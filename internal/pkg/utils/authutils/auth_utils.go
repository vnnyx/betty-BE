package authutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vnnyx/betty-BE/internal/config"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/enums"
)

type CustomClaims struct {
	ID           int64         `json:"id"`
	IsSuperAdmin bool          `json:"is_super_admin"`
	IsAdmin      bool          `json:"is_admin"`
	CompanyID    int64         `json:"company_id"`
	FranchiseID  int64         `json:"franchise_id"`
	Scopes       []enums.Scope `json:"scopes"`
	jwt.RegisteredClaims
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}

// GenerateToken generates a JWT token for the user
func GenerateToken(user *entity.User, scopes []enums.Scope) (td *TokenDetails, err error) {
	td = &TokenDetails{}

	td.AccessToken, err = GenerateJWT(user, scopes, 15*time.Minute)
	if err != nil {
		return nil, err
	}
	td.ExpiresAt = time.Now().Add(15 * time.Minute).Unix()

	td.RefreshToken, err = GenerateJWT(user, scopes, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return td, nil
}

func GenerateJWT(user *entity.User, scopes []enums.Scope, duration time.Duration) (token string, err error) {
	conf, err := config.NewConfig()
	if err != nil {
		return "", err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(conf.RSAPrivateKey))
	if err != nil {
		return "", err
	}
	claims := CustomClaims{
		ID:           user.ID,
		IsSuperAdmin: user.IsSuperAdmin,
		CompanyID:    user.CompanyID,
		FranchiseID:  user.FranchiseID,
		IsAdmin:      user.IsAdmin,
		Scopes:       scopes,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func DecodeJWT(tokenString string) (*CustomClaims, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(conf.RSAPublicKey))
	if err != nil {
		return nil, err
	}
	claims := &CustomClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func RefreshTokenRotation(oldRefreshToken string) (string, string, error) {
	claims, err := DecodeJWT(oldRefreshToken)
	if err != nil {
		return "", "", err
	}
	scopes := claims.Scopes
	newRefreshToken, err := GenerateJWT(&entity.User{
		ID:           claims.ID,
		IsSuperAdmin: claims.IsSuperAdmin,
		CompanyID:    claims.CompanyID,
		FranchiseID:  claims.FranchiseID,
		IsAdmin:      claims.IsAdmin,
	}, scopes, 24*time.Hour)
	if err != nil {
		return "", "", err
	}

	newAccessToken, err := GenerateJWT(&entity.User{
		ID:           claims.ID,
		IsSuperAdmin: claims.IsSuperAdmin,
		CompanyID:    claims.CompanyID,
		FranchiseID:  claims.FranchiseID,
		IsAdmin:      claims.IsAdmin,
	}, scopes, 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	return newRefreshToken, newAccessToken, nil
}

func ValidateJWT(tokenString string) (bool, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return false, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(conf.RSAPublicKey))
	if err != nil {
		return false, err
	}
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateSharedSecretKey(user *entity.User) (string, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return "", err
	}
	randomString, err := GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	return EncryptString(conf.EncryptKey, randomString)
}

func GenerateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func EncryptString(key string, plainText string) (string, error) {
	// Create a new cipher block from the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create a new GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce for GCM
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the data
	encrypted := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func DecryptString(key string, encryptedString string) (string, error) {
	//  Decode the encrypted data from base64
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher block from the key
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Create a new GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract the nonce from the encrypted data
	if len(encryptedData) < gcm.NonceSize() {
		return "", err
	}
	nonce, cipherText := encryptedData[:gcm.NonceSize()], encryptedData[gcm.NonceSize():]

	// Decrypt the data
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

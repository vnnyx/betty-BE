package authmiddleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vnnyx/betty-BE/internal/enums"
	"github.com/vnnyx/betty-BE/internal/errors"
	"github.com/vnnyx/betty-BE/internal/pkg/utils/authutils"
	"github.com/vnnyx/betty-BE/internal/usecase/auth"
)

type AuthMiddleware struct {
	authUC auth.AuthUC
}

func NewAuthMiddleware(authUC auth.AuthUC) *AuthMiddleware {
	return &AuthMiddleware{
		authUC: authUC,
	}
}

func (m *AuthMiddleware) AuthMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Request().Header.Peek("Authorization")
		if authHeader == nil {
			return errors.NewCustomError(string(enums.ErrUnauthorized), string(enums.ErrUnauthorized))
		}

		parts := strings.Split(string(authHeader), " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return errors.NewCustomError("Authorization header format must be 'Bearer {token}'", string(enums.ErrUnauthorized))
		}

		token := parts[1]

		decodedToken, err := authutils.DecodeJWT(string(token))
		if err != nil {
			return errors.NewCustomError(string(enums.ErrUnauthorized), string(enums.ErrUnauthorized))
		}

		if decodedToken.ExpiresAt.UnixMilli() < time.Now().UnixMilli() {
			return errors.NewCustomError(string(enums.ErrUnauthorized), string(enums.ErrUnauthorized))
		}

		valid, err := m.authUC.ValidateUser(c.Context(), decodedToken.ID)
		if err != nil {
			return errors.NewCustomError(string(enums.ErrUnauthorized), string(enums.ErrUnauthorized))
		}

		if !valid {
			return errors.NewCustomError(string(enums.ErrUnauthorized), string(enums.ErrUnauthorized))
		}

		c.Locals("userID", decodedToken.ID)
		c.Locals("isSuperAdmin", decodedToken.IsSuperAdmin)
		c.Locals("isAdmin", decodedToken.IsAdmin)
		c.Locals("companyID", decodedToken.CompanyID)
		c.Locals("franchiseID", decodedToken.FranchiseID)
		c.Locals("scopes", decodedToken.Scopes)

		return c.Next()
	}
}

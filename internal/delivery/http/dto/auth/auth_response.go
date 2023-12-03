package auth

import (
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto/company"
	"github.com/vnnyx/betty-BE/internal/enums"
)

type Credential struct {
	Scopes       []enums.Scope `json:"scopes"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresAt    int64         `json:"expires_at"`
}

type AuthDetails struct {
	ID         int64       `json:"id"`
	Credential *Credential `json:"credential"`
}

type AuthResponse struct {
	ID          int64            `json:"id"`
	Company     *company.Company `json:"company"`
	AuthDetails *AuthDetails     `json:"auth_details"`
}

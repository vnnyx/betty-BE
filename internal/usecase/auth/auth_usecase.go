package auth

import (
	"context"
	"database/sql"

	"github.com/vnnyx/betty-BE/internal/delivery/http/dto/auth"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
)

type AuthUC interface {
	GetAuthDetails(ctx context.Context, tx *sql.Tx, user *entity.User) (*auth.AuthDetails, error)
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error)
	ValidateUser(ctx context.Context, userID int64) (bool, error)
	GetNewAccessToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.AuthResponse, error)
}

package repository

import (
	"context"
	"database/sql"

	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/queryresult"
)

type UserRepository interface {
	InsertUser(ctx context.Context, tx *sql.Tx, user *entity.User) (int64, error)
	GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (user *entity.User, err error)
	GetRoleScopeByUserID(ctx context.Context, tx *sql.Tx, userID int64) (rs *queryresult.RoleScopeByUserIDResult, err error)
	GetUserByID(ctx context.Context, tx *sql.Tx, userID int64) (user *entity.User, err error)
	UpdateRefreshToken(ctx context.Context, tx *sql.Tx, userID int64, refreshToken string) error
	GetDetailUserByID(ctx context.Context, tx *sql.Tx, userID int64) (rs *queryresult.UserDetailResult, err error)
	GetUserByRefreshToken(ctx context.Context, tx *sql.Tx, refreshToken string) (user *entity.User, err error)
}

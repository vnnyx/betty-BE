package user

import (
	"context"
	"database/sql"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/queryresult"
	"github.com/vnnyx/betty-BE/internal/domain/repository"
	"github.com/vnnyx/betty-BE/internal/enums"
	"github.com/vnnyx/betty-BE/internal/errors"
	"github.com/vnnyx/betty-BE/internal/infrastructure/db"
	"github.com/vnnyx/betty-BE/internal/log"
)

type UserRepositoryImpl struct {
	logger zerolog.Logger
}

func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{
		logger: log.NewLog(),
	}
}

func (r *UserRepositoryImpl) InsertUser(ctx context.Context, tx *sql.Tx, user *entity.User) (int64, error) {
	var id int64

	args := []interface{}{
		user.IsSuperAdmin,
		user.IsAdmin,
		user.Email,
		user.Password,
		user.FullName,
		user.PhoneNumber,
		user.CompanyID,
	}

	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertUserQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert user query")
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(sqlCtx, args...).Scan(&id)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute insert user query")
		return 0, err
	}

	return int64(id), nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, tx *sql.Tx, email string) (user *entity.User, err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, GetUserByEmailQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user = &entity.User{}
	err = stmt.QueryRowContext(sqlCtx, email).Scan(
		&user.ID,
		&user.IsSuperAdmin,
		&user.IsAdmin,
		&user.Email,
		&user.Password,
		&user.RefreshToken,
		&user.FullName,
		&user.PhoneNumber,
		&user.CompanyID,
		&user.FranchiseID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewCustomError("User not found", string(enums.ErrNotFound))
		}
		r.logger.Error().Caller().Err(err).Msg("Error execute get user by email query")
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetRoleScopeByUserID(ctx context.Context, tx *sql.Tx, userID int64) (rs *queryresult.RoleScopeByUserIDResult, err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, GetRoleScopeByUserIDQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare get role scope by user id query")
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(sqlCtx, userID).Scan(
		&rs.Ids,
		&rs.RoleIds,
		&rs.Roles,
		&rs.ScopeIds,
		&rs.Scopes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewCustomError("User not found", string(enums.ErrNotFound))
		}
		r.logger.Error().Caller().Err(err).Msg("Error execute get role scope by user id query")
		return nil, err
	}

	return rs, nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, tx *sql.Tx, userID int64) (user *entity.User, err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, GetUserByIDQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare get user by id query")
		return nil, err
	}
	defer stmt.Close()

	user = &entity.User{}
	err = stmt.QueryRowContext(sqlCtx, userID).Scan(
		&user.ID,
		&user.IsSuperAdmin,
		&user.IsAdmin,
		&user.Email,
		&user.Password,
		&user.RefreshToken,
		&user.FullName,
		&user.PhoneNumber,
		&user.CompanyID,
		&user.FranchiseID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewCustomError("User not found", string(enums.ErrNotFound))
		}
		r.logger.Error().Caller().Err(err).Msg("Error execute get user by id query")
		return nil, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) UpdateRefreshToken(ctx context.Context, tx *sql.Tx, userID int64, refreshToken string) error {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, UpdateRefreshTokenQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare update refresh token query")
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(sqlCtx, refreshToken, userID)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute update refresh token query")
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) GetDetailUserByID(ctx context.Context, tx *sql.Tx, userID int64) (rs *queryresult.UserDetailResult, err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, GetDetailUserByIDQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare get detail user by id query")
		return nil, err
	}
	defer stmt.Close()

	rs = &queryresult.UserDetailResult{
		Company: &queryresult.CompanyResult{},
	}
	rs.Company.City = &entity.City{}
	rs.Company.Country = &entity.Country{}

	var franchisesJSONString string
	err = stmt.QueryRowContext(sqlCtx, userID).Scan(
		&rs.ID,
		&rs.IsSuperAdmin,
		&rs.IsAdmin,
		&rs.Email,
		&rs.Password,
		&rs.RefreshToken,
		&rs.FullName,
		&rs.PhoneNumber,
		&rs.Company.ID,
		&rs.Company.BrandName,
		&rs.Company.Address1,
		&rs.Company.Address2,
		&rs.Company.City.ID,
		&rs.Company.City.Name,
		&rs.Company.Country.ID,
		&rs.Company.Country.Name,
		&franchisesJSONString,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewCustomError("User not found", string(enums.ErrNotFound))
		}
		r.logger.Error().Caller().Err(err).Msg("Error execute get detail user by id query")
		return nil, err
	}
	if err := json.Unmarshal([]byte(franchisesJSONString), &rs.Company.Franchises); err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error unmarshal franchises json string")
		return nil, err
	}
	return rs, nil
}

func (r *UserRepositoryImpl) GetUserByRefreshToken(ctx context.Context, tx *sql.Tx, refreshToken string) (user *entity.User, err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, GetUserByRefreshTokenQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare get user by refresh token query")
		return nil, err
	}
	defer stmt.Close()

	user = &entity.User{}
	err = stmt.QueryRowContext(sqlCtx, refreshToken).Scan(
		&user.ID,
		&user.IsSuperAdmin,
		&user.IsAdmin,
		&user.Email,
		&user.Password,
		&user.RefreshToken,
		&user.FullName,
		&user.PhoneNumber,
		&user.CompanyID,
		&user.FranchiseID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewCustomError("Refresh token not valid", string(enums.ErrUnauthorized))
		}
		r.logger.Error().Caller().Err(err).Msg("Error execute get user by refresh token query")
		return nil, err
	}

	return user, nil
}

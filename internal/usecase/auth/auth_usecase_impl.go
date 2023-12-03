package auth

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto/auth"
	companydto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/company"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/repository"
	"github.com/vnnyx/betty-BE/internal/enums"
	"github.com/vnnyx/betty-BE/internal/errors"
	"github.com/vnnyx/betty-BE/internal/log"
	"github.com/vnnyx/betty-BE/internal/pkg/utils/authutils"
	"github.com/vnnyx/betty-BE/internal/pkg/utils/dbutils"
	"github.com/vnnyx/betty-BE/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

type AuthUCImpl struct {
	userRepository     repository.UserRepository
	activityRepository repository.ActivityRepository
	db                 *sql.DB
	logger             zerolog.Logger
}

func NewAuthUC(userRepository repository.UserRepository, activityRepository repository.ActivityRepository, db *sql.DB) AuthUC {
	return &AuthUCImpl{
		userRepository:     userRepository,
		activityRepository: activityRepository,
		db:                 db,
		logger:             log.NewLog(),
	}
}

func (uc *AuthUCImpl) GetAuthDetails(ctx context.Context, tx *sql.Tx, user *entity.User) (a *auth.AuthDetails, err error) {
	if tx == nil {
		tx, err = uc.db.Begin()
		if err != nil {
			uc.logger.Error().Caller().Err(err).Msg("Error begin transaction")
			return nil, err
		}

		defer func() {
			if defErr := dbutils.CommitOrRollback(tx, err); defErr != nil {
				err = defErr
			}
		}()
	}

	var scopes = []enums.Scope{
		enums.IngredientReadAccess,
		enums.IngredientCreateAccess,
		enums.IngredientUpdateAccess,
		enums.IngredientDeleteAccess,
		enums.MenuReadAccess,
		enums.MenuCreateAccess,
		enums.MenuUpdateAccess,
		enums.MenuDeleteAccess,
		enums.MenuIngredientReadAccess,
		enums.MenuIngredientCreateAccess,
		enums.MenuIngredientUpdateAccess,
		enums.MenuIngredientDeleteAccess,
		enums.FranchiseCreateAccess,
		enums.FranchiseReadAccess,
		enums.FranchiseUpdateAccess,
		enums.FranchiseDeleteAccess,
	}

	if !user.IsSuperAdmin {
		userGroup, err := uc.userRepository.GetRoleScopeByUserID(ctx, tx, user.ID)
		if err != nil {
			return nil, err
		}
		scopes = userGroup.Scopes
	}

	td, err := authutils.GenerateToken(user, scopes)
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error generate token")
		return nil, err
	}

	a = &auth.AuthDetails{
		ID: user.ID,
		Credential: &auth.Credential{
			Scopes:       scopes,
			AccessToken:  td.AccessToken,
			RefreshToken: td.RefreshToken,
			ExpiresAt:    td.ExpiresAt,
		},
	}

	return a, nil
}

func (uc *AuthUCImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	err := validation.LoginOwnerRequestValidation(req)
	if err != nil {
		return nil, err
	}

	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error begin transaction")
		return nil, err
	}
	defer func() {
		if defErr := dbutils.CommitOrRollback(tx, err); defErr != nil {
			err = defErr
		}
	}()

	user, err := uc.userRepository.GetUserByEmail(ctx, tx, req.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	userDetail, err := uc.userRepository.GetDetailUserByID(ctx, tx, user.ID)
	if err != nil {
		return nil, err
	}

	authDetails, err := uc.GetAuthDetails(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	err = uc.userRepository.UpdateRefreshToken(ctx, tx, user.ID, authDetails.Credential.RefreshToken)
	if err != nil {
		return nil, err
	}

	uc.activityRepository.BulkInsertActivity(ctx, tx, []*entity.Activity{
		{
			UserID:          user.ID,
			ActivityFieldID: int64(enums.ActivityUser),
			Action:          string(enums.ActivityActionLogin),
		},
		{
			UserID:          user.ID,
			ActivityFieldID: int64(enums.ActivityUser),
			Action:          string(enums.ActivityActionAuthenticate),
		},
	}, [][]*entity.ActivityDetail{
		{
			{
				ChangedID: nil,
				OldValue:  nil,
				NewValue:  nil,
			},
		},
		{
			{

				ChangedID: nil,
				OldValue:  nil,
				NewValue:  nil,
			},
		},
	})

	return &auth.AuthResponse{
		ID: userDetail.ID,
		Company: &companydto.Company{
			ID:            userDetail.Company.ID,
			BrandName:     userDetail.Company.BrandName,
			Franchises:    userDetail.Company.Franchises,
			StreetAddress: userDetail.Company.Address1,
			SuiteAddress:  userDetail.Company.Address2,
			City:          userDetail.Company.City,
			Country:       userDetail.Company.Country,
		},
		AuthDetails: authDetails,
	}, nil
}

func (uc *AuthUCImpl) ValidateUser(ctx context.Context, userID int64) (bool, error) {
	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error begin transaction")
		return false, err
	}
	defer func() {
		if defErr := dbutils.CommitOrRollback(tx, err); defErr != nil {
			err = defErr
		}
	}()

	user, err := uc.userRepository.GetUserByID(ctx, tx, userID)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, nil
	}

	return true, nil
}

func (uc *AuthUCImpl) GetNewAccessToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.AuthResponse, error) {
	err := validation.NewAccessTokenRequestValidation(req)
	if err != nil {
		return nil, err
	}

	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error begin transaction")
		return nil, err
	}
	defer func() {
		if defErr := dbutils.CommitOrRollback(tx, err); defErr != nil {
			err = defErr
		}
	}()

	user, err := uc.userRepository.GetUserByRefreshToken(ctx, tx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	valid, err := authutils.ValidateJWT(user.RefreshToken)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.NewCustomError("Refresh token is not valid", string(enums.ErrUnauthorized))
	}

	userDetail, err := uc.userRepository.GetDetailUserByID(ctx, tx, user.ID)
	if err != nil {
		return nil, err
	}

	authDetails, err := uc.GetAuthDetails(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	err = uc.userRepository.UpdateRefreshToken(ctx, tx, user.ID, authDetails.Credential.RefreshToken)
	if err != nil {
		return nil, err
	}

	uc.activityRepository.BulkInsertActivity(ctx, tx, []*entity.Activity{
		{
			UserID:          user.ID,
			ActivityFieldID: int64(enums.ActivityUser),
			Action:          string(enums.ActivityActionAuthenticate),
		},
	}, [][]*entity.ActivityDetail{
		{
			{
				ChangedID: nil,
				OldValue:  nil,
				NewValue:  nil,
			},
		},
	})

	return &auth.AuthResponse{
		ID: userDetail.ID,
		Company: &companydto.Company{
			ID:            userDetail.Company.ID,
			BrandName:     userDetail.Company.BrandName,
			Franchises:    userDetail.Company.Franchises,
			StreetAddress: userDetail.Company.Address1,
			SuiteAddress:  userDetail.Company.Address2,
			City:          userDetail.Company.City,
			Country:       userDetail.Company.Country,
		},
		AuthDetails: authDetails,
	}, nil
}

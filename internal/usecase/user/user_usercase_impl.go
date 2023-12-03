package user

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	authdto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/auth"
	companydto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/company"
	userdto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/user"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/repository"
	"github.com/vnnyx/betty-BE/internal/enums"
	"github.com/vnnyx/betty-BE/internal/log"
	"github.com/vnnyx/betty-BE/internal/pkg/utils/dbutils"
	"github.com/vnnyx/betty-BE/internal/usecase/auth"
	"github.com/vnnyx/betty-BE/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

type UserUCImpl struct {
	companyRepository  repository.CompanyRepository
	userRepository     repository.UserRepository
	activityRepository repository.ActivityRepository
	menuRepository     repository.MenuRepository
	authUC             auth.AuthUC
	db                 *sql.DB
	logger             zerolog.Logger
}

func NewUserUC(
	companyRepository repository.CompanyRepository,
	userRepository repository.UserRepository,
	activityRepository repository.ActivityRepository,
	menuRepository repository.MenuRepository,
	authUC auth.AuthUC,
	db *sql.DB,
) UserUC {
	return &UserUCImpl{
		companyRepository:  companyRepository,
		userRepository:     userRepository,
		activityRepository: activityRepository,
		menuRepository:     menuRepository,
		authUC:             authUC,
		db:                 db,
		logger:             log.NewLog(),
	}
}

// CreateOwner is a function to create a new owner
func (uc *UserUCImpl) CreateOwner(ctx context.Context, req *userdto.CreateOwnerRequest) (*authdto.AuthResponse, error) {
	err := validation.CreateOwnerRequestValidation(req)
	if err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error generate password hash")
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

	company, err := uc.companyRepository.InsertCompany(
		ctx, tx,
		&entity.Company{
			BrandName:     req.Company.BrandName,
			FranchiseName: req.Company.FranchiseName,
			Address1:      req.Company.StreetAddress,
			Address2:      *req.Company.SuiteAddress,
			CityID:        req.Company.CityID,
			City:          &entity.City{},
			CountryID:     req.Company.CountryID,
			Country:       &entity.Country{},
			PostalCode:    *req.Company.ZipCode,
		},
	)
	if err != nil {
		return nil, err
	}

	franchises := make([]*entity.Franchise, 0)
	franchise, err := uc.companyRepository.InsertFranchise(ctx, tx, &entity.Franchise{
		Name:      req.Company.FranchiseName,
		CompanyID: company.ID,
	})
	if err != nil {
		return nil, err
	}
	franchises = append(franchises, franchise)

	user := &entity.User{
		IsSuperAdmin: true,
		IsAdmin:      false,
		Email:        req.Email,
		Password:     string(passwordHash),
		FullName:     req.Fullname,
		PhoneNumber:  req.PhoneNumber,
		CompanyID:    company.ID,
		FranchiseID:  franchise.ID,
	}

	userID, err := uc.userRepository.InsertUser(
		ctx, tx, user,
	)
	if err != nil {
		return nil, err
	}
	user.ID = userID

	authDetails, err := uc.authUC.GetAuthDetails(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	if err = uc.userRepository.UpdateRefreshToken(ctx, tx, int64(userID), authDetails.Credential.RefreshToken); err != nil {
		return nil, err
	}

	if _, err = uc.menuRepository.InsertUnits(ctx, tx, []*entity.Unit{
		{
			IsInternational: true,
			Name:            "Kg",
			ConversionRate:  1,
			ConversionID:    nil,
			FranchiseID:     franchise.ID,
		},
		{
			IsInternational: true,
			Name:            "Liter",
			ConversionRate:  1,
			ConversionID:    nil,
			FranchiseID:     franchise.ID,
		},
	}); err != nil {
		return nil, err
	}

	activities := []*entity.Activity{
		{
			UserID:          int64(userID),
			ActivityFieldID: int64(enums.ActivityUser),
			Action:          string(enums.ActivityActionCreate),
		},
		{
			UserID:          int64(userID),
			ActivityFieldID: int64(enums.ActivityCompany),
			Action:          string(enums.ActivityActionCreate),
		},
		{
			UserID:          int64(userID),
			ActivityFieldID: int64(enums.ActivityFranchise),
			Action:          string(enums.ActivityActionCreate),
		},
	}

	activityDetails := [][]*entity.ActivityDetail{
		{
			{
				ChangedID: &userID,
				OldValue:  nil,
				NewValue:  &user.Email,
			},
		},
		{
			{
				ChangedID: &company.ID,
				OldValue:  nil,
				NewValue:  &req.Company.BrandName,
			},
		},
		{
			{
				ChangedID: &franchise.ID,
				OldValue:  nil,
				NewValue:  &req.Company.FranchiseName,
			},
		},
	}

	err = uc.activityRepository.BulkInsertActivity(ctx, tx, activities, activityDetails)
	if err != nil {
		return nil, err
	}

	return &authdto.AuthResponse{
		ID: int64(userID),
		Company: &companydto.Company{
			ID:            company.ID,
			BrandName:     req.Company.BrandName,
			Franchises:    franchises,
			StreetAddress: req.Company.StreetAddress,
			SuiteAddress:  *req.Company.SuiteAddress,
			City:          company.City,
			Country:       company.Country,
		},
		AuthDetails: authDetails,
	}, nil
}

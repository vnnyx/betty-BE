package company

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/repository"
	"github.com/vnnyx/betty-BE/internal/infrastructure/db"
	"github.com/vnnyx/betty-BE/internal/log"
)

type CompanyRepositoryImpl struct {
	logger zerolog.Logger
}

func NewCompanyRepository() repository.CompanyRepository {
	return &CompanyRepositoryImpl{
		logger: log.NewLog(),
	}
}

func (r *CompanyRepositoryImpl) InsertCompany(ctx context.Context, tx *sql.Tx, company *entity.Company) (*entity.Company, error) {
	args := []interface{}{
		company.BrandName,
		company.FranchiseName,
		company.Address1,
		company.Address2,
		company.CityID,
		company.CountryID,
		company.PostalCode,
	}

	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertCompanyQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert company query")
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(sqlCtx, args...).Scan(
		&company.ID,
		&company.BrandName,
		&company.FranchiseName,
		&company.Address1,
		&company.Address2,
		&company.City.ID,
		&company.City.Name,
		&company.City.CountryID,
		&company.City.Latitude,
		&company.City.Longitude,
		&company.Country.ID,
		&company.Country.ISO,
		&company.Country.Name,
		&company.PostalCode,
	)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute insert company query")
		return nil, err
	}

	return company, nil
}

func (r *CompanyRepositoryImpl) InsertFranchise(ctx context.Context, tx *sql.Tx, franchise *entity.Franchise) (*entity.Franchise, error) {
	args := []interface{}{
		franchise.Name,
		franchise.CompanyID,
	}

	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertFranchiseQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert franchise query")
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(sqlCtx, args...).Scan(
		&franchise.ID,
		&franchise.Name,
		&franchise.CompanyID,
	)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute insert franchise query")
		return nil, err
	}

	return franchise, nil
}

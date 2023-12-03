package repository

import (
	"context"
	"database/sql"

	"github.com/vnnyx/betty-BE/internal/domain/entity"
)

type CompanyRepository interface {
	InsertCompany(ctx context.Context, tx *sql.Tx, company *entity.Company) (*entity.Company, error)
	InsertFranchise(ctx context.Context, tx *sql.Tx, franchise *entity.Franchise) (*entity.Franchise, error)
}

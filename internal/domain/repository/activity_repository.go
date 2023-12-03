package repository

import (
	"context"
	"database/sql"

	"github.com/vnnyx/betty-BE/internal/domain/entity"
)

type ActivityRepository interface {
	BulkInsertActivity(ctx context.Context, tx *sql.Tx, activities []*entity.Activity, activityDetails [][]*entity.ActivityDetail) error
}

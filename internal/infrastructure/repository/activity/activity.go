package activity

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/repository"
	"github.com/vnnyx/betty-BE/internal/infrastructure/db"
	"github.com/vnnyx/betty-BE/internal/log"
)

type ActivityRepositoryImpl struct {
	logger zerolog.Logger
}

func NewActivityRepository() repository.ActivityRepository {
	return &ActivityRepositoryImpl{
		logger: log.NewLog(),
	}
}

func (r *ActivityRepositoryImpl) BulkInsertActivity(ctx context.Context, tx *sql.Tx, activities []*entity.Activity, activityDetails [][]*entity.ActivityDetail) error {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmtActivity, err := tx.PrepareContext(sqlCtx, InsertActivityQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert activity query")
		return err
	}
	defer stmtActivity.Close()

	stmtDetail, err := tx.PrepareContext(sqlCtx, InsertActivityDetailQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert activity detail query")
		return err
	}
	defer stmtDetail.Close()

	for i, activity := range activities {
		args := []interface{}{
			activity.ActivityFieldID,
			activity.Action,
			activity.UserID,
		}
		var id int64
		err := stmtActivity.QueryRowContext(sqlCtx, args...).Scan(&id)
		if err != nil {
			r.logger.Error().Caller().Err(err).Msg("Error execute insert activity query")
			return err
		}

		for _, activityDetail := range activityDetails[i] {
			args := []interface{}{
				id,
				activityDetail.ChangedID,
				activityDetail.OldValue,
				activityDetail.NewValue,
			}
			_, err := stmtDetail.ExecContext(sqlCtx, args...)
			if err != nil {
				r.logger.Error().Caller().Err(err).Msg("Error execute insert activity detail query")
				return err
			}
		}
	}

	return nil
}

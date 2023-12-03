package attachmentfile

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/repository"
	"github.com/vnnyx/betty-BE/internal/infrastructure/db"
	"github.com/vnnyx/betty-BE/internal/log"
)

type AttachmentFileRepositoryImpl struct {
	logger zerolog.Logger
}

func NewAttachmentFileRepository() repository.AttachmentFileRepository {
	return &AttachmentFileRepositoryImpl{
		logger: log.NewLog(),
	}
}

func (r *AttachmentFileRepositoryImpl) BulkInsertAttachmentFile(ctx context.Context, tx *sql.Tx, attachmentFiles []*entity.AttachmentFile) ([]int64, error) {
	var ids []int64

	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertAttachmentFileQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert attachment file query")
		return nil, err
	}
	defer stmt.Close()

	for _, attachmentFile := range attachmentFiles {
		var id int64

		args := []interface{}{
			attachmentFile.Path,
		}

		err = stmt.QueryRowContext(sqlCtx, args...).Scan(&id)
		if err != nil {
			r.logger.Error().Caller().Err(err).Msg("Error execute insert attachment file query")
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (r *AttachmentFileRepositoryImpl) BulkInsertPhotoMenu(ctx context.Context, tx *sql.Tx, photoMenus []*entity.PhotoMenu) ([]int64, error) {
	var ids []int64

	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertPhotoMenuQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert photo menu query")
		return nil, err
	}
	defer stmt.Close()

	for _, photoMenu := range photoMenus {
		var id int64

		args := []interface{}{
			photoMenu.MenuID,
			photoMenu.PhotoID,
		}

		err = stmt.QueryRowContext(sqlCtx, args...).Scan(&id)
		if err != nil {
			r.logger.Error().Caller().Err(err).Msg("Error execute insert photo menu query")
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

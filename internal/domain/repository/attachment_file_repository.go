package repository

import (
	"context"
	"database/sql"

	"github.com/vnnyx/betty-BE/internal/domain/entity"
)

type AttachmentFileRepository interface {
	BulkInsertAttachmentFile(ctx context.Context, tx *sql.Tx, attachmentFiles []*entity.AttachmentFile) ([]int64, error)
	BulkInsertPhotoMenu(ctx context.Context, tx *sql.Tx, photoMenus []*entity.PhotoMenu) ([]int64, error)
}

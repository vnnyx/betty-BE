package attachmentfile

const (
	InsertAttachmentFileQuery = `
		INSERT INTO attachment_file (
			path
		)
		VALUES ($1) RETURNING id
	`

	InsertPhotoMenuQuery = `
		INSERT INTO photo_menu(
			menu_id,
			attachment_file_id
		)
		VALUES ($1, $2) RETURNING id
	`
)

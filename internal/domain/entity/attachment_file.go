package entity

type AttachmentFile struct {
	ID        int64  `json:"id"`
	Path      string `json:"path"`
	IsDeleted bool   `json:"is_deleted"`
}

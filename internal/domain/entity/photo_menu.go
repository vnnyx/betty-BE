package entity

type PhotoMenu struct {
	ID        int64           `json:"id"`
	MenuID    int64           `json:"-"`
	Menu      *Menu           `json:"menu,omitempty"`
	PhotoID   int64           `json:"-"`
	Photo     *AttachmentFile `json:"photo,omitempty"`
	IsDeleted bool            `json:"is_deleted"`
}

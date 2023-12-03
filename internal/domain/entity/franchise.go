package entity

type Franchise struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	CompanyID int64           `json:"-"`
	Company   *Company        `json:"company,omitempty"`
	PhotoID   int64           `json:"-"`
	Photo     *AttachmentFile `json:"photo,omitempty"`
}

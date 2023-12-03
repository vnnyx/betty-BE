package entity

type Transaction struct {
	ID           int64        `json:"id"`
	UserID       int64        `json:"-"`
	User         *User        `json:"user,omitempty"`
	VarianMenuID int64        `json:"-"`
	VariantMenu  *VariantMenu `json:"variant_menu,omitempty"`
	Quantity     int32        `json:"quantity"`
	TotalPrice   int64        `json:"total_price"`
}

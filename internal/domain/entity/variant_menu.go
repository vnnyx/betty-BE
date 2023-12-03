package entity

type VariantMenu struct {
	ID        int64  `json:"id"`
	VariantID int64  `json:"variant_id"`
	MenuID    int64  `json:"menu_id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	IsDeleted bool   `json:"is_deleted"`
}

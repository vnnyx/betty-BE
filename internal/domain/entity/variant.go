package entity

type Variant struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsMulti   bool   `json:"is_multi"`
	IsDeleted bool   `json:"is_deleted"`
}

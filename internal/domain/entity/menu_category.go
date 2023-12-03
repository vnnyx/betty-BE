package entity

type MenuCategory struct {
	ID         int64     `json:"id"`
	MenuID     int64     `json:"-"`
	Menu       *Menu     `json:"menu,omitempty"`
	CategoryID int64     `json:"-"`
	Category   *Category `json:"category,omitempty"`
}

package entity

type Ingredient struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	MinimumStock float64 `json:"minimum_stock"`
	RecentStock  float64 `json:"recent_stock"`
	Price        float64 `json:"price"`
	UnitID       int64   `json:"-"`
	Unit         *Unit   `json:"unit,omitempty"`
	IsDeleted    bool    `json:"is_deleted"`
}

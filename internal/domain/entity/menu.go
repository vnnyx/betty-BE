package entity

type Menu struct {
	ID              int64      `json:"id"`
	Matchcode       string     `json:"matchcode"`
	Name            string     `json:"name"`
	Stock           int32      `json:"stock"`
	Price           float64    `json:"price"`
	Description     string     `json:"description"`
	StatusID        int32      `json:"-"`
	Status          *Status    `json:"status,omitempty"`
	FranchiseID     int64      `json:"-"`
	Franchise       *Franchise `json:"franchise,omitempty"`
	PopularityCount int64      `json:"popularity_count"`
	IsDeleted       bool       `json:"is_deleted"`
}

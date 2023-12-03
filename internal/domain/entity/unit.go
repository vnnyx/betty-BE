package entity

type Unit struct {
	ID              int64      `json:"id"`
	Name            string     `json:"name"`
	IsInternational bool       `json:"is_international"`
	ConversionRate  float64    `json:"conversion_rate"`
	ConversionID    *int64     `json:"conversion"`
	FranchiseID     int64      `json:"-"`
	Franchise       *Franchise `json:"franchise,omitempty"`
	IsDeleted       bool       `json:"is_deleted"`
}

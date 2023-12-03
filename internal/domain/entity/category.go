package entity

type Category struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Color       string     `json:"color"`
	FranchiseID int64      `json:"-"`
	Franchise   *Franchise `json:"franchise,omitempty"`
}

package entity

type Country struct {
	ID   int64  `json:"id"`
	ISO  string `json:"iso,omitempty"`
	Name string `json:"name"`
}

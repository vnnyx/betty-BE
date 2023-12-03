package entity

type City struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name,omitempty"`
	CountryID int64    `json:"-"`
	Country   *Country `json:"country,omitempty"`
	Latitude  float64  `json:"latitude,omitempty"`
	Longitude float64  `json:"longitude,omitempty"`
}

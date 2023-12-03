package entity

type Company struct {
	ID            int64           `json:"id"`
	BrandName     string          `json:"brand_name"`
	FranchiseName string          `json:"franchise_name"`
	Address1      string          `json:"address_1"`
	Address2      string          `json:"address_2"`
	CityID        int64           `json:"-"`
	City          *City           `json:"city,omitempty"`
	CountryID     int64           `json:"-"`
	Country       *Country        `json:"country,omitempty"`
	PhotoID       int64           `json:"-"`
	Photo         *AttachmentFile `json:"photo,omitempty"`
	PostalCode    string          `json:"postal_code,"`
}

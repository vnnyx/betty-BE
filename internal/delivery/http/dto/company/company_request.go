package company

type CreateCompanyRequest struct {
	BrandName     string  `json:"brand_name" validate:"required"`
	FranchiseName string  `json:"franchise_name" validate:"required"`
	StreetAddress string  `json:"street_address" validate:"required"`
	SuiteAddress  *string `json:"suite_address,omitempty"`
	CityID        int64   `json:"city_id" validate:"required"`
	CountryID     int64   `json:"country_id" validate:"required"`
	ZipCode       *string `json:"zip_code,omitempty"`
	Photo         *string `json:"photo,omitempty"`
}

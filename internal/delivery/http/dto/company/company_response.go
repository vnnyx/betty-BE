package company

import "github.com/vnnyx/betty-BE/internal/domain/entity"

type Company struct {
	ID            int64               `json:"id"`
	BrandName     string              `json:"brand_name"`
	Franchises    []*entity.Franchise `json:"franchises"`
	StreetAddress string              `json:"street_address"`
	SuiteAddress  string              `json:"suite_address"`
	City          *entity.City        `json:"city"`
	Country       *entity.Country     `json:"country"`
}

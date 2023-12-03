package queryresult

import (
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/enums"
)

type RoleScopeByUserIDResult struct {
	Ids      []int64
	RoleIds  []int64
	Roles    []string
	ScopeIds []int64
	Scopes   []enums.Scope
}

type InsertedMenuResult struct {
	MenuID      int64
	Matchcode   string
	Name        string
	Description string
	Price       int64
	Stock       int64
	FranchiseID int64
	Photos      string
}

type MenuIngredientResult struct {
	ID              int64
	IngredientID    int64
	IngredientName  string
	MenuID          int64
	Quantity        float64
	UnitID          int64
	IsInternational bool
	ConversionRate  float64
	ConversionID    int64
	RecentStock     float64
}

type CompanyResult struct {
	ID         int64
	BrandName  string
	Franchises []*entity.Franchise
	Address1   string
	Address2   string
	City       *entity.City
	Country    *entity.Country
}

type UserDetailResult struct {
	ID           int64
	IsSuperAdmin bool
	IsAdmin      bool
	Email        string
	Password     string
	RefreshToken string
	FullName     string
	PhoneNumber  string
	Company      *CompanyResult
}

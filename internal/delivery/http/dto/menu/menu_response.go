package menu

import "github.com/vnnyx/betty-BE/internal/delivery/http/dto"

type MenuResponse struct {
	ID               int64 `json:"id"`
	TotalBaseProduct int64 `json:"total_base_product"`
	Stock            int64 `json:"stock"`
	Price            int64 `json:"price"`
	Status           int64 `json:"status"`
}

type CreateMenuResponse struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Matchcode   string       `json:"matchcode"`
	Description string       `json:"description"`
	Price       int64        `json:"price"`
	Stock       int64        `json:"stock"`
	FranchiseID int64        `json:"franchise_id"`
	Photos      []*dto.Photo `json:"photos"`
}

type BaseProduct struct {
	ID                    int64   `json:"id"`
	Name                  string  `json:"name"`
	QuantityPerMenu       float64 `json:"quantity_per_menu"`
	TotalQuantityForServe float64 `json:"total_quantity_for_serve"`
	RecentStock           float64 `json:"recent_stock"`
}

type CreateRecipeMenuResponse struct {
	MenuID            int64          `json:"menu_id"`
	EstimatedServings int64          `json:"estimated_servings"`
	SyncAt            int64          `json:"sync_at"`
	Recipes           []*BaseProduct `json:"recipes"`
}

type CreateBaseProductResponse struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	MinimumStock float64 `json:"minimum_stock"`
	RecentStock  float64 `json:"recent_stock"`
	Price        float64 `json:"price"`
	FranchiseID  int64   `json:"franchise_id"`
}

type CreateMenuCategoryResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

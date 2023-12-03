package menu

type CreateMenuRequest struct {
	Name          string   `json:"name" validate:"required"`
	ShortMenuName string   `json:"short_menu_name" validate:"required"`
	Price         float64  `json:"price" validate:"required"`
	CategoryID    int64    `json:"category_id" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	Photos        []string `json:"photos" validate:"required"`
}

type CreateUnitRequest struct {
	Name           string  `json:"name" validate:"required"`
	ConversionID   int64   `json:"conversion_id" validate:"required"`
	ConversionRate float64 `json:"conversion_rate" validate:"required"`
}

type CreateBaseProductRequest struct {
	BaseProduceName string             `json:"base_product_name" validate:"required"`
	CustomUnit      *CreateUnitRequest `json:"custom_unit,omitempty"`
	UnitID          int64              `json:"unit_id,omitempty"`
	RequiredStock   float64            `json:"required_stock" validate:"required"`
	RecentStock     float64            `json:"recent_stock" validate:"required"`
	Price           float64            `json:"price" validate:"required"`
}

type CreateBaseProductsRequest []*CreateBaseProductRequest

type CreateMenuRecipeRequest struct {
	MenuID      int64         `json:"menu_id" validate:"required"`
	BaseProduct *[]Ingredient `json:"base_product" validate:"required"`
}

type Ingredient struct {
	BaseProductID int64   `json:"base_product_id" validate:"required"`
	Quantity      float64 `json:"quantity" validate:"required"`
}

type CreateMenuCategoryRequest struct {
	Name  string `json:"name" validate:"required"`
	Color string `json:"color" validate:"required"`
}

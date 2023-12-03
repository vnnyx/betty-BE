package menu

import (
	"context"

	"github.com/vnnyx/betty-BE/internal/delivery/http/dto/menu"
)

type MenuUC interface {
	CreateMenu(ctx context.Context, req *menu.CreateMenuRequest) (*menu.CreateMenuResponse, error)
	CreateMenuRecipe(ctx context.Context, req *menu.CreateMenuRecipeRequest) (*menu.CreateRecipeMenuResponse, error)
	CreateBaseProduct(ctx context.Context, req *menu.CreateBaseProductRequest) (*menu.CreateBaseProductResponse, error)
	CreateMenuCategory(ctx context.Context, req *menu.CreateMenuCategoryRequest) (*menu.CreateMenuCategoryResponse, error)
}

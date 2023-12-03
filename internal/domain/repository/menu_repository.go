package repository

import (
	"context"
	"database/sql"

	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/queryresult"
)

type MenuRepository interface {
	InsertMenu(ctx context.Context, tx *sql.Tx, menu *entity.Menu) (*entity.Menu, error)
	InsertMenuIngredient(ctx context.Context, tx *sql.Tx, menuIngredient []*entity.MenuIngredient) (rs []*queryresult.MenuIngredientResult, err error)
	UpdateMenu(ctx context.Context, tx *sql.Tx, menu *entity.Menu, ingredient *entity.Ingredient) (err error)
	InsertIngredients(ctx context.Context, tx *sql.Tx, ingredients []*entity.Ingredient) (ids []int64, err error)
	InsertUnits(ctx context.Context, tx *sql.Tx, units []*entity.Unit) (unitIDs []int64, err error)
	UpdateMenuPopularity(ctx context.Context, tx *sql.Tx, menuID int64) (err error)
	GeIngredientByMenuID(ctx context.Context, tx *sql.Tx, menuID int64) (rs []*queryresult.MenuIngredientResult, err error)
	InsertCategory(ctx context.Context, tx *sql.Tx, category *entity.Category) (*entity.Category, error)
	InsertMenuCategory(ctx context.Context, tx *sql.Tx, menuCategory *entity.MenuCategory) (id int64, err error)
}

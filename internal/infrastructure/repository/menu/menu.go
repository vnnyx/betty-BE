package menu

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/queryresult"
	"github.com/vnnyx/betty-BE/internal/domain/repository"
	"github.com/vnnyx/betty-BE/internal/infrastructure/db"
	"github.com/vnnyx/betty-BE/internal/log"
)

type MenuRepositoryImpl struct {
	logger zerolog.Logger
}

func NewMenuRepository() repository.MenuRepository {
	return &MenuRepositoryImpl{
		logger: log.NewLog(),
	}
}

func (r *MenuRepositoryImpl) InsertMenu(ctx context.Context, tx *sql.Tx, menu *entity.Menu) (*entity.Menu, error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertMenuQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert menu query")
		return nil, err
	}

	args := []interface{}{
		menu.Matchcode,
		menu.Name,
		menu.Price,
		menu.Description,
		menu.FranchiseID,
	}

	err = stmt.QueryRowContext(sqlCtx, args...).Scan(
		&menu.ID,
		&menu.Matchcode,
		&menu.Name,
		&menu.Description,
		&menu.Price,
		&menu.Stock,
		&menu.FranchiseID,
	)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute insert menu query")
		return nil, err
	}

	return menu, nil
}

func (r *MenuRepositoryImpl) InsertMenuIngredient(ctx context.Context, tx *sql.Tx, menuIngredient []*entity.MenuIngredient) (rs []*queryresult.MenuIngredientResult, err error) {
	rs = make([]*queryresult.MenuIngredientResult, 0)
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertMenuIngredientQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert menu ingredient query")
		return nil, err
	}

	for _, ingredient := range menuIngredient {
		args := []interface{}{
			ingredient.IngredientID,
			ingredient.MenuID,
			ingredient.Quantity,
		}
		var res = &queryresult.MenuIngredientResult{}
		err = stmt.QueryRowContext(sqlCtx, args...).Scan(
			&res.ID,
			&res.IngredientID,
			&res.IngredientName,
			&res.MenuID,
			&res.Quantity,
			&res.UnitID,
			&res.IsInternational,
			&res.ConversionRate,
			&res.ConversionID,
			&res.RecentStock,
		)
		if err != nil {
			r.logger.Error().Caller().Err(err).Msg("Error execute insert menu ingredient query")
			return nil, err
		}
		rs = append(rs, res)
	}
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute insert menu ingredient query")
		return nil, err
	}

	return rs, nil
}

func (r *MenuRepositoryImpl) UpdateMenu(ctx context.Context, tx *sql.Tx, menu *entity.Menu, ingredient *entity.Ingredient) (err error) {
	return
}

func (r *MenuRepositoryImpl) InsertIngredients(ctx context.Context, tx *sql.Tx, ingredients []*entity.Ingredient) (ids []int64, err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertIngredientQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert ingredient query")
		return nil, err
	}

	for _, ingredient := range ingredients {
		args := []interface{}{
			ingredient.Name,
			ingredient.MinimumStock,
			ingredient.RecentStock,
			ingredient.Price,
			ingredient.UnitID,
			ctx.Value("franchiseID").(int64),
		}
		var id int64
		err = stmt.QueryRowContext(sqlCtx, args...).Scan(&id)
		if err != nil {
			r.logger.Error().Caller().Err(err).Msg("Error execute insert ingredient query")
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (r *MenuRepositoryImpl) InsertUnits(ctx context.Context, tx *sql.Tx, units []*entity.Unit) (unitIDs []int64, err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertUnitQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert unit query")
		return nil, err
	}

	for _, unit := range units {
		args := []interface{}{
			unit.Name,
			unit.IsInternational,
			unit.ConversionRate,
			unit.ConversionID,
			unit.FranchiseID,
		}
		var id int64
		err = stmt.QueryRowContext(sqlCtx, args...).Scan(&id)
		if err != nil {
			r.logger.Error().Caller().Err(err).Msg("Error execute insert unit query")
			return nil, err
		}
		unitIDs = append(unitIDs, id)
	}

	return unitIDs, nil
}

func (r *MenuRepositoryImpl) UpdateMenuPopularity(ctx context.Context, tx *sql.Tx, menuID int64) (err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, UpdatePopularityCountQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare update popularity count query")
		return err
	}

	_, err = stmt.ExecContext(sqlCtx, menuID)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute update popularity count query")
		return err
	}

	return nil
}

func (r *MenuRepositoryImpl) GeIngredientByMenuID(ctx context.Context, tx *sql.Tx, menuID int64) (rs []*queryresult.MenuIngredientResult, err error) {
	rs = make([]*queryresult.MenuIngredientResult, 0)
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, GetMenuIngredientByMenuIDQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare get menu ingredient by menu id query")
		return nil, err
	}

	rows, err := stmt.QueryContext(sqlCtx, menuID)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute get menu ingredient by menu id query")
		return nil, err
	}

	for rows.Next() {
		var res = &queryresult.MenuIngredientResult{}
		err = rows.Scan(
			&res.ID,
			&res.IngredientID,
			&res.IngredientName,
			&res.MenuID,
			&res.Quantity,
			&res.UnitID,
			&res.IsInternational,
			&res.ConversionRate,
			&res.ConversionID,
			&res.RecentStock,
		)
		if err != nil {
			r.logger.Error().Caller().Err(err).Msg("Error scan get menu ingredient by menu id query")
			return nil, err
		}
		rs = append(rs, res)
	}

	return rs, nil
}

func (r *MenuRepositoryImpl) InsertCategory(ctx context.Context, tx *sql.Tx, category *entity.Category) (*entity.Category, error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertCategoryQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert category query")
		return nil, err
	}

	err = stmt.QueryRowContext(sqlCtx, category.Name, category.Color, category.FranchiseID).Scan(
		&category.ID,
		&category.Name,
		&category.Color,
	)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute insert category query")
		return nil, err
	}

	return category, nil
}

func (r *MenuRepositoryImpl) InsertMenuCategory(ctx context.Context, tx *sql.Tx, menuCategory *entity.MenuCategory) (id int64, err error) {
	sqlCtx, cancel := db.NewCockroachContext()
	defer cancel()

	stmt, err := tx.PrepareContext(sqlCtx, InsertMenuCategoryQuery)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error prepare insert menu category query")
		return 0, err
	}

	err = stmt.QueryRowContext(sqlCtx, menuCategory.MenuID, menuCategory.CategoryID).Scan(&id)
	if err != nil {
		r.logger.Error().Caller().Err(err).Msg("Error execute insert menu category query")
		return 0, err
	}

	return id, nil
}

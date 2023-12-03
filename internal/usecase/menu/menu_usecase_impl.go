package menu

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/vnnyx/betty-BE/internal/config"
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto"
	menudto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/menu"
	"github.com/vnnyx/betty-BE/internal/domain/entity"
	"github.com/vnnyx/betty-BE/internal/domain/queryresult"
	"github.com/vnnyx/betty-BE/internal/domain/repository"
	"github.com/vnnyx/betty-BE/internal/enums"
	"github.com/vnnyx/betty-BE/internal/log"
	"github.com/vnnyx/betty-BE/internal/pkg/utils"
	"github.com/vnnyx/betty-BE/internal/pkg/utils/dbutils"
	"github.com/vnnyx/betty-BE/internal/validation"
)

type MenuUCImpl struct {
	menuRepo           repository.MenuRepository
	activityRepo       repository.ActivityRepository
	attachmentFileRepo repository.AttachmentFileRepository
	db                 *sql.DB
	logger             zerolog.Logger
}

func NewMenuUC(
	menuRepo repository.MenuRepository,
	activityRepo repository.ActivityRepository,
	attachmentFileRepo repository.AttachmentFileRepository,
	db *sql.DB) MenuUC {
	return &MenuUCImpl{
		menuRepo:           menuRepo,
		activityRepo:       activityRepo,
		attachmentFileRepo: attachmentFileRepo,
		db:                 db,
		logger:             log.NewLog(),
	}
}

func (uc *MenuUCImpl) CreateMenu(ctx context.Context, req *menudto.CreateMenuRequest) (*menudto.CreateMenuResponse, error) {
	err := validation.CreateMenuRequestValidation(req)
	if err != nil {
		return nil, err
	}

	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error begin transaction")
		return nil, err
	}

	defer func() {
		if defErr := dbutils.CommitOrRollback(tx, err); defErr != nil {
			err = defErr
		}
	}()

	menu, err := uc.menuRepo.InsertMenu(ctx, tx, &entity.Menu{
		Name:        req.Name,
		Franchise:   &entity.Franchise{},
		FranchiseID: ctx.Value("franchiseID").(int64),
		Status:      &entity.Status{},
		Matchcode:   req.ShortMenuName,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	_, err = uc.menuRepo.InsertMenuCategory(ctx, tx, &entity.MenuCategory{
		MenuID:     menu.ID,
		CategoryID: req.CategoryID,
	})
	if err != nil {
		return nil, err
	}

	photosPath := make([]*entity.AttachmentFile, 0)
	for i, photo := range req.Photos {
		key := fmt.Sprintf("images/%s_%v_%v.png", req.ShortMenuName, time.Now().UnixMilli(), i)
		photosPath = append(photosPath, &entity.AttachmentFile{
			Path: key,
		})
		_, err = utils.UploadToS3(key, photo)
		if err != nil {
			return nil, err
		}
	}

	photoIds, err := uc.attachmentFileRepo.BulkInsertAttachmentFile(ctx, tx, photosPath)
	if err != nil {
		return nil, err
	}

	photoMenu := make([]*entity.PhotoMenu, 0)
	photos := make([]*dto.Photo, 0)
	photoActivityDetails := make([]*entity.ActivityDetail, 0)
	for i, id := range photoIds {
		photos = append(photos, &dto.Photo{
			ID:       id,
			PhotoURL: fmt.Sprintf("%s/%v", conf.AWS.S3URL, photosPath[i].Path),
		})
		photoActivityDetails = append(photoActivityDetails, &entity.ActivityDetail{
			ChangedID: &id,
			OldValue:  nil,
			NewValue:  &photosPath[i].Path,
		})
		photoMenu = append(photoMenu, &entity.PhotoMenu{
			MenuID:  menu.ID,
			PhotoID: id,
		})
	}

	_, err = uc.attachmentFileRepo.BulkInsertPhotoMenu(ctx, tx, photoMenu)
	if err != nil {
		return nil, err
	}

	uc.activityRepo.BulkInsertActivity(ctx, tx, []*entity.Activity{
		{
			UserID:          ctx.Value("userID").(int64),
			ActivityFieldID: int64(enums.ActivityMenu),
			Action:          string(enums.ActivityActionCreate),
		},
		{
			UserID:          ctx.Value("userID").(int64),
			ActivityFieldID: int64(enums.ActivityAttachmentFile),
			Action:          string(enums.ActivityActionCreate),
		},
	}, [][]*entity.ActivityDetail{
		{
			{
				ChangedID: &menu.ID,
				OldValue:  nil,
				NewValue:  &req.Name,
			},
		},
		photoActivityDetails,
	})

	return &menudto.CreateMenuResponse{
		ID:          menu.ID,
		Name:        menu.Name,
		Matchcode:   menu.Matchcode,
		Description: menu.Description,
		Price:       int64(menu.Price),
		Stock:       int64(menu.Stock),
		FranchiseID: menu.FranchiseID,
		Photos:      photos,
	}, nil
}

func (uc *MenuUCImpl) CreateMenuRecipe(ctx context.Context, req *menudto.CreateMenuRecipeRequest) (*menudto.CreateRecipeMenuResponse, error) {
	err := validation.CreateMenuRecipeRequestValidation(req)
	if err != nil {
		return nil, err
	}

	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error begin transaction")
		return nil, err
	}

	defer func() {
		if defErr := dbutils.CommitOrRollback(tx, err); defErr != nil {
			err = defErr
		}
	}()

	menuIngredients, err := uc.menuRepo.GeIngredientByMenuID(ctx, tx, req.MenuID)

	recipes := make([]*entity.MenuIngredient, 0)
	for _, ingredient := range *req.BaseProduct {
		recipes = append(recipes, &entity.MenuIngredient{
			MenuID:       req.MenuID,
			IngredientID: ingredient.BaseProductID,
			Quantity:     ingredient.Quantity,
		})
	}

	got, err := uc.menuRepo.InsertMenuIngredient(ctx, tx, recipes)
	if err != nil {
		return nil, err
	}

	menuIngredients = append(menuIngredients, got...)

	estimateServings, err := uc.calculateEstimateServings(menuIngredients)
	if err != nil {
		return nil, err
	}

	baseProducts := make([]*menudto.BaseProduct, 0)
	for _, ingredient := range menuIngredients {
		baseProducts = append(baseProducts, &menudto.BaseProduct{
			ID:                    ingredient.IngredientID,
			Name:                  ingredient.IngredientName,
			QuantityPerMenu:       ingredient.Quantity,
			TotalQuantityForServe: ingredient.Quantity * float64(estimateServings) * ingredient.ConversionRate,
			RecentStock:           ingredient.RecentStock,
		})
	}

	activityDetails := make([][]*entity.ActivityDetail, 0)
	for _, ingredient := range *req.BaseProduct {
		quantityStr := strconv.FormatFloat(ingredient.Quantity, 'f', 6, 64)
		activityDetails = append(activityDetails, []*entity.ActivityDetail{
			{
				ChangedID: &ingredient.BaseProductID,
				OldValue:  nil,
				NewValue:  &quantityStr,
			},
		})
	}

	uc.activityRepo.BulkInsertActivity(ctx, tx, []*entity.Activity{
		{
			UserID:          ctx.Value("userID").(int64),
			ActivityFieldID: int64(enums.ActivityMenuIngredient),
			Action:          string(enums.ActivityActionCreate),
		},
	}, activityDetails)

	return &menudto.CreateRecipeMenuResponse{
		MenuID:            req.MenuID,
		EstimatedServings: estimateServings,
		SyncAt:            time.Now().UnixMilli(),
		Recipes:           baseProducts,
	}, nil
}

func (uc *MenuUCImpl) CreateBaseProduct(ctx context.Context, req *menudto.CreateBaseProductRequest) (*menudto.CreateBaseProductResponse, error) {
	err := validation.CreateBaseProductValidation(req)
	if err != nil {
		return nil, err
	}

	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error begin transaction")
		return nil, err
	}

	defer func() {
		if defErr := dbutils.CommitOrRollback(tx, err); defErr != nil {
			err = defErr
		}
	}()
	unitID := req.UnitID
	if req.CustomUnit != nil {
		uitIds, err := uc.menuRepo.InsertUnits(ctx, tx, []*entity.Unit{
			{
				IsInternational: false,
				Name:            req.CustomUnit.Name,
				ConversionRate:  req.CustomUnit.ConversionRate,
				ConversionID:    &req.CustomUnit.ConversionID,
				FranchiseID:     ctx.Value("franchiseID").(int64),
			},
		})
		if err != nil {
			return nil, err
		}
		unitID = uitIds[0]
	}

	ingredientIds, err := uc.menuRepo.InsertIngredients(ctx, tx, []*entity.Ingredient{
		{
			Name:         req.BaseProduceName,
			MinimumStock: req.RequiredStock,
			RecentStock:  req.RecentStock,
			Price:        req.Price,
			UnitID:       unitID,
		},
	})

	if err != nil {
		return nil, err
	}

	uc.activityRepo.BulkInsertActivity(ctx, tx, []*entity.Activity{
		{
			UserID:          ctx.Value("userID").(int64),
			ActivityFieldID: int64(enums.ActivityIngredient),
			Action:          string(enums.ActivityActionCreate),
		},
	}, [][]*entity.ActivityDetail{
		{
			{
				ChangedID: &ingredientIds[0],
				OldValue:  nil,
				NewValue:  &req.BaseProduceName,
			},
		},
	})

	return &menudto.CreateBaseProductResponse{
		ID:           ingredientIds[0],
		Name:         req.BaseProduceName,
		MinimumStock: req.RequiredStock,
		RecentStock:  req.RecentStock,
		Price:        req.Price,
	}, nil
}

func (uc *MenuUCImpl) CreateMenuCategory(ctx context.Context, req *menudto.CreateMenuCategoryRequest) (*menudto.CreateMenuCategoryResponse, error) {
	err := validation.CreateMenuCategoryRequestValidation(req)
	if err != nil {
		return nil, err
	}

	tx, err := uc.db.Begin()
	if err != nil {
		uc.logger.Error().Caller().Err(err).Msg("Error begin transaction")
		return nil, err
	}

	defer func() {
		if defErr := dbutils.CommitOrRollback(tx, err); defErr != nil {
			err = defErr
		}
	}()

	menuCategory, err := uc.menuRepo.InsertCategory(ctx, tx, &entity.Category{
		Name:        req.Name,
		Color:       req.Color,
		FranchiseID: ctx.Value("franchiseID").(int64),
	})
	if err != nil {
		return nil, err
	}

	uc.activityRepo.BulkInsertActivity(ctx, tx, []*entity.Activity{
		{
			UserID:          ctx.Value("userID").(int64),
			ActivityFieldID: int64(enums.ActivityCategory),
			Action:          string(enums.ActivityActionCreate),
		},
	}, [][]*entity.ActivityDetail{
		{
			{
				ChangedID: &menuCategory.ID,
				OldValue:  nil,
				NewValue:  &req.Name,
			},
		},
	})

	return &menudto.CreateMenuCategoryResponse{
		ID:    menuCategory.ID,
		Name:  menuCategory.Name,
		Color: menuCategory.Color,
	}, nil
}

func (uc *MenuUCImpl) calculateEstimateServings(ingredients []*queryresult.MenuIngredientResult) (int64, error) {
	maxServings := int64(0)

	for _, ingredient := range ingredients {
		estimateServings := int64(ingredient.RecentStock / (ingredient.Quantity * ingredient.ConversionRate))
		if maxServings == 0 || estimateServings < maxServings {
			maxServings = estimateServings
		}
	}

	return maxServings, nil
}

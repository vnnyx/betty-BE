package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto/menu"
	"github.com/vnnyx/betty-BE/internal/errors"
)

func CreateMenuRequestValidation(req *menu.CreateMenuRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			ve := make([]*errors.ValidationError, len(validationErrors))
			for i, v := range validationErrors {
				ve[i] = errors.NewValidationError("Validation failed for field", v.Field(), v.Tag())
			}
			return errors.NewValidationErrorWithDetails(ve)
		}
		return errors.NewValidationError(err.Error(), "", "")
	}
	return nil
}

func CreateBaseProductValidation(req *menu.CreateBaseProductRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			ve := make([]*errors.ValidationError, len(validationErrors))
			for i, v := range validationErrors {
				ve[i] = errors.NewValidationError("Validation failed for field", v.Field(), v.Tag())
			}
			return errors.NewValidationErrorWithDetails(ve)
		}
		return errors.NewValidationError(err.Error(), "", "")
	}
	return nil
}

func CreateMenuRecipeRequestValidation(req *menu.CreateMenuRecipeRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			ve := make([]*errors.ValidationError, len(validationErrors))
			for i, v := range validationErrors {
				ve[i] = errors.NewValidationError("Validation failed for field", v.Field(), v.Tag())
			}
			return errors.NewValidationErrorWithDetails(ve)
		}
		return errors.NewValidationError(err.Error(), "", "")
	}

	if req.BaseProduct != nil {
		for _, ingredient := range *req.BaseProduct {
			err := validate.Struct(ingredient)
			if err != nil {
				if validationErrors, ok := err.(validator.ValidationErrors); ok {
					ve := make([]*errors.ValidationError, len(validationErrors))
					for i, v := range validationErrors {
						ve[i] = errors.NewValidationError("Validation failed for field", v.Field(), v.Tag())
					}
					return errors.NewValidationErrorWithDetails(ve)
				}
			}
		}
	}
	return nil
}

func CreateMenuCategoryRequestValidation(req *menu.CreateMenuCategoryRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			ve := make([]*errors.ValidationError, len(validationErrors))
			for i, v := range validationErrors {
				ve[i] = errors.NewValidationError("Validation failed for field", v.Field(), v.Tag())
			}
			return errors.NewValidationErrorWithDetails(ve)
		}
		return errors.NewValidationError(err.Error(), "", "")
	}
	return nil
}

package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto/user"
	"github.com/vnnyx/betty-BE/internal/errors"
)

func CreateOwnerRequestValidation(req *user.CreateOwnerRequest) error {
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

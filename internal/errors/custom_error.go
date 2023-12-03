package errors

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto"
	"github.com/vnnyx/betty-BE/internal/enums"
)

func CustomErrorHandler(c *fiber.Ctx, err error) error {
	_, ok := err.(*ValidationError)
	if ok {
		var obj interface{}
		json.Unmarshal([]byte(err.Error()), &obj)
		return c.Status(fiber.StatusBadRequest).JSON(dto.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: enums.ErrBadRequest,
			Data:   nil,
			Error:  obj,
		})
	}
	customErr, ok := err.(*CustomError)
	if ok {
		var obj interface{}
		json.Unmarshal([]byte(err.Error()), &obj)
		switch customErr.Tag {
		case string(enums.ErrUnauthorized):
			return c.Status(fiber.StatusUnauthorized).JSON(dto.WebResponse{
				Code:   fiber.StatusUnauthorized,
				Status: enums.ErrUnauthorized,
				Data:   nil,
				Error:  obj,
			})
		case string(enums.ErrForbidden):
			return c.Status(fiber.StatusForbidden).JSON(dto.WebResponse{
				Code:   fiber.StatusForbidden,
				Status: enums.ErrForbidden,
				Data:   nil,
				Error:  obj,
			})
		case string(enums.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(dto.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: enums.ErrNotFound,
				Data:   nil,
				Error:  obj,
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.WebResponse{
				Code:   fiber.StatusInternalServerError,
				Status: enums.ErrInternal,
				Data:   nil,
				Error:  obj,
			})
		}
	}
	return c.Status(fiber.StatusInternalServerError).JSON(dto.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   nil,
		Error:  "An unexpected error occurred on the server",
	})
}

type ValidationError struct {
	Message string             `json:"message"`
	Details []*ValidationError `json:"details,omitempty"`
	Field   string             `json:"field,omitempty"`
	Tag     string             `json:"tag,omitempty"`
}

func NewValidationError(message, field, tag string) *ValidationError {
	return &ValidationError{
		Message: message,
		Field:   field,
		Tag:     tag,
	}
}

func NewValidationErrorWithDetails(details []*ValidationError) *ValidationError {
	return &ValidationError{
		Message: "Validation errors",
		Details: details,
	}
}

func (e *ValidationError) Error() string {
	if e.Details != nil {
		b, err := json.Marshal(e)
		if err != nil {
			return "Error marshaling validation errors"
		}
		return string(b)
	}
	return e.Message
}

type CustomError struct {
	Message string `json:"message"`
	Tag     string `json:"-"`
}

func NewCustomError(message, tag string) *CustomError {
	return &CustomError{
		Message: message,
		Tag:     tag,
	}
}

func (e *CustomError) Error() string {
	if e.Tag != "" {
		b, err := json.Marshal(e)
		if err != nil {
			return "Error marshaling custom error"
		}
		return string(b)
	}
	return e.Message
}

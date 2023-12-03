package config

import (
	"github.com/gofiber/fiber/v2"
	custom_errors "github.com/vnnyx/betty-BE/internal/errors"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: custom_errors.CustomErrorHandler,
	}
}

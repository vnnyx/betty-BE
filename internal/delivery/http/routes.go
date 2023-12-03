package http

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/vnnyx/betty-BE/internal/delivery/http/auth"
	"github.com/vnnyx/betty-BE/internal/delivery/http/menu"
	"github.com/vnnyx/betty-BE/internal/delivery/http/user"
	"github.com/vnnyx/betty-BE/internal/middleware/authmiddleware"
)

type Routes struct {
	userHandler    *user.UserHandler
	menuHandler    *menu.MenuHandler
	authHandler    *auth.AuthHandler
	authMiddleware *authmiddleware.AuthMiddleware
}

func NewRoutes(
	userHandler *user.UserHandler,
	menuHandler *menu.MenuHandler,
	authHandler *auth.AuthHandler,
	authMiddleware *authmiddleware.AuthMiddleware,
	route *fiber.App,
) *Routes {
	return &Routes{
		userHandler:    userHandler,
		menuHandler:    menuHandler,
		authHandler:    authHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *Routes) RegisterRoutes(app *fiber.App) {
	api := app.Group("/betty-api")
	user := api.Group("/user")
	auth := api.Group("/auth")
	menu := api.Group("/menu", r.authMiddleware.AuthMiddleware())

	user.Post("/owner", r.userHandler.CreateOwner)
	auth.Post("/login", r.authHandler.Login)
	auth.Post("/refresh-token", r.authHandler.RefreshToken)
	auth.Get("/google", r.authHandler.GoogleSign)
	auth.Get("/google/callback", r.authHandler.GoogleCallback)
	menu.Post("/", r.menuHandler.AddMenu)
	menu.Post("/base-product", r.menuHandler.AddBaseProduct)
	menu.Post("/recipe", r.menuHandler.AddMenuRecipe)
	menu.Post("/category", r.menuHandler.AddMenuCategory)

	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Welcome to Betty API %v", os.Getenv("ENV")),
		})
	})
}

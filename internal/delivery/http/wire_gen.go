// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vnnyx/betty-BE/internal/config"
	auth2 "github.com/vnnyx/betty-BE/internal/delivery/http/auth"
	menu3 "github.com/vnnyx/betty-BE/internal/delivery/http/menu"
	user3 "github.com/vnnyx/betty-BE/internal/delivery/http/user"
	"github.com/vnnyx/betty-BE/internal/infrastructure/db"
	"github.com/vnnyx/betty-BE/internal/infrastructure/repository/activity"
	"github.com/vnnyx/betty-BE/internal/infrastructure/repository/attachmentfile"
	"github.com/vnnyx/betty-BE/internal/infrastructure/repository/company"
	"github.com/vnnyx/betty-BE/internal/infrastructure/repository/menu"
	"github.com/vnnyx/betty-BE/internal/infrastructure/repository/user"
	"github.com/vnnyx/betty-BE/internal/middleware/authmiddleware"
	"github.com/vnnyx/betty-BE/internal/usecase/auth"
	menu2 "github.com/vnnyx/betty-BE/internal/usecase/menu"
	user2 "github.com/vnnyx/betty-BE/internal/usecase/user"
)

// Injectors from routes_injector.go:

func InitializeRoutes(app *fiber.App, cfg *config.Config) (*Routes, error) {
	companyRepository := company.NewCompanyRepository()
	userRepository := user.NewUserRepository()
	activityRepository := activity.NewActivityRepository()
	menuRepository := menu.NewMenuRepository()
	sqlDB, err := db.NewCockroachDatabase(cfg)
	if err != nil {
		return nil, err
	}
	authUC := auth.NewAuthUC(userRepository, activityRepository, sqlDB)
	userUC := user2.NewUserUC(companyRepository, userRepository, activityRepository, menuRepository, authUC, sqlDB)
	userHandler := user3.NewUserHandler(userUC)
	attachmentFileRepository := attachmentfile.NewAttachmentFileRepository()
	menuUC := menu2.NewMenuUC(menuRepository, activityRepository, attachmentFileRepository, sqlDB)
	menuHandler := menu3.NewMenuHandler(menuUC)
	oauth2Config := config.NewGoogleConfig(cfg)
	authHandler := auth2.NewAuthHandler(authUC, oauth2Config)
	authMiddleware := authmiddleware.NewAuthMiddleware(authUC)
	routes := NewRoutes(userHandler, menuHandler, authHandler, authMiddleware, app)
	return routes, nil
}

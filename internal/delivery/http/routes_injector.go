//go:build wireinject
// +build wireinject

package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/vnnyx/betty-BE/internal/config"
	"github.com/vnnyx/betty-BE/internal/delivery/http/auth"
	"github.com/vnnyx/betty-BE/internal/delivery/http/menu"
	"github.com/vnnyx/betty-BE/internal/delivery/http/user"
	"github.com/vnnyx/betty-BE/internal/infrastructure/db"
	activityRepo "github.com/vnnyx/betty-BE/internal/infrastructure/repository/activity"
	afRepo "github.com/vnnyx/betty-BE/internal/infrastructure/repository/attachmentfile"
	companyRepo "github.com/vnnyx/betty-BE/internal/infrastructure/repository/company"
	menuRepo "github.com/vnnyx/betty-BE/internal/infrastructure/repository/menu"
	userRepo "github.com/vnnyx/betty-BE/internal/infrastructure/repository/user"
	"github.com/vnnyx/betty-BE/internal/middleware/authmiddleware"
	authUC "github.com/vnnyx/betty-BE/internal/usecase/auth"
	menuUC "github.com/vnnyx/betty-BE/internal/usecase/menu"
	userUC "github.com/vnnyx/betty-BE/internal/usecase/user"
)

func InitializeRoutes(app *fiber.App, cfg *config.Config) (*Routes, error) {
	wire.Build(
		config.NewGoogleConfig,
		db.NewCockroachDatabase,
		userRepo.NewUserRepository,
		activityRepo.NewActivityRepository,
		companyRepo.NewCompanyRepository,
		menuRepo.NewMenuRepository,
		afRepo.NewAttachmentFileRepository,
		authUC.NewAuthUC,
		userUC.NewUserUC,
		menuUC.NewMenuUC,
		authmiddleware.NewAuthMiddleware,
		auth.NewAuthHandler,
		user.NewUserHandler,
		menu.NewMenuHandler,
		NewRoutes,
	)
	return nil, nil
}

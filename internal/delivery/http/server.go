package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vnnyx/betty-BE/internal/config"
	"github.com/vnnyx/betty-BE/internal/infrastructure/db"
)

func StartHTTPServer() {
	app := fiber.New(config.NewFiberConfig())
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	if db.RunMigration(cfg) != nil {
		panic(err)
	}

	routes, err := InitializeRoutes(app, cfg)
	if err != nil {
		panic(err)
	}

	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":9000"))
}

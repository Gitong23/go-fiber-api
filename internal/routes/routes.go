package routes

import (
	healthzHandler "github.com/Gitong23/go-fiber-hex-api/internal/adapters/http_handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	healthzHandler := healthzHandler.NewHealthzHandler()
	app.Get("/healthz", healthzHandler.HealthCheck)

	apiV1 := app.Group("/api/v1")
	SetupAuthRoutes(apiV1)
}

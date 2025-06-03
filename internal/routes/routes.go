package routes

import (
	healthzHandler "github.com/Gitong23/go-fiber-hex-api/internal/adapters/http_handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	healthzHandler := healthzHandler.NewHealthzHandler()
	app.Get("/healthz", healthzHandler.HealthCheck)

	//set auth middleware
	// This is where you would typically set up authentication middleware
	// For example, you might use JWT or session-based authentication
	// app.Use(middleware.AuthMiddleware)
	// Set up API versioning
	// This groups all API routes under /api/v1

	apiV1 := app.Group("/api/v1")
	SetupAuthRoutes(apiV1)
}

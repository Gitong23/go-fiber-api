package user

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// User routes will be implemented here
	api := app.Group("/api/v1")
	users := api.Group("/users")

	// Protected routes would use auth middleware
	_ = users // placeholder to avoid unused variable error
}

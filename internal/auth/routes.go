package auth

import (
	"github.com/Gitong23/go-fiber-hex-api/config"
	"github.com/Gitong23/go-fiber-hex-api/internal/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userRepo := user.NewUserRepository()
	authService := NewAuthService(userRepo, config.AppConfig.JWT.Secret)
	handler := NewHandler(authService)

	api := app.Group("/api/v1")
	auth := api.Group("/auth")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}

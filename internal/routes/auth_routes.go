package routes

import (
	"github.com/Gitong23/go-fiber-hex-api/config"
	authAdapters "github.com/Gitong23/go-fiber-hex-api/internal/adapters/http_handler"
	userAdapters "github.com/Gitong23/go-fiber-hex-api/internal/adapters/mongo"
	authCore "github.com/Gitong23/go-fiber-hex-api/internal/core/auth"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	userRepo := userAdapters.NewUserRepository()
	authService := authCore.NewAuthService(userRepo, config.AppConfig.JWT.Secret)
	authHandler := authAdapters.NewAuthHandler(authService)

	auth := router.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
}

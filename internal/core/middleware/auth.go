package middleware

import (
	"net/http"
	"strings"

	"github.com/Gitong23/go-fiber-hex-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
)

// AuthMiddleware checks if the user is authenticated
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(AuthorizationHeader)
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid authorization header format",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		// Set user ID in context
		c.Locals("userID", claims.Subject)
		return c.Next()
	}
}

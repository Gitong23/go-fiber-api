package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Process request
	err := c.Next()

	// Log after request
	log.Printf("%s %s - %d - %v",
		c.Method(),
		c.Path(),
		c.Response().StatusCode(),
		time.Since(start),
	)

	return err
}

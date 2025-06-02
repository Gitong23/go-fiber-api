package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Gitong23/go-fiber-hex-api/config"
	"github.com/Gitong23/go-fiber-hex-api/internal/auth"
	"github.com/Gitong23/go-fiber-hex-api/internal/health"
	"github.com/Gitong23/go-fiber-hex-api/internal/middleware"
	"github.com/Gitong23/go-fiber-hex-api/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// Initialize database
	db.Connect()

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Setup middleware
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(middleware.LoggingMiddleware)

	// Setup routes
	health.SetupRoutes(app)
	auth.SetupRoutes(app)

	// Channel to listen for interrupt signal to trigger shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		port := fmt.Sprintf(":%d", cfg.App.Port)
		log.Printf("Server starting on port %s", port)
		if err := app.Listen(port); err != nil {
			log.Printf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	db.Disconnect()

	log.Println("Server exited")
}

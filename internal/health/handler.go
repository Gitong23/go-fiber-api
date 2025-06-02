package health

import (
	"context"
	"time"

	"github.com/Gitong23/go-fiber-hex-api/internal/response"
	"github.com/Gitong23/go-fiber-hex-api/pkg/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	// Check database connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	status := "healthy"
	dbStatus := "connected"

	if err := db.MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		status = "unhealthy"
		dbStatus = "disconnected"
	}

	return c.JSON(response.SuccessResponse("Health check", fiber.Map{
		"status":    status,
		"database":  dbStatus,
		"timestamp": time.Now(),
	}))
}

func SetupRoutes(app *fiber.App) {
	handler := NewHandler()
	app.Get("/health", handler.HealthCheck)
}

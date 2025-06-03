package adapters

import (
	"net/http"

	core "github.com/Gitong23/go-fiber-hex-api/internal/core/auth"
	"github.com/Gitong23/go-fiber-hex-api/internal/core/response"
	"github.com/gofiber/fiber/v2"
)

type IauthHandler struct {
	service core.IauthService
}

func NewAuthHandler(service core.IauthService) *IauthHandler {
	return &IauthHandler{service: service}
}

func (h *IauthHandler) Register(c *fiber.Ctx) error {
	var req core.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse("Invalid request payload"))
	}

	user, err := h.service.Register(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(response.SuccessResponse("User registered successfully", user))
}

func (h *IauthHandler) Login(c *fiber.Ctx) error {
	var req core.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse("Invalid request payload"))
	}

	token, err := h.service.Login(req)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(response.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(response.SuccessResponse("Login successful", token))
}

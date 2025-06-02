package auth

import (
	"net/http"

	"github.com/Gitong23/go-fiber-hex-api/internal/response"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse("Invalid request payload"))
	}

	user, err := h.service.Register(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(response.SuccessResponse("User registered successfully", user))
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse("Invalid request payload"))
	}

	token, err := h.service.Login(req)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(response.ErrorResponse(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(response.SuccessResponse("Login successful", token))
}

type RegisterRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=50"`
	Password  string `json:"password" validate:"required,min=6"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

package user

// import (
// 	"net/http"

// 	"github.com/Gitong23/go-fiber-hex-api/internal/response"
// 	"github.com/gofiber/fiber/v2"
// )

// type Handler struct {
// 	userService service.UserService
// }

// func NewHandler(userService service.UserService) *Handler {
// 	return &Handler{userService: userService}
// }

// func (h *Handler) GetUser(c *fiber.Ctx) error {
// 	userID := c.Params("id")
// 	user, err := h.userService.GetUserByID(userID)
// 	if err != nil {
// 		return response.ErrorResponse(c, http.StatusNotFound, "User not found")
// 	}
// 	return response.SuccessResponse(c, http.StatusOK, user)
// }

// func (h *Handler) CreateUser(c *fiber.Ctx) error {
// 	var userInput service.CreateUserInput
// 	if err := c.BodyParser(&userInput); err != nil {
// 		return response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
// 	}

// 	user, err := h.userService.CreateUser(userInput)
// 	if err != nil {
// 		return response.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
// 	}
// 	return response.SuccessResponse(c, http.StatusCreated, user)
// }

// func (h *Handler) UpdateUser(c *fiber.Ctx) error {
// 	userID := c.Params("id")
// 	var userInput service.UpdateUserInput
// 	if err := c.BodyParser(&userInput); err != nil {
// 		return response.ErrorResponse(c, http.StatusBadRequest, "Invalid input")
// 	}

// 	user, err := h.userService.UpdateUser(userID, userInput)
// 	if err != nil {
// 		return response.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user")
// 	}
// 	return response.SuccessResponse(c, http.StatusOK, user)
// }

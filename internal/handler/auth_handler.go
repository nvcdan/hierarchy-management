package handler

import (
	"hierarchy-management/internal/errors"
	"hierarchy-management/internal/response"
	"hierarchy-management/internal/service"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		response.HandleError(c, errors.NewValidationError("credentials", "Invalid username or password"))
		return
	}

	err := h.userService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": credentials.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		response.HandleError(c, errors.NewInternalError("Failed to generate token"))
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		IsSuccess: true,
		Message:   "Login successful",
		Data:      map[string]string{"token": tokenString},
	})
}

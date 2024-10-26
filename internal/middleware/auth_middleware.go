package middleware

import (
	"hierarchy-management/internal/response"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func abortUnauthorized(c *gin.Context, code, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, response.APIResponse{
		IsSuccess: false,
		Message:   message,
		Error: map[string]string{
			"code":    code,
			"message": message,
		},
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			abortUnauthorized(c, "unauthorized", "No Authorization header provided")
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			abortUnauthorized(c, "invalid_auth_header", "Authorization header must start with 'Bearer '")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != os.Getenv("JWT_SECRET") {
			abortUnauthorized(c, "invalid_token", "The provided token is invalid")
			return
		}

		c.Next()
	}
}

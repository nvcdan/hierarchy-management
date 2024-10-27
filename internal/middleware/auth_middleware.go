package middleware

import (
	"errors"
	"hierarchy-management/internal/response"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			abortUnauthorized(c, "invalid_token", "Invalid or expired token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("username", claims["username"])
		}

		c.Next()
	}
}

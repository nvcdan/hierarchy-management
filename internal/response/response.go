package response

import (
	"hierarchy-management/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Error     interface{} `json:"error,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func HandleError(c *gin.Context, err error) {
	statusCode, errorResponse := mapErrorToResponse(err)
	c.JSON(statusCode, APIResponse{
		IsSuccess: false,
		Message:   "An error occurred",
		Error:     errorResponse,
	})
}

func mapErrorToResponse(err error) (int, interface{}) {
	switch e := err.(type) {
	case *errors.InternalError:
		return http.StatusInternalServerError, ErrorResponse{
			Code:    "internal_error",
			Message: e.Message,
		}
	case *errors.NotFoundError:
		return http.StatusNotFound, ErrorResponse{
			Code:    "not_found",
			Message: e.Error(),
		}
	case *errors.ValidationError:
		return http.StatusBadRequest, []ErrorResponse{
			{
				Code:    "validation_error",
				Message: e.Error(),
			},
		}
	case *errors.DuplicateEntryError:
		return http.StatusConflict, ErrorResponse{
			Code:    "duplicate_entry",
			Message: e.Error(),
		}
	case *errors.AuthenticationError:
		return http.StatusUnauthorized, ErrorResponse{
			Code:    "authentication_failed",
			Message: e.Message,
		}
	default:
		return http.StatusInternalServerError, ErrorResponse{
			Code:    "unknown_error",
			Message: "An unexpected error occurred",
		}
	}
}

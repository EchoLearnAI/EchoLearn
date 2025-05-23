package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError represents a standard error response format.
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// NewError creates a new APIError and sends it as a JSON response.
func NewError(c *gin.Context, code int, message string) {
	c.JSON(code, APIError{Code: code, Message: message})
}

// NewErrorWithDetails creates a new APIError with additional details.
func NewErrorWithDetails(c *gin.Context, code int, message string, details string) {
	c.JSON(code, APIError{Code: code, Message: message, Details: details})
}

// Common error responses
func BadRequest(c *gin.Context, details ...string) {
	msg := "Bad Request"
	detailStr := ""
	if len(details) > 0 {
		detailStr = details[0]
	}
	NewErrorWithDetails(c, http.StatusBadRequest, msg, detailStr)
}

func NotFound(c *gin.Context, resource ...string) {
	msg := "Resource not found"
	if len(resource) > 0 {
		msg = resource[0] + " not found"
	}
	NewError(c, http.StatusNotFound, msg)
}

func InternalServerError(c *gin.Context, details ...string) {
	msg := "Internal Server Error"
	detailStr := ""
	if len(details) > 0 {
		detailStr = details[0]
	}
	NewErrorWithDetails(c, http.StatusInternalServerError, msg, detailStr)
}

func Unauthorized(c *gin.Context, details ...string) {
    msg := "Unauthorized"
    detailStr := ""
    if len(details) > 0 {
        detailStr = details[0]
    }
    NewErrorWithDetails(c, http.StatusUnauthorized, msg, detailStr)
}

func Forbidden(c *gin.Context, details ...string) {
    msg := "Forbidden"
    detailStr := ""
    if len(details) > 0 {
        detailStr = details[0]
    }
    NewErrorWithDetails(c, http.StatusForbidden, msg, detailStr)
} 
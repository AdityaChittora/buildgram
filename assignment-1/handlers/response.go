// Package handlers provides the HTTP handler functions for the BuildGram API.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// successResponse sends a standard success JSON envelope.
func successResponse(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
	})
}

// errorResponse sends a standard error JSON envelope.
func errorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
	})
}

// notFound is a convenience wrapper for 404 error responses.
func notFound(c *gin.Context, message string) {
	errorResponse(c, http.StatusNotFound, message)
}

// badRequest is a convenience wrapper for 400 error responses.
func badRequest(c *gin.Context, message string) {
	errorResponse(c, http.StatusBadRequest, message)
}

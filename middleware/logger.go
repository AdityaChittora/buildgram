// Package middleware provides custom Gin middleware for BuildGram.
package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger returns a Gin middleware that logs the HTTP method, path, and
// latency for every incoming request to stdout in the BuildGram log format:
//
//	[BuildGram] POST /api/v1/users | 204.75µs
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Let the request be handled by subsequent middleware / the route handler.
		c.Next()

		latency := time.Since(start)
		fmt.Printf("[BuildGram] %s %s | %v\n", c.Request.Method, c.Request.URL.Path, latency)
	}
}

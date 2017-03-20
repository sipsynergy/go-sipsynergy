package middleware

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// AuthRequired will check the request headers for the given API Key.
// If it is not found the request will be cancelled and a 401 will be returned.
func AuthRequired(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-API-KEY") != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Not Authorized",
			})
			return
		}

		c.Next()
	}
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			status := c.Writer.Status()
			if status < 400 {
				status = http.StatusInternalServerError
			}
			c.JSON(status, gin.H{"error": err.Error()})
		}
	}
}

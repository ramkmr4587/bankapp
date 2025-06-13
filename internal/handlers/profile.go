package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	uidAny, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
		return
	}
	userID, ok := uidAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id has wrong type"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}

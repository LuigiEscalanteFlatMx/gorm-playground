package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthyCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"system": "ping",
	})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"system": "ping",
	})
}

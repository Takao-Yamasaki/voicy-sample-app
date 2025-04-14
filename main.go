package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	apiGroup := r.Group("/api")
	apiGroup.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "register endpoint",
		})
	})

	r.Run(":8080")
}

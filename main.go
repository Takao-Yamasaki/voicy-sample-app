package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"voicy-sample-app/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	apiGroup := r.Group("/api")
	apiGroup.POST("/register", controllers.Register)

	r.Run(":8080")
}

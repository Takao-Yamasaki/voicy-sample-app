package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"voicy-sample-app/handlers"
)

func main() {
	r := gin.Default()

	// healthcheck用のエンドポイント
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	apiGroup := r.Group("/api")
	apiGroup.GET("/sounds", handlers.GetSounds)

	r.Run(":8080")
}

package main

import (
	"voicy-sample-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// healthcheck用のエンドポイント
	r.GET("/healthcheck", handlers.HealthCheck)

	apiGroup := r.Group("/api")
	apiGroup.GET("/sounds", handlers.GetSounds)

	r.Run(":8080")
}

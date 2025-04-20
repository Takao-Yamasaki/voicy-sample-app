package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Sound struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Personality string `json:"personality"`
	CreatedAt   string `json:"created_at"`
}

var sounds []Sound

func main() {
	r := gin.Default()

	// healthcheck用のエンドポイント
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	apiGroup := r.Group("/api")
	apiGroup.GET("/sounds", getSounds)

	r.Run(":8080")
}

func getSounds(c *gin.Context) {
	// JSONファイルを開く
	file, err := os.Open("data.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to open data file"})
		return
	}
	defer file.Close()

	// JSONデータをデコードする
	if err := json.NewDecoder(file).Decode(&sounds); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to decode data"})
		return
	}

	c.JSON(http.StatusOK, sounds)
}

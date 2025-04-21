package handlers

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

func GetSounds(c *gin.Context) {
	// JSONファイルを開く
	file, err := os.Open("data/data.json")
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

// TODO:　単体テストを実装

package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"voicy-sample-app/models"

	"github.com/gin-gonic/gin"
)

// ヘルスチェック用の関数
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

var sounds []models.Sound

// 音声データの一覧を取得する関数
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

package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"voicy-sample-app/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// HealtCheck関数のテスト
func TestHealthCheck(t *testing.T) {
	// Ginのルーターをセットアップ
	r := gin.Default()
	r.GET("/healthcheck", handlers.HealthCheck)

	tests := []struct {
		name         string
		httpMethod   string
		url          string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "ヘルスチェック成功",
			httpMethod:   "GET",
			url:          "/healthcheck",
			expectedCode: http.StatusOK,
			expectedBody: `{"status":"healthy"}`,
		},
		{
			name:         "ヘルスチェック失敗（不正なURL）",
			httpMethod:   "GET",
			url:          "/invalid",
			expectedCode: http.StatusNotFound,
			expectedBody: "404 page not found",
		},
		{
			name:         "ヘルスチェック失敗（メソッド不許可）",
			httpMethod:   "POST",
			url:          "/healthcheck",
			expectedCode: http.StatusNotFound,
			expectedBody: "404 page not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のHTTPリクエストを作成
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.httpMethod, tt.url, nil)
			// リクエストを処理
			r.ServeHTTP(w, req)
			// ステータスコードの確認
			assert.Equal(t, tt.expectedCode, w.Code)
			// レスポンスボディの確認
			assert.Equal(t, tt.expectedBody, w.Body.String())
		})
	}
}

// GetSounds関数のテスト
func TestGetSounds(t *testing.T) {
	// テスト用にGinのルーターを設定
	router := gin.Default()
	router.GET("/api/sounds", handlers.GetSounds)

	tests := []struct {
		name           string
		httpMethod     string
		url            string
		expectedCode   int
		expectedSounds int
	}{
		// {
		// 	name:           "正常系（データが2件存在する場合）",
		// 	httpMethod:     "GET",
		// 	url:            "/api/sounds",
		// 	expectedCode:   http.StatusOK,
		// 	expectedSounds: 2,
		// },
		{
			name:           "異常系（不正なURL）",
			httpMethod:     "GET",
			url:            "/api/invalid",
			expectedCode:   http.StatusNotFound,
			expectedSounds: 0,
		},
		{
			name:           "異常系（不正なメソッド）",
			httpMethod:     "POST",
			url:            "/api/sounds",
			expectedCode:   http.StatusNotFound,
			expectedSounds: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のリクエストを作成
			req, _ := http.NewRequest(tt.httpMethod, tt.url, nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			// ステータスコードの確認
			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}

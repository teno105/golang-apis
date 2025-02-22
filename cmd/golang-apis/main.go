package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"golang-apis/infra"
	"golang-apis/internal/models"
)

// 특정 게임 ID에 해당하는 데이터를 로드하는 함수
func loadGameData(gameID string) (*GameData, error) {
	maintenancePath := filepath.Join("data", gameID, "maintenance.json")
	noticePath := filepath.Join("data", gameID, "notice.json")

	var gameData GameData

	// maintenance.json 읽기
	if err := loadJSONFile(maintenancePath, &gameData.Maintenance); err != nil {
		return nil, err
	}

	// notice.json 읽기
	if err := loadJSONFile(noticePath, &gameData.Notice); err != nil {
		return nil, err
	}

	return &gameData, nil
}

func main() {

	r := gin.Default()

	// GET /v2/init_data/games/:id 엔드포인트
	r.GET("/v2/init_data/games/:id", func(c *gin.Context) {
		gameID := c.Param("id") // URL에서 game ID 가져오기
		data, err := loadGameData(gameID) // 해당 ID의 데이터를 로드
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game data not found"})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	// 서버 실행
	fmt.Println("Server is running on port 8080...")
	r.Run(":8080")
}

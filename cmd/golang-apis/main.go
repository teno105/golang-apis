package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"golang-apis/infra"
	"golang-apis/internal/models"
)

func main() {

	// Initialize Gin router
	r := gin.Default()

	// GET /v2/init_data/games/:id 엔드포인트
	r.GET("/v2/init_data/games/:id", func(c *gin.Context) {
		gameID := c.Param("id") // URL에서 game ID 가져오기
		data, err := loadNotice(gameID) // 해당 ID의 gamedata.json 로드
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game data not found"})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	fmt.Println("Server is running on port 8080")
	r.Run(":8080")
}

package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"golang-apis/infra"
	"golang-apis/internal/models"
)

func fetchMaintenances(db *gorm.DB, wg *sync.WaitGroup, maintenances *[]models.Maintenance) {
	defer wg.Done()
	result := db.Find(maintenances)
	if result.Error != nil {
		fmt.Println("Error reading from Maintenance table:", result.Error)
	}
}

func createMaintenance(db *gorm.DB, c *gin.Context) {

	var maintenance models.Maintenance
	if err := c.ShouldBindJSON(&maintenance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := infra.DB.Create(&maintenance)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, maintenance)
}

func main() {

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Migrate the schema
	db.AutoMigrate(&models.Maintenance{})

	// Initialize Gin router
	r := gin.Default()

	// Define GET endpoint to fetch Maintenance records
	r.GET("/maintenances", func(c *gin.Context) {
		var maintenances []models.Maintenance
		var wg sync.WaitGroup
		wg.Add(1)
		go fetchMaintenances(db, &wg, &maintenances)
		wg.Wait()

		c.JSON(http.StatusOK, maintenances)
	})

	// GET /v2/init_data/games/:id 엔드포인트
	r.GET("/v2/init_data/games/:id", func(c *gin.Context) {
		gameID := c.Param("id") // URL에서 game ID 가져오기
		data, err := loadGameData("game_id" + gameID) // game_id + id 형태로 파일명 생성
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Game data not found"})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	fmt.Println("Server is running on port 8080")
	r.Run(":8080")
}

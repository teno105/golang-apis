// cmd/golang-apis/main.go
package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Maintenance represents the maintenance table in the database
type Maintenance struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"size:255"`
	Body        string `gorm:"type:text"`
	DetailedUrl string `gorm:"size:255"`
}

func printHelloWorld() {
	fmt.Println("Hello, World!")
}

func fetchMaintenances(db *gorm.DB, wg *sync.WaitGroup, maintenances *[]Maintenance) {
	defer wg.Done()
	result := db.Find(maintenances)
	if result.Error != nil {
		fmt.Println("Error reading from Maintenance table:", result.Error)
	}
}

func main() {
	printHelloWorld()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Migrate the schema
	db.AutoMigrate(&Maintenance{})

	// Initialize Gin router
	r := gin.Default()

	// Define GET endpoint to fetch Maintenance records
	r.GET("/maintenances", func(c *gin.Context) {
		var maintenances []Maintenance
		var wg sync.WaitGroup
		wg.Add(1)
		go fetchMaintenances(db, &wg, &maintenances)
		wg.Wait()

		c.JSON(http.StatusOK, maintenances)
	})

	// Start the Gin server
	r.Run(":8080")
}

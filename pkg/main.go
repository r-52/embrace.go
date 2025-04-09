package main

import (
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/r-52/embrace/models"
)

func main() {
	prepareEnv()

	// Initialize the database
	// and run the migrations
	models.OpenDatabase()

	router := gin.Default()

	router.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"backend":   "v0.0.0",
			"startedAt": time.Now(),
		})
	})

	router.Run()

}

func prepareEnv() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Error getting current working directory")
	}
	err = godotenv.Load(path.Join(cwd, "..", ".env"))
	if err != nil {
		panic("Error loading .env file")
	}
}

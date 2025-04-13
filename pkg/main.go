package main

import (
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/r-52/embrace/models"
	"github.com/r-52/embrace/models/dto/user"
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

	apiV1 := router.Group("/api/v1")
	setupUserRoutes(apiV1)
	setupCompanyRoutes(apiV1)

	router.Run()

}

func setupCompanyRoutes(apiV1 *gin.RouterGroup) {
	companies := apiV1.Group("/companies")
	companies.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Company created",
		})
	})
}

func setupUserRoutes(apiV1 *gin.RouterGroup) {
	users := apiV1.Group("/users")
	users.POST("/create", func(c *gin.Context) {
		var req user.CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created",
		})
	})
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

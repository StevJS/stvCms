package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"stvCms/internal/config"
)
//prueba commit 
func main() {
	loadEnv()
	//startDatabase()
	startServer()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func startServer() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"response": "pong!"})
	})
	// post group
	postGroup := router.Group("/post")
	postGroup.GET("")

	// users group
	//userGroup := router.Group("/user")
	//userGroup.GET("")

	err := router.Run(":" + os.Getenv("GIN_PORT"))
	if err != nil {
		return
	}
}

func startDatabase() {
	config.Init()
}

package main

import (
	"log"

	"github.com/Kim-DaeHan/my-gin-app/config"
	"github.com/joho/godotenv"

	"github.com/Kim-DaeHan/my-gin-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// .env 파일 로딩
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()

	//run database
	config.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:8080")
}

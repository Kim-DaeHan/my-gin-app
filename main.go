package main

import (
	"context"
	"log"

	"github.com/Kim-DaeHan/my-gin-app/config"
	"github.com/Kim-DaeHan/my-gin-app/routes"
	"github.com/joho/godotenv"

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
	client := config.InitializeDB()
	defer client.Disconnect(context.Background()) // 애플리케이션이 종료되기 전에 연결 닫기
	routes.UserRoute(router)

	router.Run("localhost:8080")
}

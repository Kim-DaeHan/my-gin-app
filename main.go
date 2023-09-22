// main.go

package main

import (
	"github.com/Kim-DaeHan/my-gin-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin 라우터 생성
	router := gin.Default()

	// 라우트 설정
	routes.SetupRoutes(router)

	// 서버 시작
	router.Run(":8080")
}

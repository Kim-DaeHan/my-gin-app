package main

import (
	"github.com/Kim-DaeHan/my-gin-app/config"

	"github.com/Kim-DaeHan/my-gin-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	config.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:8080")
}

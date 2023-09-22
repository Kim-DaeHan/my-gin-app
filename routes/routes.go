// routes/routes.go

package routes

import (
	"github.com/Kim-DaeHan/my-gin-app/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes는 애플리케이션의 라우팅을 설정합니다.
func SetupRoutes(router *gin.Engine) {
	// 사용자 관련 라우트 설정
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		// 다른 사용자 관련 라우트 설정
	}

	// 다른 라우트 그룹 및 핸들러 함수를 추가할 수 있습니다.
}

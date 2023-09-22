package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kim-DaeHan/my-gin-app/models"
)

// GetUsers는 모든 사용자를 가져와 응답합니다.
func GetUsers(c *gin.Context) {
	users, err := models.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// 다른 컨트롤러 함수들을 여기에 추가할 수 있습니다.

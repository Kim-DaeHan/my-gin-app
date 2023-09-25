// tests/user_controller_test.go

package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Kim-DaeHan/my-gin-app/controllers" // 테스트할 컨트롤러 임포트

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// Gin 엔진을 모의(Mock) 모드로 설정
	gin.SetMode(gin.TestMode)

	// 라우터 생성
	router := gin.New()
	router.GET("/users", controllers.GetAllUser) // 테스트할 컨트롤러 사용

	// GET /users 요청을 수행
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// 상태 코드 확인 (200 OK 예상)
	assert.Equal(t, http.StatusOK, resp.Code)

	// 응답 본문을 JSON으로 파싱하고 확인 (사용자 목록 예상)
	expected := `[]` // 예시: 빈 사용자 목록
	assert.JSONEq(t, expected, resp.Body.String())
}

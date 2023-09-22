package models

// User는 사용자 모델을 나타냅니다.
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// 다른 모델 및 데이터베이스 관련 로직을 여기에 추가할 수 있습니다.

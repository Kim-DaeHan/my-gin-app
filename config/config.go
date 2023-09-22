package config

import (
	"os"
)

// Config는 애플리케이션의 설정을 저장하는 구조체입니다.
type Config struct {
	MongoDBURI  string
	MongoDBName string
	// 다른 설정 필드들을 추가할 수 있습니다.
}

// LoadConfig는 환경 변수를 읽고 설정을 반환합니다.
func LoadConfig() *Config {
	return &Config{
		MongoDBURI:  os.Getenv("MONGODB_URI"),  // MongoDB 연결 문자열
		MongoDBName: os.Getenv("MONGODB_NAME"), // MongoDB 데이터베이스 이름
		// 다른 설정 필드들을 여기에 추가할 수 있습니다.
	}
}

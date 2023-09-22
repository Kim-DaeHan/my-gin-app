package config

// AppConfig는 애플리케이션의 설정을 저장하는 구조체입니다.
type AppConfig struct {
	// 애플리케이션 설정 필드들을 정의합니다.
}

// LoadConfig는 설정을 로드하는 함수입니다.
func LoadConfig() *AppConfig {
	// 설정을 읽어와 AppConfig 인스턴스를 반환합니다.
	return &AppConfig{}
}

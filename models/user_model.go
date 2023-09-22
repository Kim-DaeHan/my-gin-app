package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User는 사용자 모델을 나타냅니다.
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GetUserCollection는 MongoDB의 "users" 컬렉션을 반환합니다.
func GetUserCollection() (*mongo.Collection, error) {
	// MongoDB 연결 설정
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// MongoDB 컬렉션 가져오기
	collection := client.Database("mydb").Collection("users")
	return collection, nil
}

// FindAllUsers는 모든 사용자를 조회합니다.
func FindAllUsers() ([]User, error) {
	collection, err := GetUserCollection()
	if err != nil {
		return nil, err
	}

	// MongoDB에서 사용자 조회
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []User
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

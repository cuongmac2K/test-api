package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

func ConnectToMongoDB(username string, password string) {
	// Tạo thông tin đăng nhập
	credential := options.Credential{
		Username: username,
		Password: password,
	}

	// Tạo tùy chọn kết nối với thông tin đăng nhập
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Kiểm tra kết nối
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	MongoClient = client
	fmt.Println("Connected to MongoDB!")

}

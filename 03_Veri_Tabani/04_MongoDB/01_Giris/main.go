package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Client ayarları oluşturma
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// MongoDB'ye bağlantı
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Bağlantı kontrol
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB'ye Bağlanıldı!")
}

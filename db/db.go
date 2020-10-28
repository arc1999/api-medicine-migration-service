package db

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var mongodb *mongo.Database

//Init -
func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		os.Getenv("DATA_MONGODB_URI"),
	))

	if err != nil {
		log.Fatal(err)
	}
	mongodb = client.Database(os.Getenv("DATA_MONGODB_DATABASE"))
	log.Println("MongoDb Connected")
}

//GetDB - return Db instance
func GetMongoDB() *mongo.Database {
	return mongodb
}

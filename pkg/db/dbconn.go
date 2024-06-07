package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dipankarupd/writersparadise/pkg/config"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName string
var Client *mongo.Client

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	connectionString := os.Getenv("MONGO_URI")
	dbName = config.DB_NAME

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongodb successfully")

	Client = client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection instance ready")
	return collection
}

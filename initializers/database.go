package initializers

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToDB() {
	// load env variables from .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	LoadEnvVariables()

	// get value from .env
	// MONGO_URI := os.Getenv("MONGO_URI")

	// connect to the database
	clientOptions := options.Client().ApplyURI("mongodb://htetko:secret@db:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	fmt.Println("Connected to MongoDB")

	//client.Database(dbName).Collection(colName)

}

var RedisClient *redis.Client

func ConnectToRedis() {
	// Connect to Redis
	opt, err := redis.ParseURL("redis://redis:6379")

	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	RedisClient = client
	fmt.Println("Connected to Redis")
}

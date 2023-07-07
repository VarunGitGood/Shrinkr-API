package database

import (
	"api/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	mdb    *mongo.Database
)

func ConnectMongo() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.Config("MONGO_URL")).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		fmt.Println(err)
	}
	mdb = client.Database("shrinkr")
	fmt.Println("Connected to MongoDB")
}

func RegisterUser(user *User) error {
	collection := mdb.Collection("users")
	// check if user already exists
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("User registered")
	return nil
}

func GetUser(email string) (*User, error) {
	collection := mdb.Collection("users")
	filter := bson.D{{"username", email}}
	var user User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

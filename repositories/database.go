package repositories

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDb() *mongo.Database {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kali:kali123@cluster-fulls-tech-kzjnc.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)

	db := client.Database("fullsTechDB")
	return db

}

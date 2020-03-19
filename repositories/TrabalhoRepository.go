package repositories

import (
	"context"
	"fullstech-backend-go/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TrabalhoRepository struct{}

func (TrabalhoRepository) FindAllTrabalhos() []models.Trabalho {
	trabalhos := []models.Trabalho{}
	db := connectDb()

	findOptions := options.Find()
	findOptions.SetLimit(2)

	cur, err := db.Collection("trabalhos").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem models.Trabalho
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		trabalhos = append(trabalhos, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return trabalhos
}

func (TrabalhoRepository) Save(trabalho models.Trabalho) interface{} {
	db := connectDb()

	insertedResult, err := db.Collection("trabalhos").InsertOne(context.TODO(), trabalho)
	if err != nil {
		panic(err)
	}

	return insertedResult.InsertedID
}

func (TrabalhoRepository) Delete(id string) interface{} {

	db := connectDb()

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalln(err)
	}

	filter := bson.M{"_id": idPrimitive}

	deleteResult, err := db.Collection("trabalhos").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return deleteResult
}

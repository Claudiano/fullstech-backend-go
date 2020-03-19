package repositories

import (
	"context"
	"fullstech-backend-go/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LembreteRepository struct{}

func (LembreteRepository) FindAllLembretes() []models.Lembrete {
	lembretes := []models.Lembrete{}
	db := connectDb()

	findOptions := options.Find()
	findOptions.SetLimit(2)

	cur, err := db.Collection("lembretes").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem models.Lembrete
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		lembretes = append(lembretes, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return lembretes
}

func (LembreteRepository) Save(lembrete models.Lembrete) interface{} {
	db := connectDb()

	insertedResult, err := db.Collection("lembretes").InsertOne(context.TODO(), lembrete)
	if err != nil {
		panic(err)
	}

	return insertedResult.InsertedID
}

func (LembreteRepository) Delete(id string) interface{} {
	db := connectDb()

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatalln(err)
	}

	filter := bson.M{"_id": idPrimitive}

	deleteResult, err := db.Collection("lembretes").DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return deleteResult
}

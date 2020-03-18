package repositories

import (
	"context"
	"fmt"
	"fullstech-backend-go/models"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TrabalhoRepository struct{}

func (TrabalhoRepository) FindAllTrabalhos() []models.Trabalho {
	trabalhos := []models.Trabalho{}
	db := connectDb()

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := db.Collection("trabalhos").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
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

	// Close the cursor once finished
	cur.Close(context.TODO())

	return trabalhos
}

func (TrabalhoRepository) Save(trabalho models.Trabalho) string {
	db := connectDb()

	insertedResult, err := db.Collection("trabalhos").InsertOne(context.TODO(), trabalho)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintln("Usuario salvo com id: ", insertedResult.InsertedID)
}

func (TrabalhoRepository) Delete(id int64) string {
	return "exclui um trabalho, com  id = " + strconv.FormatInt(id, 10)
}

package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertJSON(client *mongo.Client, dbName, collName string, data interface{}) error {
	db := client.Database(dbName)
	coll := db.Collection(collName)

	// Vérifier si les données sont dans le bon format
	bsonData, err := bson.Marshal(data)
	if err != nil {
		return err
	}

	_, err = coll.InsertOne(context.Background(), bsonData)
	if err != nil {
		return err
	}
	return nil
}

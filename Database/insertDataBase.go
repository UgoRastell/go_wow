package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func InsertJSON(databaseName string, collectionName string, document interface{}) error {
    // Create a context object for the operation
    ctx := context.TODO()

    // Get a handle to the database
    db := client.Database(databaseName)

    // Get a handle to the collection
    coll := db.Collection(collectionName)

    // Insert the document into the collection
    _, err := coll.InsertOne(ctx, document)
    if err != nil {
        return err
    }

    return nil
}


package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
)

func InsertJSON(client *mongo.Client, dbName, collName string, data interface{}) error {
    coll := client.Database(dbName).Collection(collName)
    _, err := coll.InsertOne(context.Background(), data)
    if err != nil {
        return err
    }
    return nil
}

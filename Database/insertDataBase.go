package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "fmt"
)

func InsertDocument(client *mongo.Client, dbName string, collectionName string, document interface{}) error {
    // Insérer le document dans la collection
    collection := client.Database(dbName).Collection(collectionName)
    _, err := collection.InsertOne(context.Background(), document)
    if err != nil {
        return fmt.Errorf("Erreur lors de l'insertion du document dans la collection : %v", err)
    }

    return nil
}

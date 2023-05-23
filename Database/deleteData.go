package db

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func DeleteDocuments(client *mongo.Client, dbName string, collectionName string, filter bson.M) error {
    // Sélectionner la base de données et la collection
    collection := client.Database(dbName).Collection(collectionName)

    // Supprimer les documents correspondants au filtre spécifié
    deleteResult, err := collection.DeleteMany(context.Background(), filter)
    if err != nil {
        return fmt.Errorf("Erreur lors de la suppression des documents : %v", err)
    }

    fmt.Printf("Nombre de documents supprimés : %d\n", deleteResult.DeletedCount)

    return nil
}

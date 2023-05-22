package db

import (
    "context"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func CheckUserExists(client *mongo.Client, guildID string, userID string, username string) (bool, error) {
	// Get the collection from the database
	collection := client.Database("gowow").Collection("users")

	// Set the filter to search for the user
	filter := bson.M{"guild_id": guildID, "user_id": userID}

	// Set options to limit the result to one document
	options := options.FindOne().SetProjection(bson.M{"_id": 1})

	// Execute the query
	result := collection.FindOne(context.Background(), filter, options)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		// An error occurred while executing the query
		return false, result.Err()
	}

	// User exists in the database
	return true, nil
}
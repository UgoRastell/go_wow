package db

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func AddNewUserToDatabase(client *mongo.Client, guildID string, userID string, username string) error {
	user := bson.M{
		"guild_id": guildID,
		"user_id":  userID,
		"username": username,
	}

	err := InsertDocument(client, "gowow", "users", user)
	if err != nil {
		fmt.Printf("Erreur lors de l'insertion des données dans la collection : %v\n", err)
		return err
	}
	return nil
}
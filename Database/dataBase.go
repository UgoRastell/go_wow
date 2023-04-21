package db

import (
    "context"
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnexionDatabase() (*mongo.Client, error) {

    err := godotenv.Load("./tokens/.env")
    if err != nil {
        fmt.Println("Error loading .env file: ", err)
        return nil, err
    }

    // Récupérer le token de bot depuis les variables d'environnement
    access, ok := os.LookupEnv("MDP_DATABASE")
    if !ok {
        fmt.Println("MDP_DATABASE environment variable not found.")
        return nil, err
    }

    // Use the SetServerAPIOptions() method to set the Stable API version to 1
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI("mongodb+srv://GO_U:" + access + "@gowow.qjwmv9s.mongodb.net/test").SetServerAPIOptions(serverAPI)

    // Create a new client and connect to the server
    client, err := mongo.Connect(context.Background(), opts)
    if err != nil {
        return nil, err
    }

    // Send a ping to confirm a successful connection
    if err = client.Ping(context.Background(), nil); err != nil {
        return nil, err
    }

    fmt.Println("You successfully connected to MongoDB!")
    return client, nil
}

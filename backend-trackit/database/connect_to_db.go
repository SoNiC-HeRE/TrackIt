
package database

import (
    "context"
    "log"
    "os"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    DB     *mongo.Database
    Client *mongo.Client
)

func InitDatabase() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatal("MONGODB_URI not set in environment")
    }

    
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Failed to ping MongoDB:", err)
    }

    
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "task_management"
    }

    DB = client.Database(dbName)
    Client = client

    log.Println("Connected to MongoDB successfully")
}


func GetCollection(name string) *mongo.Collection {
    return DB.Collection(name)
}
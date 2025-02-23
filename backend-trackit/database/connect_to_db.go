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
    // DB holds the reference to the MongoDB database instance
    DB *mongo.Database
    // Client holds the reference to the MongoDB client instance
    Client *mongo.Client
)

// getEnv retrieves an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

// InitDatabase initializes the MongoDB connection
func InitDatabase() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    mongoURI := getEnv("MONGODB_URI", "")
    if mongoURI == "" {
        log.Fatalf("Environment variable MONGODB_URI is not set")
    }

    // Create a new MongoDB client
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Verify the connection by pinging the database
    if err = client.Ping(ctx, nil); err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    dbName := getEnv("DB_NAME", "data_trackit")

    DB = client.Database(dbName)
    Client = client

    log.Println("✅ Successfully connected to MongoDB:", dbName)
}

// GetCollection returns a reference to the specified MongoDB collection
func GetCollection(name string) *mongo.Collection {
    return DB.Collection(name)
}

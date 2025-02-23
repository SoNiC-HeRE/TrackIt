package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "backend-trackit/handlers"
    "backend-trackit/middleware"
    "backend-trackit/database"
)

func main() {
    // Load .env in development
    if os.Getenv("GO_ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("Warning: .env file not found")
        }
    }

    // Initialize database
    database.InitDatabase()

    // Initialize Gin
    r := gin.Default()
    
    // Add middleware
    r.Use(middleware.CORSMiddleware())

    // Routes
    api := r.Group("/api")
    {
        api.POST("/register", handlers.Register)
        api.POST("/login", handlers.Login)
        api.POST("/logout", handlers.Logout)
        api.GET("/me", handlers.GetMe)
        api.GET("/tasks", handlers.GetTasks)
        api.POST("/tasks", handlers.CreateTask)
        api.PUT("/tasks/:id", handlers.UpdateTask)
        api.DELETE("/tasks/:id", handlers.DeleteTask)
        api.POST("/ai/suggestions", handlers.GetAISuggestions)
    }

    // Health check endpoint
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "healthy",
            "message": "Server is running",
        })
    })

    // Get port from environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "10000" // Render's default port
    }

    // Create server with explicit host and port binding
    server := &http.Server{
        Addr:    fmt.Sprintf("0.0.0.0:%s", port),
        Handler: r,
    }

    // Log server start
    log.Printf("Server starting on http://0.0.0.0:%s", port)

    // Start server
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal("Error starting server:", err)
    }
}
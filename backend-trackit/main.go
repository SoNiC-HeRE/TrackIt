package main

import (
    "log"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin router
    r := gin.Default()

    // Health check endpoint
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  "healthy",
            "message": "Server is running",
        })
    })

    // Get port from environment variable (default: 10000)
    port := os.Getenv("PORT")
    if port == "" {
        port = "10000"
    }

    // Start server
    log.Printf("Server starting on port %s...", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Error starting server:", err)
    }
}

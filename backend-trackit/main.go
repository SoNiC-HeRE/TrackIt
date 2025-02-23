package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"backend-trackit/config"
	"backend-trackit/database"
	"backend-trackit/middleware"
	"backend-trackit/routes"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database connection
	database.InitDatabase()

	// Initialize Gin Router
	r := gin.Default()

	// Apply middleware
	r.Use(middleware.CORSMiddleware())

	// Setup API routes
	routes.RegisterRoutes(r)

	// Health check endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "Server is running",
		})
	})

	// Start server with graceful shutdown handling
	startServer(r)
}

// startServer initializes and starts the HTTP server with graceful shutdown
func startServer(router *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000" // Default port
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	// Channel to listen for termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Server started on http://localhost:%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for termination signal
	<-quit
	log.Println("Shutting down server...")

	// Gracefully shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}

package handlers

import (
    "fmt"    
    "log"
    "os"
    "time"
    "github.com/gin-gonic/gin"
    "task-management/internal/services"
)

type AIRequest struct {
    Prompt string `json:"prompt" binding:"required"`
}

type AIResponse struct {
    Suggestions string    `json:"suggestions"`
    Timestamp   time.Time `json:"timestamp"`
}

func GetAISuggestions(c *gin.Context) {
    var request AIRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        log.Printf("Invalid request: %v", err)
        c.JSON(400, gin.H{"error": "Invalid request. Prompt is required"})
        return
    }

    
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        log.Printf("OpenAI API key not set")
        c.JSON(500, gin.H{"error": "OpenAI API key not configured"})
        return
    }

    log.Printf("Initializing AI service with key length: %d", len(apiKey))

   
    aiService, err := services.NewAIService(apiKey)
    if err != nil {
        log.Printf("Failed to initialize AI service: %v", err)
        c.JSON(500, gin.H{"error": "Failed to initialize AI service"})
        return
    }

    log.Printf("Generating response for prompt: %s", request.Prompt)

    
    suggestions, err := aiService.GenerateResponse(request.Prompt)
    if err != nil {
        log.Printf("Failed to generate suggestions: %v", err)
        c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to generate suggestions: %v", err)})
        return
    }

    response := AIResponse{
        Suggestions: suggestions,
        Timestamp:   time.Now(),
    }

    c.JSON(200, response)
}
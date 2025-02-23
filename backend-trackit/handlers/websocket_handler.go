package handlers

import (
    "log"
	"os"
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "github.com/golang-jwt/jwt/v4"
    "backend-trackit/services"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true 
    },
}

func HandleWebSocket(c *gin.Context) {
    
    token := c.Query("token")
    if token == "" {
        log.Printf("No token provided")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
        return
    }

    
    claims := jwt.MapClaims{}
    _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        log.Printf("Invalid token: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    // Get user ID from claims
    userID, ok := claims["user_id"].(string)
    if !ok {
        log.Printf("User ID not found in token")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    log.Printf("WebSocket connection attempt from user: %s", userID)

    // Upgrade connection
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Failed to upgrade connection: %v", err)
        return
    }

    // Create client
    client := &services.Client{
        Hub:  services.WebsocketHub,
        ID:   userID,
        Conn: conn,
        Send: make(chan []byte, 256),
    }

    // Register client
    client.Hub.Register <- client

    // Start client routines
    go client.WritePump()
    go client.ReadPump()

    log.Printf("WebSocket connection established for user: %s", userID)
}
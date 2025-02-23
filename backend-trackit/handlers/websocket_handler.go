package handlers

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "github.com/gorilla/websocket"
    "backend-trackit/services"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin:     func(r *http.Request) bool { return true },
}

func HandleWebSocket(c *gin.Context) {
    token := c.Query("token")
    if token == "" {
        log.Println("No token provided")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
        return
    }

    userID, err := validateToken(token)
    if err != nil {
        log.Printf("Token validation failed: %v", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    log.Printf("WebSocket connection attempt from user: %s", userID)

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Failed to upgrade connection: %v", err)
        return
    }

    client := &services.Client{
        Hub:  services.WebsocketHub,
        ID:   userID,
        Conn: conn,
        Send: make(chan []byte, 256),
    }

    client.Hub.Register <- client

    go client.WritePump()
    go client.ReadPump()

    log.Printf("WebSocket connection established for user: %s", userID)
}

func validateToken(token string) (string, error) {
    claims := jwt.MapClaims{}
    _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil {
        return "", err
    }

    userID, ok := claims["user_id"].(string)
    if !ok {
        return "", jwt.ErrSignatureInvalid
    }
    return userID, nil
}
package middleware

import (
    "fmt"
    "os"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
    UserId string `json:"user_id"`
    jwt.StandardClaims
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        headers := c.Writer.Header()
        headers.Set("Access-Control-Allow-Origin", "*")
        headers.Set("Access-Control-Allow-Credentials", "true")
        headers.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        headers.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

func GenerateToken(userId string) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return "", fmt.Errorf("JWT_SECRET not set")
    }

    claims := Claims{
        UserId: userId,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
            IssuedAt:  time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (*Claims, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return nil, fmt.Errorf("JWT_SECRET not set")
    }

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        fmt.Println("Token validation error:", err) // Debugging line
        return nil, fmt.Errorf("invalid token: %v", err)
    }
    
    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == "" {
            c.JSON(401, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        claims, err := ValidateToken(tokenString)
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("userId", claims.UserId)
        c.Next()
    }
}
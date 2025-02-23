
package handlers

import (
    "context"
    "time"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "golang.org/x/crypto/bcrypt"
    "task-management/internal/database"
    "task-management/internal/middleware"
)

type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name     string            `bson:"name" json:"name"`
    Email    string            `bson:"email" json:"email"`
    Password string            `bson:"password" json:"-"`
}

func Register(c *gin.Context) {
    var input struct {
        Name     string `json:"name" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

   
    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var existingUser User
    err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
    if err == nil {
        c.JSON(400, gin.H{"error": "Email already registered"})
        return
    }

    
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to hash password"})
        return
    }

    
    user := User{
        ID:       primitive.NewObjectID(),
        Name:     input.Name,
        Email:    input.Email,
        Password: string(hashedPassword),
    }

    _, err = collection.InsertOne(ctx, user)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }

   
    token, err := middleware.GenerateToken(user.ID.Hex())
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(200, gin.H{
        "token": token,
        "user": gin.H{
            "id":    user.ID.Hex(),
            "name":  user.Name,
            "email": user.Email,
        },
    })
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var user User
    err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
    if err != nil {
        c.JSON(401, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(401, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := middleware.GenerateToken(user.ID.Hex())
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(200, gin.H{
        "token": token,
        "user": gin.H{
            "id":    user.ID.Hex(),
            "name":  user.Name,
            "email": user.Email,
        },
    })
}

func GetMe(c *gin.Context) {
    userId, exists := c.Get("userId")
    if !exists {
        c.JSON(401, gin.H{"error": "Unauthorized"})
        return
    }

    objectId, err := primitive.ObjectIDFromHex(userId.(string))
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid user ID"})
        return
    }

    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var user User
    err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
    if err != nil {
        c.JSON(404, gin.H{"error": "User not found"})
        return
    }

    c.JSON(200, gin.H{
        "user": gin.H{
            "id":    user.ID.Hex(),
            "name":  user.Name,
            "email": user.Email,
        },
    })
}

func Logout(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Logged out successfully",
    })
}
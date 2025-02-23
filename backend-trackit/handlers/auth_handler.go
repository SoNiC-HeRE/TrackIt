package handlers

import (
    "context"
    "log"
    "time"
    "strings"
    "fmt"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"

    "backend-trackit/database"
    "backend-trackit/middleware"
)

// Collection names
const userCollection = "users"

// User model struct
type User struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name     string             `bson:"name" json:"name"`
    Email    string             `bson:"email" json:"email"`
    Password string             `bson:"password" json:"-"`
}

// Register handles user registration
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

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    collection := database.GetCollection(userCollection)

    // Check if email is already registered
    if userExists(ctx, collection, input.Email) {
        c.JSON(400, gin.H{"error": "Email already registered"})
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error hashing password:", err)
        c.JSON(500, gin.H{"error": "Failed to hash password"})
        return
    }

    // Create user object
    user := User{
        ID:       primitive.NewObjectID(),
        Name:     input.Name,
        Email:    input.Email,
        Password: string(hashedPassword),
    }

    if _, err := collection.InsertOne(ctx, user); err != nil {
        log.Println("Error inserting user:", err)
        c.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }

    // Generate JWT token
    token, err := middleware.GenerateToken(user.ID.Hex())
    fmt.Println("Extracted Token:", token)
    if err != nil {
        log.Println("Error generating token:", err)
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(200, gin.H{
        "token": token,
        "user":  mapUserResponse(user),
    })
}

// Login handles user authentication
func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    collection := database.GetCollection(userCollection)

    var user User
    if err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user); err != nil {
        c.JSON(401, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(401, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := middleware.GenerateToken(user.ID.Hex())
    if err != nil {
        log.Println("Error generating token:", err)
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(200, gin.H{
        "token": token,
        "user":  mapUserResponse(user),
    })
}

// GetMe retrieves the authenticated user's data
func GetMe(c *gin.Context) {
    userID, exists := c.Get("userId")
    if !exists {
        c.JSON(401, gin.H{"error": "Unauthorized"})
        return
    }

    objectID, err := primitive.ObjectIDFromHex(userID.(string))
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid user ID"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    collection := database.GetCollection(userCollection)

    var user User
    if err := collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user); err != nil {
        c.JSON(404, gin.H{"error": "User not found"})
        return
    }

    c.JSON(200, gin.H{"user": mapUserResponse(user)})
}

// Logout handles user logout
func Logout(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(401, gin.H{"error": "Authorization header required"})
        return
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    if tokenString == "" {
        c.JSON(401, gin.H{"error": "Invalid token format"})
        return
    }

    _, err := middleware.ValidateToken(tokenString)
    if err != nil {
        c.JSON(401, gin.H{"error": "Invalid token"})
        return
    }

    // Usually, JWTs are stateless, so we don't "destroy" them on the server.
    // Instead, we ask the client to discard the token.
    c.JSON(200, gin.H{"message": "Logged out successfully"})
}


// userExists checks if a user with the given email already exists
func userExists(ctx context.Context, collection *mongo.Collection, email string) bool {
    var user User
    err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
    return err == nil
}

// mapUserResponse formats user data for JSON response
func mapUserResponse(user User) gin.H {
    return gin.H{
        "id":    user.ID.Hex(),
        "name":  user.Name,
        "email": user.Email,
    }
}

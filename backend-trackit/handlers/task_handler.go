package handlers

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "backend-trackit/models"
    "backend-trackit/database"
    "backend-trackit/services"
)

// CreateTask handles creating a new task
func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        respondWithError(c, 400, "Invalid request payload", err)
        return
    }

    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
    task.ID = primitive.NewObjectID()
    task.CreatedBy = userID
    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()

    if err := insertTask(task); err != nil {
        respondWithError(c, 500, "Failed to create task", err)
        return
    }

    generateAISuggestions(task)

    c.JSON(201, gin.H{"message": "Task created successfully", "task": task})
}

// GetTasks retrieves all tasks for a user
func GetTasks(c *gin.Context) {
    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))

    tasks, err := fetchUserTasks(userID)
    if err != nil {
        respondWithError(c, 500, "Failed to fetch tasks", err)
        return
    }

    c.JSON(200, gin.H{"tasks": tasks})
}

// UpdateTask modifies an existing task
func UpdateTask(c *gin.Context) {
    taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))
    var updateData models.Task

    if err := c.ShouldBindJSON(&updateData); err != nil {
        respondWithError(c, 400, "Invalid request payload", err)
        return
    }

    updateData.UpdatedAt = time.Now()

    if modifiedCount, err := updateTask(taskID, updateData); err != nil {
        respondWithError(c, 500, "Failed to update task", err)
        return
    } else if modifiedCount == 0 {
        c.JSON(404, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(200, gin.H{"message": "Task updated successfully"})
}

// DeleteTask removes a task if the user is authorized
func DeleteTask(c *gin.Context) {
    taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))
    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))

    if deletedCount, err := deleteTask(taskID, userID); err != nil {
        respondWithError(c, 500, "Failed to delete task", err)
        return
    } else if deletedCount == 0 {
        c.JSON(404, gin.H{"error": "Task not found or unauthorized"})
        return
    }

    c.JSON(200, gin.H{"message": "Task deleted successfully"})
}

// ------------------ Helper Functions ------------------

// Insert a task into the database
func insertTask(task models.Task) error {
    collection := database.Client.Database("data_trackit").Collection("tasks")
    _, err := collection.InsertOne(context.Background(), task)
    return err
}

// Fetch tasks for a user
func fetchUserTasks(userID primitive.ObjectID) ([]models.Task, error) {
    collection := database.Client.Database("data_trackit").Collection("tasks")
    cursor, err := collection.Find(context.Background(), bson.M{
        "$or": []bson.M{
            {"created_by": userID},
            {"assigned_to": userID},
        },
    })
    if err != nil {
        return nil, err
    }

    var tasks []models.Task
    if err := cursor.All(context.Background(), &tasks); err != nil {
        return nil, err
    }
    return tasks, nil
}

// Update a task in the database
func updateTask(taskID primitive.ObjectID, updateData models.Task) (int64, error) {
    collection := database.Client.Database("data_trackit").Collection("tasks")
    result, err := collection.UpdateOne(
        context.Background(),
        bson.M{"_id": taskID},
        bson.M{"$set": updateData},
    )
    return result.ModifiedCount, err
}

// Delete a task from the database
func deleteTask(taskID, userID primitive.ObjectID) (int64, error) {
    collection := database.Client.Database("data_trackit").Collection("tasks")
    result, err := collection.DeleteOne(context.Background(), bson.M{
        "_id":        taskID,
        "created_by": userID,
    })
    return result.DeletedCount, err
}

// Generate AI-based suggestions for a task
func generateAISuggestions(task models.Task) {
    aiService, err := services.NewAIService(os.Getenv("OPENAI_API_KEY"))
    if err != nil {
        log.Printf("AI Service initialization failed: %v", err)
        return
    }

    suggestions, err := aiService.GenerateResponse(task.Title + ": " + task.Description)
    if err != nil {
        log.Printf("Error generating AI suggestions: %v", err)
        return
    }

    suggCollection := database.Client.Database("data_trackit").Collection("ai_suggestions")
    _, err = suggCollection.InsertOne(context.Background(), models.AITaskSuggestion{
        TaskID:      task.ID,
        Suggestion:  suggestions,
        GeneratedAt: time.Now(),
    })
    if err != nil {
        log.Printf("Failed to save AI suggestion: %v", err)
    }
}

// Respond with an error in a structured format
func respondWithError(c *gin.Context, status int, message string, err error) {
    log.Printf("%s: %v", message, err)
    c.JSON(status, gin.H{"error": message})
}

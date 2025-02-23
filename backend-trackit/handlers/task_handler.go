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

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
    task.ID = primitive.NewObjectID()
    task.CreatedBy = userID
    task.CreatedAt = time.Now()
    task.UpdatedAt = time.Now()

    collection := database.Client.Database("taskmanagement").Collection("tasks")
    _, err := collection.InsertOne(context.Background(), task)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create task"})
        return
    }

  
    aiService, err := services.NewAIService(os.Getenv("OPENAI_API_KEY"))
    if err == nil {
        suggestions, err := aiService.GenerateResponse(task.Title + ": " + task.Description)
        if err != nil {
            log.Printf("Error generating AI suggestions: %v", err)
        } else {
            
            suggCollection := database.Client.Database("taskmanagement").Collection("ai_suggestions")
            suggCollection.InsertOne(context.Background(), models.AITaskSuggestion{
                TaskID:      task.ID,
                Suggestion:  suggestions,
                GeneratedAt: time.Now(),
            })
        }
    }

    c.JSON(201, task)
}

func GetTasks(c *gin.Context) {
    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))
    
    collection := database.Client.Database("taskmanagement").Collection("tasks")
    cursor, err := collection.Find(context.Background(), bson.M{
        "$or": []bson.M{
            {"created_by": userID},
            {"assigned_to": userID},
        },
    })
    
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to fetch tasks"})
        return
    }

    var tasks []models.Task
    if err = cursor.All(context.Background(), &tasks); err != nil {
        c.JSON(500, gin.H{"error": "Failed to decode tasks"})
        return
    }

    c.JSON(200, tasks)
}

func UpdateTask(c *gin.Context) {
    taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))
    var updateData models.Task
    
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    updateData.UpdatedAt = time.Now()

    collection := database.Client.Database("taskmanagement").Collection("tasks")
    result, err := collection.UpdateOne(
        context.Background(),
        bson.M{"_id": taskID},
        bson.M{"$set": updateData},
    )

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to update task"})
        return
    }

    if result.ModifiedCount == 0 {
        c.JSON(404, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(200, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
    taskID, _ := primitive.ObjectIDFromHex(c.Param("id"))
    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))

    collection := database.Client.Database("taskmanagement").Collection("tasks")
    result, err := collection.DeleteOne(context.Background(), bson.M{
        "_id": taskID,
        "created_by": userID,
    })

    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to delete task"})
        return
    }

    if result.DeletedCount == 0 {
        c.JSON(404, gin.H{"error": "Task not found or unauthorized"})
        return
    }

    c.JSON(200, gin.H{"message": "Task deleted successfully"})
}
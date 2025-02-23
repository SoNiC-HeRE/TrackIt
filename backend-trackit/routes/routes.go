package routes

import (
	"github.com/gin-gonic/gin"
	"backend-trackit/handlers"
)

// RegisterRoutes sets up all API endpoints
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.POST("/logout", handlers.Logout)
		api.GET("/me", handlers.GetMe)
		api.GET("/tasks", handlers.GetTasks)
		api.POST("/tasks", handlers.CreateTask)
		api.PUT("/tasks/:id", handlers.UpdateTask)
		api.DELETE("/tasks/:id", handlers.DeleteTask)
		api.POST("/ai/suggestions", handlers.GetAISuggestions)
	}
}

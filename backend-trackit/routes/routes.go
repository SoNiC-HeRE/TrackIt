package routes

import (
	"github.com/gin-gonic/gin"
	"backend-trackit/handlers"
	"backend-trackit/middleware"
)

// RegisterRoutes sets up all API endpoints
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware()) // Apply authentication middleware
		{
			protected.POST("/logout", handlers.Logout) // Move logout inside protected routes
			protected.GET("/me", handlers.GetMe)
			protected.GET("/tasks", handlers.GetTasks)
			protected.POST("/tasks", handlers.CreateTask)
			protected.PUT("/tasks/:id", handlers.UpdateTask)
			protected.DELETE("/tasks/:id", handlers.DeleteTask)
			protected.POST("/ai/suggestions", handlers.GetAISuggestions)
		}
	}
}

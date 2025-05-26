package api

import (
	"github.com/EchoLearnAI/EchoLearn/api/handlers"
	_ "github.com/EchoLearnAI/EchoLearn/docs" // docs is required for Swag
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configures the Gin router with all API routes and applies CORS.
func SetupRouter(corsMW gin.HandlerFunc) *gin.Engine {
	r := gin.Default()

	// Apply CORS middleware first
	r.Use(corsMW)

	// Swagger endpoint - must be defined AFTER CORS if it needs to be accessible cross-origin
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// User routes
		userRoutes := apiv1.Group("/users")
		{
			userRoutes.POST("", handlers.CreateUser)
			userRoutes.GET("/:id", handlers.GetUserByID)
		}

		// Question routes
		questionRoutes := apiv1.Group("/questions")
		{
			questionRoutes.POST("", handlers.AddQuestion) // Admin purposes
			questionRoutes.GET("/random", handlers.GetRandomQuestion)
			questionRoutes.GET("/category/:name", handlers.GetQuestionsByCategory)
			questionRoutes.GET("/:id", handlers.GetQuestionByID)
		}

		// Session routes
		sessionRoutes := apiv1.Group("/session")
		{
			sessionRoutes.POST("/start", handlers.StartSession)
			sessionRoutes.POST("/submit", handlers.SubmitAnswer)
			sessionRoutes.GET("/:id/summary", handlers.GetSessionSummary)
		}

		// Auth routes
		authRoutes := apiv1.Group("/auth")
		{
			authRoutes.POST("/login", handlers.Login) // New login route
		}
	}

	return r
}

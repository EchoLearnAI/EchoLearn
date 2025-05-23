package api

import (
	"github.com/gin-gonic/gin"
	"github.com/EchoLearnAI/EchoLearn/api/handlers"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/EchoLearnAI/EchoLearn/docs" // docs is required for Swag
)

SetupRouter configures the Gin router with all API routes.
@title EchoLearn API
@version 1.0
@description This is the API for the EchoLearn English learning app.
@termsOfService http://swagger.io/terms/
@contact.name API Support
@contact.url http://www.swagger.io/support
@contact.email support@swagger.io
@license.name Apache 2.0
@license.url http://www.apache.org/licenses/LICENSE-2.0.html
@host localhost:8080
@BasePath /api/v1

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// TODO: Add CORS middleware if React Native app is served from a different origin during dev
	// r.Use(cors.Default()) // Example: gin-contrib/cors

	// Swagger endpoint
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
	}

	return r
} 
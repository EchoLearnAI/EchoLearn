package routes

import (
	"echolearn/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App) {
	// User routes
	users := app.Group("/users")
	users.Post("/", controllers.CreateUser)
	users.Get("/", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUser)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)

	// Question routes
	questions := app.Group("/questions")
	questions.Get("/random", controllers.GetRandomQuestion)
	questions.Get("/category/:name", controllers.GetQuestionsByCategory)
	questions.Get("/:id", controllers.GetQuestion)
	questions.Post("/", controllers.CreateQuestion)
	questions.Put("/:id", controllers.UpdateQuestion)
	questions.Delete("/:id", controllers.DeleteQuestion)

	// Session routes
	sessions := app.Group("/sessions")
	sessions.Post("/start", controllers.StartSession)
	sessions.Post("/submit", controllers.SubmitAnswer)
	sessions.Post("/finish", controllers.FinishSession)
	sessions.Get("/:id/summary", controllers.GetSessionSummary)
} 
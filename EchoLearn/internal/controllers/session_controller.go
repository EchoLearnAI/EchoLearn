package controllers

import (
	"echolearn/internal/db"
	"echolearn/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// StartSession initializes a new game session
func StartSession(c *fiber.Ctx) error {
	type SessionRequest struct {
		UserID string          `json:"user_id"`
		Mode   models.GameMode `json:"mode"`
	}

	var req SessionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Validate user ID
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Check if user exists
	var user models.User
	if err := db.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Validate game mode
	if req.Mode != models.SurvivalMode && 
	   req.Mode != models.TopicMode && 
	   req.Mode != models.InfiniteMode {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid game mode",
		})
	}

	// Create new session
	session := models.Session{
		UserID: userID,
		Mode:   req.Mode,
	}

	if err := db.DB.Create(&session).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create session",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(session)
}

// SubmitAnswer records a user's answer in a session
func SubmitAnswer(c *fiber.Ctx) error {
	type AnswerRequest struct {
		SessionID  string `json:"session_id"`
		QuestionID string `json:"question_id"`
		OptionID   string `json:"option_id"`
	}

	var req AnswerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Parse UUIDs
	sessionID, err := uuid.Parse(req.SessionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	questionID, err := uuid.Parse(req.QuestionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid question ID",
		})
	}

	optionID, err := uuid.Parse(req.OptionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid option ID",
		})
	}

	// Get the session
	var session models.Session
	if err := db.DB.First(&session, "id = ?", sessionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Session not found",
		})
	}

	// Check if session is already completed
	if session.IsComplete {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Session is already completed",
		})
	}

	// Find the option to check if it's correct
	var option models.Option
	if err := db.DB.First(&option, "id = ? AND question_id = ?", optionID, questionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Option not found or doesn't belong to the question",
		})
	}

	// Create answer record
	answer := models.Answer{
		SessionID:  sessionID,
		QuestionID: questionID,
		OptionID:   optionID,
		IsCorrect:  option.IsCorrect,
	}

	// Start a transaction
	tx := db.DB.Begin()

	// Save the answer
	if err := tx.Create(&answer).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save answer",
		})
	}

	// Update session score and mistakes
	if option.IsCorrect {
		session.Score++
	} else {
		session.Mistakes++
	}

	// Check if the session should be completed based on game mode
	if session.Mode == models.SurvivalMode && session.Mistakes >= 3 {
		session.IsComplete = true
		now := time.Now()
		session.EndedAt = &now
	}

	if err := tx.Save(&session).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update session",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	type AnswerResponse struct {
		IsCorrect  bool           `json:"is_correct"`
		IsComplete bool           `json:"is_complete"`
		Score      int            `json:"score"`
		Mistakes   int            `json:"mistakes"`
		Option     models.Option  `json:"option"`
	}

	return c.JSON(AnswerResponse{
		IsCorrect:  option.IsCorrect,
		IsComplete: session.IsComplete,
		Score:      session.Score,
		Mistakes:   session.Mistakes,
		Option:     option,
	})
}

// FinishSession completes a session
func FinishSession(c *fiber.Ctx) error {
	type FinishRequest struct {
		SessionID string `json:"session_id"`
	}

	var req FinishRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	sessionID, err := uuid.Parse(req.SessionID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	var session models.Session
	if err := db.DB.First(&session, "id = ?", sessionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Session not found",
		})
	}

	if session.IsComplete {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Session is already completed",
		})
	}

	// Mark the session as complete
	now := time.Now()
	session.IsComplete = true
	session.EndedAt = &now

	if err := db.DB.Save(&session).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to finish session",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Session finished successfully",
		"session": session,
	})
}

// GetSessionSummary generates a summary for a completed session
func GetSessionSummary(c *fiber.Ctx) error {
	id := c.Params("id")

	sessionID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	var session models.Session
	if err := db.DB.First(&session, "id = ?", sessionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Session not found",
		})
	}

	// If session is not complete, return current state
	if !session.IsComplete {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Session is not completed yet",
			"session": session,
		})
	}

	// Get answers with their questions
	var answers []models.Answer
	if err := db.DB.Preload("Question").
		Where("session_id = ?", sessionID).Find(&answers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch session answers",
		})
	}

	// Calculate statistics
	summary := models.SessionSummary{
		TotalQuestions: len(answers),
		CorrectAnswers: session.Score,
		Mistakes:       session.Mistakes,
		Score:          session.Score,
		GameMode:       session.Mode,
	}

	if session.EndedAt != nil && !session.StartedAt.IsZero() {
		summary.Duration = session.EndedAt.Sub(session.StartedAt)
	}

	// Calculate scores by category
	categoryMap := make(map[string]*models.CategoryScore)
	
	for _, answer := range answers {
		category := answer.Question.Category
		
		if _, exists := categoryMap[category]; !exists {
			categoryMap[category] = &models.CategoryScore{
				Category: category,
			}
		}
		
		categoryMap[category].TotalQuestions++
		if answer.IsCorrect {
			categoryMap[category].CorrectAnswers++
		}
	}
	
	// Convert map to slice
	for _, score := range categoryMap {
		summary.ByCategory = append(summary.ByCategory, *score)
	}

	return c.JSON(summary)
} 
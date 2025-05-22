package controllers

import (
	"echolearn/internal/db"
	"echolearn/internal/models"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandomQuestion fetches a random question with its options and grammar rule
func GetRandomQuestion(c *fiber.Ctx) error {
	var count int64
	if err := db.DB.Model(&models.Question{}).Count(&count).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to count questions",
		})
	}
	
	if count == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No questions available",
		})
	}
	
	// Get random question
	offset := rand.Intn(int(count))
	
	var question models.Question
	if err := db.DB.Preload("Options").Preload("Grammar").
		Offset(offset).Limit(1).Find(&question).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch random question",
		})
	}
	
	return c.JSON(question)
}

// GetQuestionsByCategory fetches questions by category
func GetQuestionsByCategory(c *fiber.Ctx) error {
	category := c.Params("name")
	limit := 10 // Default limit
	
	var questions []models.Question
	if err := db.DB.Preload("Options").Preload("Grammar").
		Where("category = ?", category).Limit(limit).Find(&questions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch questions by category",
		})
	}
	
	if len(questions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No questions found for this category",
		})
	}
	
	return c.JSON(questions)
}

// GetQuestion gets a question by ID
func GetQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid question ID",
		})
	}
	
	var question models.Question
	if err := db.DB.Preload("Options").Preload("Grammar").
		First(&question, "id = ?", uid).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Question not found",
		})
	}
	
	return c.JSON(question)
}

// CreateQuestion creates a new question with its options and grammar rule
func CreateQuestion(c *fiber.Ctx) error {
	question := new(models.Question)
	
	if err := c.BodyParser(question); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	
	// Check that there are exactly 4 options
	if len(question.Options) != 4 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Question must have exactly 4 options",
		})
	}
	
	// Check that one and only one option is correct
	correctCount := 0
	for _, option := range question.Options {
		if option.IsCorrect {
			correctCount++
		}
	}
	
	if correctCount != 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Question must have exactly 1 correct option",
		})
	}
	
	// Begin transaction for creating question with options
	tx := db.DB.Begin()
	
	// Create grammar rule if it doesn't have an ID
	if question.Grammar.ID == uuid.Nil {
		if err := tx.Create(&question.Grammar).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create grammar rule",
			})
		}
		question.GrammarID = question.Grammar.ID
	}
	
	// Create question
	if err := tx.Create(&question).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create question",
		})
	}
	
	// Create options
	for i := range question.Options {
		question.Options[i].QuestionID = question.ID
		if err := tx.Create(&question.Options[i]).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create options",
			})
		}
	}
	
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}
	
	// Reload the complete question with associations
	db.DB.Preload("Options").Preload("Grammar").First(&question, "id = ?", question.ID)
	
	return c.Status(fiber.StatusCreated).JSON(question)
}

// UpdateQuestion updates a question by ID
func UpdateQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid question ID",
		})
	}
	
	var existingQuestion models.Question
	if err := db.DB.Preload("Options").First(&existingQuestion, "id = ?", uid).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Question not found",
		})
	}
	
	var updatedQuestion models.Question
	if err := c.BodyParser(&updatedQuestion); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	
	// Set ID to ensure we update the correct question
	updatedQuestion.ID = uid
	
	// Update question fields
	db.DB.Model(&existingQuestion).Updates(updatedQuestion)
	
	return c.JSON(existingQuestion)
}

// DeleteQuestion deletes a question by ID
func DeleteQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid question ID",
		})
	}
	
	// Delete options first (use transaction for atomicity)
	tx := db.DB.Begin()
	
	if err := tx.Where("question_id = ?", uid).Delete(&models.Option{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete options",
		})
	}
	
	// Delete the question
	result := tx.Delete(&models.Question{}, "id = ?", uid)
	if result.Error != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	
	if result.RowsAffected == 0 {
		tx.Rollback()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Question not found",
		})
	}
	
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Question deleted successfully",
	})
} 
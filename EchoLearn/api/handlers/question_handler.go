package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/EchoLearnAI/EchoLearn/db"
	"github.com/EchoLearnAI/EchoLearn/models"
	"github.com/EchoLearnAI/EchoLearn/utils"
)

// AddQuestion godoc
// @Summary Add a new question (admin)
// @Description Adds a new question to the database. Requires admin privileges (not implemented in MVP).
// @Tags questions
// @Accept  json
// @Produce  json
// @Param   question  body   models.Question  true  "Question object"
// @Success 201 {object} models.Question
// @Failure 400 {object} utils.APIError "Invalid input"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /questions [post]
func AddQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// Basic validation
	if question.Text == "" || len(question.Options) != 4 || question.Category == "" || question.GrammarRule.Title == "" {
		utils.BadRequest(c, "Missing required question fields (text, 4 options, category, grammar rule title)")
		return
	}

	correctOptionExists := false
	for _, opt := range question.Options {
		if opt.Text == "" {
			utils.BadRequest(c, "Option text cannot be empty")
			return
		}
		if opt.IsCorrect {
			correctOptionExists = true
		}
	}
	if !correctOptionExists {
		utils.BadRequest(c, "At least one option must be marked as correct")
		return
	}

	dbInstance := db.GetDB()

	// Check if grammar rule exists, or create it
	var grammarRule models.GrammarRule
	err := dbInstance.Where("title = ?", question.GrammarRule.Title).First(&grammarRule).Error
	if err != nil && !db.IsErrNotFound(err) {
		utils.InternalServerError(c, "Error checking grammar rule: "+err.Error())
		return
	}

	// If grammar rule doesn't exist by title, create it
	if db.IsErrNotFound(err) {
		newRule := models.GrammarRule{
			Title:       question.GrammarRule.Title,
			Description: question.GrammarRule.Description,
			Examples:    question.GrammarRule.Examples,
		}
		if err := dbInstance.Create(&newRule).Error; err != nil {
			utils.InternalServerError(c, "Could not create grammar rule: "+err.Error())
			return
		}
		question.GrammarRuleID = newRule.ID
		question.GrammarRule = newRule // Ensure the response includes the full rule
	} else {
		question.GrammarRuleID = grammarRule.ID
		question.GrammarRule = grammarRule // Ensure response includes existing rule details
	}

	// The BeforeCreate hook on Question will generate its ID.
	// For Options, their BeforeCreate hook will generate their IDs.
	// We need to set QuestionID for each option manually *before* creating the Question record
	// if we were to create options separately. However, GORM handles this with associations if structured correctly.
	// The foreignKey in Question.Options `gorm:"foreignKey:QuestionID"` and `constraint:OnDelete:CASCADE`
	// should handle the relationship and cascading deletes.

	if err := dbInstance.Create(&question).Error; err != nil {
		utils.InternalServerError(c, "Could not create question: "+err.Error())
		return
	}

	// Reload the question to ensure all associated data (like options with IDs) is fresh
	var createdQuestion models.Question
	if err := dbInstance.Preload("Options").Preload("GrammarRule").First(&createdQuestion, "id = ?", question.ID).Error; err != nil {
	    utils.InternalServerError(c, "Could not retrieve created question with associations: "+err.Error())
	    return
	}

	c.JSON(http.StatusCreated, createdQuestion)
}

// GetRandomQuestion godoc
// @Summary Get a random question
// @Description Retrieves a single random question from the database
// @Tags questions
// @Produce  json
// @Success 200 {object} models.Question
// @Failure 404 {object} utils.APIError "No questions available"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /questions/random [get]
func GetRandomQuestion(c *gin.Context) {
	dbInstance := db.GetDB()
	var question models.Question
	var count int64

	dbInstance.Model(&models.Question{}).Count(&count)
	if count == 0 {
		utils.NotFound(c, "No questions available")
		return
	}

	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(int(count))

	if err := dbInstance.Preload("Options").Preload("GrammarRule").Offset(offset).First(&question).Error; err != nil {
		utils.InternalServerError(c, "Could not retrieve random question: "+err.Error())
		return
	}
	c.JSON(http.StatusOK, question)
}

// GetQuestionsByCategory godoc
// @Summary Get questions by category
// @Description Retrieves 10 questions for a given category
// @Tags questions
// @Produce  json
// @Param   name  path   string  true  "Category name"
// @Success 200 {array}  models.Question
// @Failure 404 {object} utils.APIError "No questions found for this category"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /questions/category/{name} [get]
func GetQuestionsByCategory(c *gin.Context) {
	categoryName := c.Param("name")
	dbInstance := db.GetDB()
	var questions []models.Question

	if err := dbInstance.Preload("Options").Preload("GrammarRule").Where("category = ?", categoryName).Limit(10).Find(&questions).Error; err != nil {
		utils.InternalServerError(c, "Could not retrieve questions by category: "+err.Error())
		return
	}

	if len(questions) == 0 {
		utils.NotFound(c, "No questions found for category: "+categoryName)
		return
	}

	c.JSON(http.StatusOK, questions)
}

// GetQuestionByID godoc
// @Summary Get a question by ID with detailed feedback
// @Description Retrieves a specific question by its ID, including all options and grammar rule details.
// @Tags questions
// @Produce  json
// @Param   id  path   string  true  "Question ID"
// @Success 200 {object} models.Question
// @Failure 404 {object} utils.APIError "Question not found"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /questions/{id} [get]
func GetQuestionByID(c *gin.Context) {
	questionID := c.Param("id")
	dbInstance := db.GetDB()
	var question models.Question

	if err := dbInstance.Preload("Options").Preload("GrammarRule").First(&question, "id = ?", questionID).Error; err != nil {
		if db.IsErrNotFound(err) {
			utils.NotFound(c, "Question")
		} else {
			utils.InternalServerError(c, "Could not retrieve question: "+err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, question)
} 
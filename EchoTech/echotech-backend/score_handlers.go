package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ScoreSubmission struct {
	Topic      string `json:"topic" binding:"required"`
	Difficulty string `json:"difficulty" binding:"required"`
	Points     int    `json:"points" binding:"required,gte=0"`
	Total      int    `json:"total" binding:"required,gt=0"`
}

// SubmitScore godoc
// @Summary Submit a user's quiz score
// @Description Saves a quiz score for the authenticated user
// @Tags scores
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param score body ScoreSubmission true "Score details"
// @Success 201 {object} ScoreDTO "Score saved successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /scores [post]
func SubmitScore(c *gin.Context) {
	var submission ScoreSubmission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	userIDStr, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format in token"})
		return
	}

	// Find the topic by slug to get its ID
	var topic Topic
	if err := DB.First(&topic, "slug = ?", submission.Topic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid topic: " + submission.Topic})
		return
	}

	score := Score{
		UserID:      uint(userID),
		TopicID:     topic.ID, // Use the ID of the found topic
		Difficulty:  submission.Difficulty,
		Points:      submission.Points,
		Total:       submission.Total,
		AttemptedAt: time.Now(),
	}

	if err := DB.Create(&score).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save score: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, score)
}

// GetUserScores godoc
// @Summary Get scores for the authenticated user
// @Description Retrieves all quiz scores for the currently logged-in user
// @Tags scores
// @Security BearerAuth
// @Produce json
// @Success 200 {array} ScoreDTO "List of user scores"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/me/scores [get]
func GetUserScores(c *gin.Context) {
	userIDStr, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr.(string), 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format in token"})
		return
	}

	var scores []Score
	if err := DB.Where("user_id = ?", uint(userID)).Order("attempted_at desc").Find(&scores).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve scores: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, scores)
}

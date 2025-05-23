package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/EchoLearnAI/EchoLearn/db"
	"github.com/EchoLearnAI/EchoLearn/models"
	"github.com/EchoLearnAI/EchoLearn/utils"
)

// StartSessionRequest defines the expected body for starting a new session.
type StartSessionRequest struct {
	UserID string           `json:"user_id" binding:"required"`
	Mode   models.GameMode `json:"mode" binding:"required"`
}

// SubmitAnswerRequest defines the expected body for submitting an answer.
type SubmitAnswerRequest struct {
	SessionID        string `json:"session_id" binding:"required"`
	QuestionID       string `json:"question_id" binding:"required"`
	SelectedOptionID string `json:"selected_option_id" binding:"required"`
}

// StartSession godoc
// @Summary Start a new game session
// @Description Initializes a new session for a user based on the selected game mode.
// @Tags sessions
// @Accept  json
// @Produce  json
// @Param   session_request  body   StartSessionRequest  true  "Session Start Info"
// @Success 201 {object} models.GameSession
// @Failure 400 {object} utils.APIError "Invalid input or game mode"
// @Failure 404 {object} utils.APIError "User not found"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /session/start [post]
func StartSession(c *gin.Context) {
	var req StartSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	dbInstance := db.GetDB()

	// Verify user exists
	var user models.User
	if err := dbInstance.First(&user, "id = ?", req.UserID).Error; err != nil {
		if db.IsErrNotFound(err) {
			utils.NotFound(c, "User")
		} else {
			utils.InternalServerError(c, "Error fetching user: "+err.Error())
		}
		return
	}

	// Create new game session
	newSession := models.GameSession{
		UserID: req.UserID,
		Mode:   req.Mode,
		// AnsweredQuestionsJSON will be initialized as an empty array string
		AnsweredQuestionsJSON: "[]",
	}

	switch req.Mode {
	case models.SurvivalMode:
		newSession.MaxMistakes = 3
	case models.FiveTopicMode:
		newSession.TotalQuestions = 50 // 10 questions per 5 topics (simplified for MVP)
		// Logic for selecting questions for 5-topic mode would go here.
		// For MVP, we might just pull random questions up to TotalQuestions as they are requested.
	case models.InfiniteMode:
		// No specific setup needed beyond base
	default:
		utils.BadRequest(c, "Invalid game mode specified")
		return
	}

	if err := dbInstance.Create(&newSession).Error; err != nil {
		utils.InternalServerError(c, "Could not start session: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, newSession)
}

// SubmitAnswer godoc
// @Summary Submit an answer for the current question in a session
// @Description Records the user's answer, updates score/mistakes, and determines if the session ends.
// @Tags sessions
// @Accept  json
// @Produce  json
// @Param   answer_request  body   SubmitAnswerRequest  true  "Answer Submission Info"
// @Success 200 {object} map[string]interface{} "result: correct/incorrect, session_active: bool, session_details: models.GameSession"
// @Failure 400 {object} utils.APIError "Invalid input"
// @Failure 404 {object} utils.APIError "Session or Question not found"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /session/submit [post]
func SubmitAnswer(c *gin.Context) {
	var req SubmitAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	dbInstance := db.GetDB()
	var session models.GameSession
	var question models.Question

	// Fetch active session
	if err := dbInstance.First(&session, "id = ? AND is_active = ?", req.SessionID, true).Error; err != nil {
		if db.IsErrNotFound(err) {
			utils.NotFound(c, "Active session")
		} else {
			utils.InternalServerError(c, "Error fetching session: "+err.Error())
		}
		return
	}

	// Fetch question and its options
	if err := dbInstance.Preload("Options").First(&question, "id = ?", req.QuestionID).Error; err != nil {
		if db.IsErrNotFound(err) {
			utils.NotFound(c, "Question")
		} else {
			utils.InternalServerError(c, "Error fetching question: "+err.Error())
		}
		return
	}

	var selectedOption models.Option
	var correctOption models.Option
	isCorrect := false

	for _, opt := range question.Options {
		if opt.ID == req.SelectedOptionID {
			selectedOption = opt
		}
		if opt.IsCorrect {
			correctOption = opt
		}
	}

	if selectedOption.ID == "" {
		utils.BadRequest(c, "Selected option not found for this question")
		return
	}

	if selectedOption.ID == correctOption.ID {
		isCorrect = true
		session.Score++
	} else {
		session.MistakesMade++
	}

	// Record the answer
	var answeredQuestions []models.AnsweredQuestionDetail
	if err := json.Unmarshal([]byte(session.AnsweredQuestionsJSON), &answeredQuestions); err != nil {
		utils.InternalServerError(c, "Error processing session answers: "+err.Error())
		return
	}

	answerDetail := models.AnsweredQuestionDetail{
		QuestionID:      question.ID,
		QuestionText:    question.Text,
		SelectedOptionID: selectedOption.ID,
		CorrectOptionID: correctOption.ID,
		IsCorrect:       isCorrect,
		Explanation:     correctOption.Explanation, // Provide explanation of the correct answer
	}
	answeredQuestions = append(answeredQuestions, answerDetail)

	updatedAnswersJSON, err := json.Marshal(answeredQuestions)
	if err != nil {
		utils.InternalServerError(c, "Error updating session answers: "+err.Error())
		return
	}
	session.AnsweredQuestionsJSON = string(updatedAnswersJSON)
	session.CurrentQuestion++ // Increment question count for all modes

	// Check game mode specific end conditions
	sessionActive := true
	switch session.Mode {
	case models.SurvivalMode:
		if session.MistakesMade >= session.MaxMistakes {
			session.IsActive = false
			session.EndedAt = time.Now()
			sessionActive = false
		}
	case models.FiveTopicMode:
		if session.CurrentQuestion >= session.TotalQuestions {
			session.IsActive = false
			session.EndedAt = time.Now()
			sessionActive = false
		}
	case models.InfiniteMode:
		// Infinite mode only ends via a separate API call (e.g., /session/finish) or client-side decision.
		// For MVP, we'll assume it continues until client stops or can add a manual Finish endpoint later.
		// If user wants to finish Infinite mode, they can tap Finish, which can then make a call to /session/:id/summary to end implicitly
		// Or we could add an explicit /session/:id/finish endpoint
		break
	}

	if err := dbInstance.Save(&session).Error; err != nil {
		utils.InternalServerError(c, "Could not update session: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":          isCorrect,
		"session_active":  sessionActive,
		"session_details": session,
		"correct_option_id": correctOption.ID,
		"explanation": correctOption.Explanation, // Provide immediate feedback
	})
}

// GetSessionSummary godoc
// @Summary Get a game session summary
// @Description Retrieves the summary of a completed or ongoing game session, including score and answer breakdown.
// @Tags sessions
// @Produce  json
// @Param   id  path   string  true  "Session ID"
// @Success 200 {object} map[string]interface{} "session: models.GameSession, answered_questions: []models.AnsweredQuestionDetail"
// @Failure 404 {object} utils.APIError "Session not found"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /session/{id}/summary [get]
func GetSessionSummary(c *gin.Context) {
	sessionID := c.Param("id")
	dbInstance := db.GetDB()
	var session models.GameSession

	if err := dbInstance.First(&session, "id = ?", sessionID).Error; err != nil {
		if db.IsErrNotFound(err) {
			utils.NotFound(c, "Session")
		} else {
			utils.InternalServerError(c, "Error fetching session: "+err.Error())
		}
		return
	}

	// If the session is infinite and still active, the client might be calling this to "finish" it.
	// Or, it's just a request for an interim summary.
	// For MVP, if it's an infinite mode and still marked active, we can mark it as ended now.
	if session.Mode == models.InfiniteMode && session.IsActive {
		session.IsActive = false
		session.EndedAt = time.Now()
		if err := dbInstance.Save(&session).Error; err != nil {
			utils.InternalServerError(c, "Error updating infinite session status: "+err.Error())
			return
		}
	}

	var answeredQuestions []models.AnsweredQuestionDetail
	if err := json.Unmarshal([]byte(session.AnsweredQuestionsJSON), &answeredQuestions); err != nil {
		utils.InternalServerError(c, "Error parsing answered questions: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session":            session,
		"answered_questions": answeredQuestions,
	})
} 
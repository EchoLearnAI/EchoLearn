package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/EchoLearnAI/EchoLearn/db"
	"github.com/EchoLearnAI/EchoLearn/models"
	"github.com/EchoLearnAI/EchoLearn/utils"
	"github.com/gin-gonic/gin"
)

// StartSessionRequest defines the expected body for starting a new session.
type StartSessionRequest struct {
	UserID string          `json:"user_id" binding:"required"`
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
		UserID:                req.UserID,
		Mode:                  req.Mode,
		IsActive:              true,       // A new session should be active
		StartedAt:             time.Now(), // Record start time
		AnsweredQuestionsJSON: "[]",
	}

	switch req.Mode {
	case models.GameModeMistakes:
		newSession.MaxMistakes = 3
	case models.GameModeCategoryChallenge:
		// Future: if req contains CategoryName, store it on newSession (model needs CategoryName field)
		// and the question fetching logic below should respect it.
		break
	case models.GameModeInfinite:
		break
	default:
		utils.BadRequest(c, "Invalid game mode specified: "+string(req.Mode))
		return
	}

	// Attempt to set an initial question for the session
	var firstQuestion models.Question
	// For simplicity, pick any random question.
	// TODO: For CategoryChallenge, this should be filtered by category if category is part of StartSessionRequest.
	if err := dbInstance.Order("RANDOM()").First(&firstQuestion).Error; err != nil {
		if !db.IsErrNotFound(err) { // If it's any error other than "not found"
			utils.InternalServerError(c, "Failed to fetch initial question: "+err.Error())
			return
		}
		// If no questions found, CurrentQuestionID will remain empty.
		// The frontend should handle this (e.g., "No questions available for this mode").
		// Or, we could prevent session creation if no questions exist.
		log.Println("StartSession: No questions found in DB to assign as first question.")
	} else {
		newSession.CurrentQuestionID = firstQuestion.ID
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
		QuestionID:       question.ID,
		QuestionText:     question.Text,
		SelectedOptionID: selectedOption.ID,
		CorrectOptionID:  correctOption.ID,
		IsCorrect:        isCorrect,
		Explanation:      correctOption.Explanation, // Provide explanation of the correct answer
	}
	answeredQuestions = append(answeredQuestions, answerDetail)

	updatedAnswersJSON, err := json.Marshal(answeredQuestions)
	if err != nil {
		utils.InternalServerError(c, "Error updating session answers: "+err.Error())
		return
	}
	session.AnsweredQuestionsJSON = string(updatedAnswersJSON)
	session.CurrentQuestion++ // Increment question count for all modes

	// Default to session still being active
	sessionActive := true

	// If session is still considered active, try to get a next question
	if session.IsActive { // Check if session is still active before attempting to fetch next question
		var nextQuestion models.Question
		// Exclude already answered questions if possible (more complex, for future)
		// For now, just get any random question not the current one (if possible)
		// This query attempts to get a random question that is not the one just answered.
		// If only one question exists, this might still return the same one or none if we strictly exclude.
		// A more robust solution would involve tracking all answered questions in the session.
		query := dbInstance.Order("RANDOM()")
		if req.QuestionID != "" { // req.QuestionID is the question just answered
			query = query.Not("id = ?", req.QuestionID)
		}

		err := query.First(&nextQuestion).Error
		if err != nil {
			if db.IsErrNotFound(err) {
				// No more questions available
				log.Printf("SubmitAnswer: No next question found for session %s. Ending session.", session.ID)
				session.IsActive = false
				session.EndedAt = time.Now()
				sessionActive = false          // Update local variable to reflect session ended
				session.CurrentQuestionID = "" // Clear current question ID
			} else {
				utils.InternalServerError(c, "Error fetching next question: "+err.Error())
				return
			}
		} else {
			session.CurrentQuestionID = nextQuestion.ID
		}
	}

	// Check game mode specific end conditions
	switch session.Mode {
	case models.GameModeMistakes:
		if session.MistakesMade >= session.MaxMistakes {
			session.IsActive = false
			session.EndedAt = time.Now()
			sessionActive = false
		}
	case models.GameModeCategoryChallenge:
		// Add end conditions for Category Challenge if applicable (e.g., number of questions)
		// if session.CurrentQuestion >= session.TotalQuestions { ... }
		break
	case models.GameModeInfinite:
		// Infinite mode only ends via a separate API call (e.g., /session/finish) or client-side decision.
		// For MVP, we'll assume it continues until client stops or can add a manual Finish endpoint later.
		// If user wants to finish Infinite mode, they can tap Finish, which can then make a call to /session/:id/summary to end implicitly
		// Or we could add an explicit /session/:id/finish endpoint

		// If session was marked inactive due to no more questions (above), ensure IsActive is false.
		if !session.IsActive { // This check handles the case where no more questions were found for infinite mode
			sessionActive = false
		}
		break
	}

	if err := dbInstance.Save(&session).Error; err != nil {
		utils.InternalServerError(c, "Could not update session: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":            isCorrect,
		"session_active":    sessionActive,
		"session_details":   session,
		"correct_option_id": correctOption.ID,
		"explanation":       correctOption.Explanation, // Provide immediate feedback
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
	if session.Mode == models.GameModeInfinite && session.IsActive {
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

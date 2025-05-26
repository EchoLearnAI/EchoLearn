package handlers

import (
	"net/http"

	"github.com/EchoLearnAI/EchoLearn/db"
	"github.com/EchoLearnAI/EchoLearn/models"
	errorUtils "github.com/EchoLearnAI/EchoLearn/utils"    // For API error responses
	passwordUtils "github.com/EchoLearnAI/EchoLearn/utils" // For password hashing
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginRequest defines the expected request body for logging in.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Log in a user
// @Description Authenticates a user with email and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   credentials body LoginRequest true "User login credentials"
// @Success 200 {object} models.User "Successfully logged in"
// @Failure 400 {object} errorUtils.APIError "Invalid request payload or validation error"
// @Failure 401 {object} errorUtils.APIError "Invalid email or password"
// @Failure 404 {object} errorUtils.APIError "User not found"
// @Failure 500 {object} errorUtils.APIError "Internal server error"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorUtils.BadRequest(c, "Invalid request payload: "+err.Error())
		return
	}

	dbInstance := db.GetDB()
	var user models.User

	if err := dbInstance.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errorUtils.NewErrorWithDetails(c, http.StatusUnauthorized, "Authentication Failed", "Invalid email or password")
		} else {
			errorUtils.InternalServerError(c, "Failed to retrieve user: "+err.Error())
		}
		return
	}

	if !passwordUtils.CheckPasswordHash(req.Password, user.PasswordHash) {
		errorUtils.NewErrorWithDetails(c, http.StatusUnauthorized, "Authentication Failed", "Invalid email or password")
		return
	}

	// Login successful, return user data (PasswordHash is excluded by json:"-")
	c.JSON(http.StatusOK, gin.H{"data": user})
}

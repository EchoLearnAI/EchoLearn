package handlers

import (
	"net/http"

	"github.com/EchoLearnAI/EchoLearn/db"
	"github.com/EchoLearnAI/EchoLearn/models"
	errorUtils "github.com/EchoLearnAI/EchoLearn/utils"    // For API error responses (assuming errors.go is in package utils)
	passwordUtils "github.com/EchoLearnAI/EchoLearn/utils" // For password hashing
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUserRequest defines the expected request body for creating a user.
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Creates a new user with name, email, and password.
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user body CreateUserRequest true "User creation request"
// @Success 201 {object} models.User "Successfully created user"
// @Failure 400 {object} errorUtils.APIError "Invalid request payload or validation error"
// @Failure 409 {object} errorUtils.APIError "Email already exists"
// @Failure 500 {object} errorUtils.APIError "Internal server error"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorUtils.BadRequest(c, "Invalid request payload: "+err.Error())
		return
	}

	hashedPassword, err := passwordUtils.HashPassword(req.Password)
	if err != nil {
		errorUtils.InternalServerError(c, "Failed to hash password")
		return
	}

	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	dbInstance := db.GetDB()
	if err := dbInstance.Create(&user).Error; err != nil {
		if db.IsUniqueConstraintError(err, "users.email") { // Check specific field if possible or just general unique error
			errorUtils.NewErrorWithDetails(c, http.StatusConflict, "Registration Error", "Email already exists")
		} else {
			errorUtils.InternalServerError(c, "Failed to create user: "+err.Error())
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Retrieves user details for a given user ID.
// @Tags users
// @Produce json
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} models.User "Successfully retrieved user"
// @Failure 400 {object} errorUtils.APIError "Invalid user ID format"
// @Failure 404 {object} errorUtils.APIError "User not found"
// @Failure 500 {object} errorUtils.APIError "Internal server error"
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	dbInstance := db.GetDB()
	if err := dbInstance.Preload("Sessions").First(&user, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errorUtils.NotFound(c, "User")
		} else {
			errorUtils.InternalServerError(c, "Failed to retrieve user: "+err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// TODO: Add UpdateUserProgress handler if needed, for now progress is tied to sessions.

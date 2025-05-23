package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/EchoLearnAI/EchoLearn/db"
	"github.com/EchoLearnAI/EchoLearn/models"
	"github.com/EchoLearnAI/EchoLearn/utils"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user to the system
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body   models.User  true  "User info"
// @Success 201 {object} models.User
// @Failure 400 {object} utils.APIError "Invalid input"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// Basic validation
	if user.Name == "" {
		utils.BadRequest(c, "User name cannot be empty")
		return
	}

	dbInstance := db.GetDB()
	if err := dbInstance.Create(&user).Error; err != nil {
		utils.InternalServerError(c, "Could not create user: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Retrieve user details by their ID
// @Tags users
// @Produce  json
// @Param   id  path   string  true  "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} utils.APIError "User not found"
// @Failure 500 {object} utils.APIError "Server error"
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	var user models.User
	dbInstance := db.GetDB()

	if err := dbInstance.First(&user, "id = ?", userID).Error; err != nil {
		if db.IsErrNotFound(err) {
			utils.NotFound(c, "User")
		} else {
			utils.InternalServerError(c, "Could not retrieve user: "+err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// TODO: Add UpdateUserProgress handler if needed, for now progress is tied to sessions. 
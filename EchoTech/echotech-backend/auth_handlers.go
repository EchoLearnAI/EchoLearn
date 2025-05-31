package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Use an environment variable in production!
var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY")) // Ensure JWT_SECRET_KEY is set

const tokenDuration = 24 * time.Hour // JWT token validity duration

type AuthCredentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body AuthCredentials true "User registration credentials"
// @Success 201 {object} map[string]interface{} "User created successfully (userId in response)"
// @Failure 400 {object} map[string]string "Invalid input or user already exists"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var creds AuthCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := User{Email: creds.Email, Password: string(hashedPassword)}

	// Check if user already exists
	var existingUser User
	if err := DB.Where("email = ?", creds.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error checking user"})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "userId": user.ID})
}

// Login godoc
// @Summary Login an existing user
// @Description Authenticates a user and returns a JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body AuthCredentials true "User login credentials"
// @Success 200 {object} map[string]string "Login successful, token provided"
// @Failure 400 {object} map[string]string "Invalid input or invalid credentials"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var creds AuthCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	var user User
	if err := DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error finding user"})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	expirationTime := time.Now().Add(tokenDuration)
	claims := &jwt.RegisteredClaims{
		Subject:   user.Email,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer:    "echotech-backend",
		ID:        fmt.Sprintf("%d", user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "userId": user.ID, "email": user.Email})
}

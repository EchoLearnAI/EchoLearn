package controllers

import (
	"echolearn/internal/db"
	"echolearn/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUser gets a user by ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	
	user := new(models.User)
	result := db.DB.First(&user, "id = ?", uid)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	
	return c.JSON(user)
}

// GetAllUsers gets all users
func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	
	result := db.DB.Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	
	return c.JSON(users)
}

// UpdateUser updates a user by ID
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	
	user := new(models.User)
	result := db.DB.First(&user, "id = ?", uid)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	
	db.DB.Save(&user)
	
	return c.JSON(user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	
	result := db.DB.Delete(&models.User{}, "id = ?", uid)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}
	
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
} 
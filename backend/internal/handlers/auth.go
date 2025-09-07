package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/fundament/internal/models"
	"github.com/yourusername/fundament/internal/utils"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	// Check if database is available
	if h.DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not available. Please ensure PostgreSQL is running.",
		})
	}

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	if len(req.Password) < 6 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Password must be at least 6 characters long",
		})
	}

	// Check if user already exists
	var existingUser models.User
	if err := h.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{
			"error": "User with this email already exists",
		})
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Create user
	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	// Check if database is available
	if h.DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not available. Please ensure PostgreSQL is running.",
		})
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate input
	if req.Email == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	// Debug logging
	log.Printf("🔐 Login attempt for email: %s", req.Email)

	// Find user
	var user models.User
	if err := h.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		log.Printf("❌ User not found: %s", req.Email)
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	log.Printf("✅ User found: %s (ID: %d)", user.Email, user.ID)

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		log.Printf("❌ Password mismatch for user: %s", req.Email)
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	log.Printf("✅ Password verified for user: %s", req.Email)

	// Generate JWT token
	token, err := utils.GenerateJWT(&user)
	if err != nil {
		log.Printf("❌ JWT generation failed for user: %s, error: %v", req.Email, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	log.Printf("✅ Login successful for user: %s", req.Email)

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

func (h *AuthHandler) Health(c *fiber.Ctx) error {
	// Check if database is available
	if h.DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"status": "unhealthy",
			"database": "not connected",
			"message": "Database not available. Please ensure PostgreSQL is running.",
		})
	}

	// Try a simple database query to ensure it's working
	var result int64
	if err := h.DB.Model(&models.User{}).Count(&result).Error; err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status": "unhealthy",
			"database": "error",
			"message": "Database query failed",
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "healthy",
		"database": "connected",
		"message": "Service is running normally",
	})
}

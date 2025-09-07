package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/fundament/internal/utils"
)

func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		log.Printf("🔐 JWT Middleware - Path: %s, Auth Header: %s", c.Path(), authHeader)

		if authHeader == "" {
			log.Printf("❌ JWT Middleware - Missing authorization header")
			return c.Status(401).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Printf("❌ JWT Middleware - Invalid authorization header format: %v", tokenParts)
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}

		tokenString := tokenParts[1]
		log.Printf("🔍 JWT Middleware - Validating token: %s...", tokenString[:20]+"...")

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			log.Printf("❌ JWT Middleware - Token validation failed: %v", err)
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		log.Printf("✅ JWT Middleware - Token validated for user: %s (ID: %d)", claims.Email, claims.UserID)

		// Store user information in context
		c.Locals("userID", claims.UserID)
		c.Locals("email", claims.Email)

		return c.Next()
	}
}

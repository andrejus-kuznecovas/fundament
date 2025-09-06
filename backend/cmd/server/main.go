package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/yourusername/fundament/internal/database"
	"github.com/yourusername/fundament/internal/handlers"
	"github.com/yourusername/fundament/internal/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)
	notesHandler := handlers.NewNotesHandler(db)

	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS_ORIGIN"),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// API routes
	api := app.Group("/api")

	// Authentication routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected notes routes
	notes := api.Group("/notes")
	notes.Use(middleware.JWTAuth())
	notes.Get("/", notesHandler.GetNotes)
	notes.Post("/", notesHandler.CreateNote)
	notes.Get("/:id", notesHandler.GetNote)
	notes.Put("/:id", notesHandler.UpdateNote)
	notes.Delete("/:id", notesHandler.DeleteNote)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

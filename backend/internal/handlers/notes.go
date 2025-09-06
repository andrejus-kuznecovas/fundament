package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/fundament/internal/models"
	"gorm.io/gorm"
)

type NotesHandler struct {
	DB *gorm.DB
}

type CreateNoteRequest struct {
	Content string `json:"content" validate:"required"`
}

type UpdateNoteRequest struct {
	Content string `json:"content" validate:"required"`
}

func NewNotesHandler(db *gorm.DB) *NotesHandler {
	return &NotesHandler{DB: db}
}

func (h *NotesHandler) GetNotes(c *fiber.Ctx) error {
	// Check if database is available
	if h.DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not available. Please ensure PostgreSQL is running.",
		})
	}

	userID := c.Locals("userID").(uint)

	var notes []models.Note
	if err := h.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&notes).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve notes",
		})
	}

	return c.JSON(fiber.Map{
		"notes": notes,
	})
}

func (h *NotesHandler) CreateNote(c *fiber.Ctx) error {
	// Check if database is available
	if h.DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not available. Please ensure PostgreSQL is running.",
		})
	}

	userID := c.Locals("userID").(uint)

	var req CreateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Content == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Content is required",
		})
	}

	note := models.Note{
		UserID:  userID,
		Content: req.Content,
	}

	if err := h.DB.Create(&note).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create note",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"note": note,
	})
}

func (h *NotesHandler) GetNote(c *fiber.Ctx) error {
	// Check if database is available
	if h.DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not available. Please ensure PostgreSQL is running.",
		})
	}

	userID := c.Locals("userID").(uint)
	noteIDStr := c.Params("id")

	noteID, err := strconv.ParseUint(noteIDStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid note ID",
		})
	}

	var note models.Note
	if err := h.DB.Where("id = ? AND user_id = ?", uint(noteID), userID).First(&note).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	return c.JSON(fiber.Map{
		"note": note,
	})
}

func (h *NotesHandler) UpdateNote(c *fiber.Ctx) error {
	// Check if database is available
	if h.DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not available. Please ensure PostgreSQL is running.",
		})
	}

	userID := c.Locals("userID").(uint)
	noteIDStr := c.Params("id")

	noteID, err := strconv.ParseUint(noteIDStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid note ID",
		})
	}

	var req UpdateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Content == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Content is required",
		})
	}

	var note models.Note
	if err := h.DB.Where("id = ? AND user_id = ?", uint(noteID), userID).First(&note).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	note.Content = req.Content
	if err := h.DB.Save(&note).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update note",
		})
	}

	return c.JSON(fiber.Map{
		"note": note,
	})
}

func (h *NotesHandler) DeleteNote(c *fiber.Ctx) error {
	// Check if database is available
	if h.DB == nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "Database not available. Please ensure PostgreSQL is running.",
		})
	}

	userID := c.Locals("userID").(uint)
	noteIDStr := c.Params("id")

	noteID, err := strconv.ParseUint(noteIDStr, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid note ID",
		})
	}

	var note models.Note
	if err := h.DB.Where("id = ? AND user_id = ?", uint(noteID), userID).First(&note).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Note not found",
		})
	}

	if err := h.DB.Delete(&note).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete note",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}

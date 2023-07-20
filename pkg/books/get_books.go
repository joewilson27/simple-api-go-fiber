package books

import (
	"github.com/gofiber/fiber/v2"
	"simple-go-fiber/pkg/common/models"
)

// this method belongs to struct handler in controller
func (h handler) GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}
package books

import (
	"github.com/gofiber/fiber/v2"
	"simple-go-fiber/pkg/common/models"
)

// this method belongs to struct handler in controller
func (h handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete book from db
	h.DB.Delete(&book)

	return c.SendStatus(fiber.StatusOK)
}
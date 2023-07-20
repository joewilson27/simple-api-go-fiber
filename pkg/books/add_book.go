package books

import (
	"github.com/gofiber/fiber/v2"
  "simple-go-fiber/pkg/common/models"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// this method belongs to struct handler in controller
func (h handler) AddBook(c *fiber.Ctx) error {
	body := new(AddBookRequestBody)

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	book := new(models.Book)
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	// insert new db entry
	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&book)

}
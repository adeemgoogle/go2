package GET

import (
	"github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func AuthorsBook(c *fiber.Ctx) error {
	authorId := c.Params("id")

	var books []models.Book
	err := database.Db.Select(&books, "Select id, title, genre, isbn  from book where authorid=$1", authorId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "do not work select",
		})
	}

	return c.Status(200).JSON(authorId)
}

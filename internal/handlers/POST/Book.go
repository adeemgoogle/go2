package POST

import (
	"github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {
	var err error
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Not right format",
		})
	}

	_, err = database.Db.Exec("INSERT INTO book(id, title, genre, isbn, authorid) VALUES ($1, $2, $3, $4, $5 )", book.ID, book.Title, book.Genre, book.ISBN, book.AuthorID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create book",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "created",
	})
}

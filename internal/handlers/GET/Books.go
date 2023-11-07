package GET

import (
	"github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	rows, err := database.Db.Query("SELECT id, title, genre, isbn, authorid FROM book")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Genre, &book.ISBN, &book.AuthorID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

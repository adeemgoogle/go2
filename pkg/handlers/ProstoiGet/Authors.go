package ProstoiGet

import (
	"github.com/adeemgoogle/go2/pkg/database"
	"github.com/adeemgoogle/go2/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetAuthors(c *fiber.Ctx) error {
	rows, err := database.Db.Query("SELECT id, fullname, pseudonym, spesialization FROM author")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	defer rows.Close()

	var authors []models.Author

	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.FullName, &author.Pseudonym, &author.Specialization); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	return c.Status(fiber.StatusOK).JSON(authors)
}

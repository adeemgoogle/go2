package GET

import (
	models2 "github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func GetAuthors(c *fiber.Ctx) error {
	rows, err := database.Db.Query("SELECT id, fullname, pseudonym, specialization FROM author")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	defer rows.Close()

	var authors []models2.Author

	for rows.Next() {
		var author models2.Author
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

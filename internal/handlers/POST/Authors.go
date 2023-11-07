package POST

import (
	"github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func CreateAuthor(c *fiber.Ctx) error {
	var err error
	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Not right format",
		})
	}

	_, err = database.Db.Exec("INSERT INTO author(id, fullname, pseudonym, specialization) VALUES ($1, $2, $3, $4)", author.ID, author.FullName, author.Pseudonym, author.Specialization)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create authors",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "created",
		"author":  author,
	})
}

package DELETE

import (
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func AuthorDelete(c *fiber.Ctx) error {
	authorID := c.Params("id")
	_, err := database.Db.Exec("DELETE  FROM book where authorid = $1", authorID)
	if err != nil {
		return c.Status(500).SendString("Error with deleting")
	}

	_, err = database.Db.Exec("DELETE FROM author where id = $1", authorID)
	if err != nil {
		return c.Status(500).SendString("Error with deleting")
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Deleted",
	})
}

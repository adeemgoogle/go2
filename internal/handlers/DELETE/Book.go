package DELETE

import (
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func BookDelete(c *fiber.Ctx) error {
	bookId := c.Params("id")
	var err error

	if err != nil {
		return c.Status(400).SendString("not correct id")
	}
	_, err = database.Db.Query("DELETE  FROM book where id = $1", bookId)
	if err != nil {
		return c.Status(500).SendString("Error with deleting")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Deleted",
	})
}

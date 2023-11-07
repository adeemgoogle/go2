package POST

import (
	"github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func CreateBRbook(c *fiber.Ctx) error {
	var err error
	BRbook := new(models.BorrowedBook)

	if err := c.BodyParser(BRbook); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Not right format",
		})
	}

	_, err = database.Db.Exec("INSERT INTO borrowedbooks(memberid, bookid) VALUES ($1, $2 )", BRbook.MemberID, BRbook.BookID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create borrewed book",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "created",
	})
}

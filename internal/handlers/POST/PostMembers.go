package POST

import (
	"github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func CreateMember(c *fiber.Ctx) error {
	var err error
	member := new(models.Member)

	if err := c.BodyParser(member); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Not right format",
		})
	}

	_, err = database.Db.Exec("INSERT INTO members(memberid, fullname) VALUES ($1, $2)", member.MemberID, member.FullName)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create member",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "created",
	})
}

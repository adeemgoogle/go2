package GET

import (
	"github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func GetMembers(c *fiber.Ctx) error {
	rows, err := database.Db.Query("SELECT memberid, fullname FROM members")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	defer rows.Close()

	var members []models.Member

	for rows.Next() {
		var member models.Member
		if err := rows.Scan(&member.MemberID, &member.FullName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	return c.Status(fiber.StatusOK).JSON(members)
}

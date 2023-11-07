package PATCH

import "github.com/gofiber/fiber/v2"

func UpdAuthors(c *fiber.Ctx) error {
	var upd map[string]interface{}
	if err := c.BodyParser(&upd); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("not right data")
	}

	query := "UPDATE AUTHOR SET"
	args := []interface{}{}
	for key, value := range upd {
		query += " " + key + "=?,"
		args = append(args, value)
	}
	query += "where id = ? "
	//args = append(args, au)

	return c.Status(200).JSON(fiber.Map{
		"message": "Deleted",
	})
}

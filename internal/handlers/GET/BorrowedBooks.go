package GET

import (
	"github.com/adeemgoogle/go2/internal/models"
	"github.com/adeemgoogle/go2/internal/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func BorrowedBooks(c *fiber.Ctx) error {
	rows, err := database.Db.Query("SELECT member_id, book_id FROM borrowedbooks")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	defer rows.Close()

	var borrowedBook []models.BorrowedBook

	for rows.Next() {
		var borrowedBooks models.BorrowedBook
		if err := rows.Scan(&borrowedBooks.MemberID, &borrowedBooks.BookID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve authors",
			})
		}
		borrowedBook = append(borrowedBook, borrowedBooks)
	}

	if err := rows.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve authors",
		})
	}
	return c.Status(fiber.StatusOK).JSON(borrowedBook)
}

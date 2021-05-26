package user

import (
	"time"
	"upvoteTest/database"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}

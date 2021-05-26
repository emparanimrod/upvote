package user

import (
	"time"
	"upvoteTest/database"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint      `gorm:"primaryKey; autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" gorm:"string"`
	UpdatedAt time.Time `json:"updated_at" gorm:"string"`
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user User
	payload := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(err)
	}

	if payload.Name == "" || payload.Email == "" {
		c.JSON("Fill in all required fields")
		return c.SendStatus(401)
	}
	user.Name = payload.Name
	user.Email = payload.Email
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := db.Create(&user).Error
	if err != nil {
		return c.Status(401).JSON("invalid data")
	}
	return c.JSON("User Created")
}

package main

import (
	"fmt"
	"upvoteTest/database"
	"upvoteTest/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}

func initDatabase() {
	var err error
	dsn := "upvote:clarion103@tcp(127.0.0.1:3306)/upvote_test?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen(":3000")
}

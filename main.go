package main

import (
	"fmt"
	"upvoteTest/database"
	"upvoteTest/post"
	"upvoteTest/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	// Grouped API according to the naming convention
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	//	get users
	v1.Get("/users", user.GetUsers)
	//	Create User
	v1.Post("/user", user.CreateUser)
	// Get Posts
	v1.Get("/posts", post.GetPosts)
	// Create Post
	v1.Post("/post", post.CreatePost)
	//	Upvote
	v1.Get("/upvote/:id", post.Upvote)
	//	Downvote
	v1.Get("/downvote/:id", post.DownVote)
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
	database.DBConn.AutoMigrate(&post.Post{})
	fmt.Println("Database Migrated")
}

func Setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	setupRoutes(app)
	app.Use(logger.New())
	return app
}

func main() {
	app := Setup()

	app.Listen(":3000")
}

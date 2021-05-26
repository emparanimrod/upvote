package main

import (
	"fmt"
	"log"
	"os"
	"upvoteTest/database"
	"upvoteTest/post"
	"upvoteTest/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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
	// Get Post
	v1.Get("/post/:id", post.GetPost)
	// Create Post
	v1.Post("/post", post.CreatePost)
	//	Upvote
	v1.Get("/upvote", post.Upvote)
	//	Downvote
	v1.Get("/downvote", post.DownVote)
}

func initDatabase() {
	var err error
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPW := os.Getenv("DB_PASSWORD")
	dbURL := os.Getenv("DB_URL")

	//dsn := "upvote:clarion103@tcp(127.0.0.1:3306)/upvote_test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPW, dbURL, dbName)
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
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not Found")
	})
	return app
}

func main() {
	app := Setup()

	app.Listen(":3000")
}

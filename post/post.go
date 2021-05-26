package post

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"upvoteTest/database"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Upvotes   int       `json:"upvotes"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func GetPosts(c *fiber.Ctx) error {
	db := database.DBConn
	var posts []Post
	db.Preload("Author").Find(&posts)
	return c.JSON(&posts)
}

func CreatePost(c *fiber.Ctx) error {
	db := database.DBConn
	var post Post

	if err := c.BodyParser(&post); err != nil {
		return c.JSON(err)
	}

	post.Upvotes = 0

	err := db.Create(&post).Error
	if err != nil {
		return c.Status(401).JSON(err)
	}
	return c.JSON(post)
}

func Upvote(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var post Post

	err := db.Where("id = ?", id).Find(&post).Error
	if err != nil {
		return c.Status(500).JSON("An Error Occurred")
	}
	if post.ID == 0 {
		return c.Status(404).JSON("Post Does not Exist")
	}
	totalVotes := post.Upvotes + 1
	error := db.Model(&post).Update("upvotes", totalVotes).Error
	if error != nil {
		return c.Status(500).JSON("An Error Occurred")
	}
	return c.JSON(post)
}

func DownVote(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var post Post

	err := db.Where("id = ?", id).Find(&post).Error
	if err != nil {
		return c.Status(500).JSON("An Error Occurred")
	}
	if post.ID == 0 {
		return c.Status(404).JSON("Post Does not Exist")
	}
	totalVotes := post.Upvotes - 1
	error := db.Model(&post).Update("upvotes", totalVotes).Error
	if error != nil {
		return c.Status(500).JSON("An Error Occurred")
	}
	return c.JSON(post)
}

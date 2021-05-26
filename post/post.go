package post

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"upvoteTest/database"
	"upvoteTest/user"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Upvotes   int       `json:"upvotes" gorm:"default:0"`
	VotesCast int       `json:"votes_cast" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func GetPosts(c *fiber.Ctx) error {
	db := database.DBConn
	var posts []Post
	db.Preload("Author").Find(&posts)
	return c.JSON(&posts)
}

func GetPost(c *fiber.Ctx) error {
	db := database.DBConn
	var post Post
	id := c.Params("id")
	err := db.First(&post, id).Error
	if err != nil {
		return c.Status(404).JSON("Couldn't find that")
	}
	return c.JSON(post)
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
	id := c.Query("post")
	userId := c.Query("user")
	var post Post
	var user user.User

	//check if user exists
	userErr := db.First(&user, "id = ?", userId).Error
	if userErr != nil {
		return c.Status(401).JSON("Are you a valid user?")
	}
	//check for post and update
	err := db.Where("id = ?", id).Find(&post).Error
	if err != nil {
		return c.Status(500).JSON("An Error Occurred")
	}
	if post.ID == 0 {
		return c.Status(404).JSON("Post Does not Exist")
	}
	totalVotes := post.Upvotes + 1
	totalVotesCast := post.VotesCast + 1
	error := db.Model(&post).Update("upvotes", totalVotes).Update("votes_cast", totalVotesCast).Error
	if error != nil {
		return c.Status(500).JSON("An Error Occurred")
	}
	return c.JSON(post)
}

func DownVote(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Query("post")
	userId := c.Query("user")
	var post Post
	var user user.User

	//check if user exists
	userErr := db.First(&user, "id = ?", userId).Error
	if userErr != nil {
		return c.Status(401).JSON("Are you a valid user?")
	}
	//check for post and update
	err := db.Where("id = ?", id).Find(&post).Error
	if err != nil {
		return c.Status(500).JSON("An Error Occurred")
	}
	if post.ID == 0 {
		return c.Status(404).JSON("Post Does not Exist")
	}
	totalVotes := post.Upvotes - 1
	totalVotesCast := post.VotesCast + 1
	error := db.Model(&post).Update("upvotes", totalVotes).Update("votes_cast", totalVotesCast).Error
	if error != nil {
		return c.Status(500).JSON("An Error Occurred")
	}
	return c.JSON(post)
}

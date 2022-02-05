package models

import (
	"log"
	"time"

	"github.com/paulobezerra/goblog/src/configs"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id         int       `json:"id"`
	Slug       string    `json:"slug"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Abstract   string    `json:"abstract"`
	AuthorId   string    `json:"author_id"`
	Author     User      `json:"author" gorm:"foreignKey:AuthorId"`
	CategoryId int       `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryId"`
	Tags       []Tag     `gorm:"many2many:posts_tags;"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func GetPost(id int) *Post {
	db := configs.GetConnect()

	var post Post
	db.First(&post, id)

	return &post
}

func FindOnePostBySlug(slug string) *Post {
	db := configs.GetConnect()

	var post Post
	db.First(&post, "slug = ?", slug)
	return &post
}

func FindAllPosts() []Post {
	db := configs.GetConnect()

	var posts []Post
	db.Find(&posts)

	return posts
}

func (post *Post) Create() bool {
	db := configs.GetConnect()
	if err := db.Create(&post).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (post *Post) Save() bool {
	db := configs.GetConnect()
	if err := db.Save(&post).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (post *Post) Delete() bool {
	db := configs.GetConnect()
	if err := db.Delete(&post).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

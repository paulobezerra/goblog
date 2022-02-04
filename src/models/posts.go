package models

import (
	"time"

	"github.com/paulobezerra/goblog/src/db"
)

type Post struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      string    `json:"tags"`
	Abstract  string    `json:"abstract"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetPost(id string) *Post {
	db := db.GetConnect()

	var post Post
	db.First(&post, id)

	return &post
}

func FindOnePostBySlug(slug string) *Post {
	db := db.GetConnect()

	var post Post
	db.First(&post, "slug = ?", slug)

	return &post
}

func FindAllPosts() []Post {
	db := db.GetConnect()

	var posts []Post
	db.Find(&posts)

	return posts
}

func CreatePost(slug string, title string, abstract string, content string, tags string) Post {
	db := db.GetConnect()

	post := Post{Slug: slug, Title: title, Abstract: abstract, Content: content, Tags: tags}

	db.Create(&post)

	return post
}

func UpdatePost(id string, slug string, title string, abstract string, content string, tags string) Post {
	db := db.GetConnect()

	var post Post
	db.First(&post, id)
	post.Slug = slug
	post.Title = title
	post.Abstract = abstract
	post.Content = content
	post.Tags = tags

	db.Save(&post)

	return post
}

func DeletePost(id string) {
	db := db.GetConnect()

	var post Post
	db.Delete(&post, id)
}

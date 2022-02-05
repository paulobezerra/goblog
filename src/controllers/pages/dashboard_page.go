package pages

import "github.com/paulobezerra/goblog/src/models"

type Dashboard struct {
	User        models.User
	PostsActive string
	UsersActive string
}

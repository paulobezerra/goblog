package dto

import "github.com/paulobezerra/goblog/src/models"

type DashboardDto struct {
	User             models.User
	PostsActive      string
	UsersActive      string
	CategoriesActive string
}

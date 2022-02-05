package pages

import "github.com/paulobezerra/goblog/src/models"

type UserIndexPage struct {
	Dashboard
	Users []models.User
}

func NewUserIndexPage(user models.User, users []models.User) UserIndexPage {
	return UserIndexPage{
		Dashboard: Dashboard{
			User: user,
		},
		Users: users,
	}
}

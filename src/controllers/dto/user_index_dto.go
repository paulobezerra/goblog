package dto

import "github.com/paulobezerra/goblog/src/models"

type UserIndexDto struct {
	DashboardDto
	Users []models.User
}

func NewUserIndexDto(user models.User, users []models.User) UserIndexDto {
	return UserIndexDto{
		DashboardDto: DashboardDto{
			User: user,
		},
		Users: users,
	}
}

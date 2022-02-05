package dto

import (
	"net/http"
	"strconv"

	"github.com/paulobezerra/goblog/src/models"
)

type UserDto struct {
	FormDto
	DashboardDto
	models.User
}

func NewUserDto(title string, user models.User) UserDto {
	return UserDto{
		FormDto: FormDto{
			FormTitle:          title,
			ValidationMessages: map[string]string{},
		},
		DashboardDto: DashboardDto{
			UsersActive: "active",
			User:        user,
		},
	}
}

func (form *UserDto) SetUserId(id string) {
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt
}

func (form *UserDto) LoadFormData(r *http.Request) {
	form.Username = r.FormValue("username")
	form.Firstname = r.FormValue("firstname")
	form.Lastname = r.FormValue("lastname")
	form.Password = r.FormValue("password")
}

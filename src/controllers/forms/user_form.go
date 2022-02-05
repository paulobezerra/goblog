package forms

import (
	"net/http"
	"strconv"

	"github.com/paulobezerra/goblog/src/controllers/pages"
	"github.com/paulobezerra/goblog/src/models"
)

type UserFormData struct {
	Form
	pages.Dashboard
	models.User
}

func NewUserFormData(title string, user models.User) UserFormData {
	return UserFormData{
		Form: Form{
			FormTitle:          title,
			ValidationMessages: map[string]string{},
		},
		Dashboard: pages.Dashboard{
			UsersActive: "active",
			User:        user,
		},
	}
}

func (form *UserFormData) GetUser() models.User {
	var user models.User

	if form.Id > 0 {
		user.Id = form.Id
	}

	user.Username = form.Username
	user.Firstname = form.Firstname
	user.Lastname = form.Lastname
	user.Password = form.Password

	return user
}

func (form *UserFormData) SetUser(user models.User) {

	if user.Id > 0 {
		form.Id = user.Id
	}

	form.Username = user.Username
	form.Firstname = user.Firstname
	form.Lastname = user.Lastname
}

func (form *UserFormData) SetUserId(id string) {
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt
}

func (form *UserFormData) LoadFormData(r *http.Request) {
	form.Username = r.FormValue("username")
	form.Firstname = r.FormValue("firstname")
	form.Lastname = r.FormValue("lastname")
	form.Password = r.FormValue("password")
}

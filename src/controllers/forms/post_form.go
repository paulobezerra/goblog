package forms

import (
	"net/http"
	"strconv"

	"github.com/paulobezerra/goblog/src/controllers/pages"
	"github.com/paulobezerra/goblog/src/models"
)

type PostFormData struct {
	Form
	pages.Dashboard
	models.Post
}

func NewPostFormData(title string, user models.User) PostFormData {
	return PostFormData{
		Form: Form{
			FormTitle:          title,
			ValidationMessages: map[string]string{},
		},
		Dashboard: pages.Dashboard{
			PostsActive: "active",
			User:        user,
		},
	}
}

func (form *PostFormData) SetPostId(id string) {
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt
}

func (form *PostFormData) LoadFormData(r *http.Request) {
	form.Slug = r.FormValue("slug")
	form.Title = r.FormValue("title")
	form.Abstract = r.FormValue("abstract")
	form.Content = r.FormValue("content")
	form.Tags = r.FormValue("tags")
}

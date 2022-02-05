package dto

import (
	"net/http"
	"strconv"

	"github.com/paulobezerra/goblog/src/models"
)

type PostDto struct {
	FormDto
	DashboardDto
	models.Post
}

func NewPostDto(title string, user models.User) PostDto {
	return PostDto{
		FormDto: FormDto{
			FormTitle:          title,
			ValidationMessages: map[string]string{},
		},
		DashboardDto: DashboardDto{
			PostsActive: "active",
			User:        user,
		},
	}
}

func (form *PostDto) SetPostId(id string) {
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt
}

func (form *PostDto) LoadFormData(r *http.Request) {
	form.Slug = r.FormValue("slug")
	form.Title = r.FormValue("title")
	form.Abstract = r.FormValue("abstract")
	form.Content = r.FormValue("content")
	form.Tags = r.FormValue("tags")
}

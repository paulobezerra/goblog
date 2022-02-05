package dto

import (
	"net/http"
	"strconv"

	"github.com/paulobezerra/goblog/src/models"
)

type TagDto struct {
	FormDto
	DashboardDto
	models.Tag
}

func NewTagDto(title string, user models.User) TagDto {
	return TagDto{
		FormDto: FormDto{
			FormTitle:          title,
			ValidationMessages: map[string]string{},
		},
		DashboardDto: DashboardDto{
			TagsActive: "active",
			User:       user,
		},
	}
}

func (form *TagDto) SetTagId(id string) {
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt
}

func (form *TagDto) LoadFormData(r *http.Request) {
	form.Description = r.FormValue("description")
}

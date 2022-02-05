package dto

import (
	"net/http"
	"strconv"

	"github.com/paulobezerra/goblog/src/models"
)

type CategoryDto struct {
	FormDto
	DashboardDto
	models.Category
}

func NewCategoryDto(title string, user models.User) CategoryDto {
	return CategoryDto{
		FormDto: FormDto{
			FormTitle:          title,
			ValidationMessages: map[string]string{},
		},
		DashboardDto: DashboardDto{
			CategoriesActive: "active",
			User:             user,
		},
	}
}

func (form *CategoryDto) SetCategoryId(id string) {
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt
}

func (form *CategoryDto) LoadFormData(r *http.Request) {
	form.Description = r.FormValue("description")
}

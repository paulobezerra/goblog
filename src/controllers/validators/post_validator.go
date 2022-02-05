package validators

import (
	"github.com/paulobezerra/goblog/src/controllers/dto"
	"github.com/paulobezerra/goblog/src/models"
)

func ValidatePost(form *dto.PostDto) bool {
	var valid = true

	if form.Slug == "" {
		form.ValidationMessages["Slug"] = "Slug deve ser informado"
		valid = false
	}

	if form.Title == "" {
		form.ValidationMessages["Title"] = "Título deve ser informado"
		valid = false
	}

	if form.Abstract == "" {
		form.ValidationMessages["Abstract"] = "Resumo deve ser informado"
		valid = false
	}

	// if form.Tags == "" {
	// 	form.ValidationMessages["Tags"] = "Tags deve ser informado"
	// 	valid = false
	// }

	post := models.FindOnePostBySlug(form.Slug)
	if post.Id > 0 && post.Id != form.Id {
		form.ValidationMessages["Slug"] = "Já existe um post com este slug"
		valid = false
	}

	return valid
}

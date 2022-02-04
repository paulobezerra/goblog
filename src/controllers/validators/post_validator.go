package validators

import (
	"strconv"

	"github.com/paulobezerra/goblog/src/models"
)

func ValidatePost(id string, slug string, title string, abstract string, content string, tags string) (map[string]string, bool) {
	var messages = map[string]string{}
	var valid = true

	if slug == "" {
		messages["Slug"] = "Slug deve ser informado"
		valid = false
	}

	if title == "" {
		messages["Title"] = "Título deve ser informado"
		valid = false
	}

	if abstract == "" {
		messages["Abstract"] = "Resumo deve ser informado"
		valid = false
	}

	if tags == "" {
		messages["Tags"] = "Tags deve ser informado"
		valid = false
	}

	post := models.FindOnePostBySlug(slug)
	idInt, _ := strconv.Atoi(id)
	if post.Id > 0 && post.Id != idInt {
		messages["slug"] = "Já existe um post com este slug"
		valid = false
	}

	return messages, valid
}

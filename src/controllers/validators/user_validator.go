package validators

import (
	"github.com/paulobezerra/goblog/src/controllers/dto"
	"github.com/paulobezerra/goblog/src/models"
)

func ValidateUser(form dto.UserDto) bool {
	var valid = true

	if form.Username == "" {
		form.ValidationMessages["Username"] = "Nome do usuário deve ser informado"
		valid = false
	}

	if form.Firstname == "" {
		form.ValidationMessages["Firstname"] = "Primeiro nome deve ser informado"
		valid = false
	}

	if form.Lastname == "" {
		form.ValidationMessages["Lastname"] = "Sobrenome deve ser informado"
		valid = false
	}

	user := models.FindOneUserByUsername(form.Username)
	if user.Id != 0 && form.Id != user.Id {
		form.ValidationMessages["Username"] = "Já existe um usuário com este nome"
		valid = false
	}

	if form.Id == 0 && form.Password == "" {
		form.ValidationMessages["Password"] = "Senha deve ser informado"
		valid = false
	}

	if form.Password != "" && len(form.Password) < 6 {
		form.ValidationMessages["Password"] = "Senha deve conter no mínimo 6 caracteres"
		valid = false
	}

	return valid
}

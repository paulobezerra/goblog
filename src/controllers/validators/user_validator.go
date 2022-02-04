package validators

import "github.com/paulobezerra/goblog/src/models"

func ValidateUser(username string, firstname string, lastname string, password string, update bool) (map[string]string, bool) {
	var messages = map[string]string{}
	var valid = true

	if username == "" {
		messages["Username"] = "Nome do usuário deve ser informado"
		valid = false
	}

	if firstname == "" {
		messages["Firstname"] = "Primeiro nome deve ser informado"
		valid = false
	}

	if lastname == "" {
		messages["Lastname"] = "Sobrenome deve ser informado"
		valid = false
	}
	if !update {
		user := models.FindOneUserByUsername(username)
		if user.Id != 0 {
			messages["Username"] = "Já existe um usuário com este nome"
			valid = false
		}

		if password == "" {
			messages["Password"] = "Senha deve ser informado"
			valid = false
		} else if len(password) < 6 {
			messages["Password"] = "Senha deve conter no mínimo 6 caracteres"
			valid = false
		}
	}

	return messages, valid
}

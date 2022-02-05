package validators

import (
	"fmt"

	"github.com/paulobezerra/goblog/src/controllers/dto"
)

func ValidateCategory(form *dto.CategoryDto) bool {
	var valid = true
	fmt.Println(form)
	if form.Description == "" {
		form.ValidationMessages["Description"] = "Descrição deve ser informada"
		valid = false
	}

	return valid
}

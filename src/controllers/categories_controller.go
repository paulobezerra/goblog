package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/paulobezerra/goblog/src/controllers/dto"
	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/controllers/validators"
	"github.com/paulobezerra/goblog/src/models"
)

func IndexCategories(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	categories := models.FindAllCategories()
	data := map[string]interface{}{
		"CategoriesActive": "active",
		"User":             user,
		"Categories":       categories,
	}
	helpers.RenderTemplate(w, "layout_admin", "categories/index", data)
}

func FormCreateCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	data := dto.NewCategoryDto("Nova categoria", user)
	helpers.RenderTemplate(w, "layout_admin", "categories/form", data)
}

func CreateCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewCategoryDto("Nova categoria", user)
	form.LoadFormData(r)
	if validators.ValidateCategory(&form) && form.Category.Create() {
		http.Redirect(w, r, "/admin/categories", http.StatusFound)
		return
	}
	helpers.RenderTemplate(w, "layout_admin", "categories/form", form)
}

func FormUpdateCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewCategoryDto("Editar categoria", user)
	form.SetCategoryId(p.ByName("id"))
	category := models.GetCategory(form.Id)
	if category.Id == 0 {
		http.Redirect(w, r, "/admin/categories", http.StatusFound)
		return
	}
	form.Category = *category
	helpers.RenderTemplate(w, "layout_admin", "categories/form", form)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewCategoryDto("Editar categoria", user)
	form.SetCategoryId(p.ByName("id"))
	form.LoadFormData(r)
	if validators.ValidateCategory(&form) && form.Category.Save() {
		http.Redirect(w, r, "/admin/categories", http.StatusFound)
		return
	}
	helpers.RenderTemplate(w, "layout_admin", "categories/form", form)
}

func ViewCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewCategoryDto("Dados da categoria", user)
	form.SetCategoryId(p.ByName("id"))
	category := models.GetCategory(form.Id)
	if category.Id == 0 {
		http.Redirect(w, r, "/admin/categories", http.StatusFound)
		return
	}
	form.Category = *category
	helpers.RenderTemplate(w, "layout_admin", "categories/view", form)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	id := p.ByName("id")
	idInt, _ := strconv.Atoi(id)
	category := models.GetCategory(idInt)
	fmt.Println(id, idInt, category)
	if category.Id > 0 {
		category.Delete()
	}
	http.Redirect(w, r, "/admin/categories", http.StatusFound)
}

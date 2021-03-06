package controllers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/paulobezerra/goblog/src/controllers/dto"
	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/controllers/validators"
	"github.com/paulobezerra/goblog/src/models"
)

func IndexUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	users := models.FindAllUsers()
	page := dto.NewUserIndexDto(userAuth, users)
	helpers.RenderTemplate(w, "layout_admin", "users/index", page)
}

func FormCreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	data := dto.NewUserDto("Novo usuário", user)
	helpers.RenderTemplate(w, "layout_admin", "users/form", data)
}

func CreateUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	form := dto.NewUserDto("Novo usuário", userAuth)

	form.LoadFormData(r)

	if validators.ValidateUser(form) && form.User.Create() {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
	}

	helpers.RenderTemplate(w, "layout_admin", "users/form", form)
}

func FormUpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	form := dto.NewUserDto("Editar usuário", userAuth)
	form.SetUserId(p.ByName("id"))
	user := models.GetUser(form.Id)
	if user.Id == 0 {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
		return
	}
	form.User = user
	helpers.RenderTemplate(w, "layout_admin", "users/form", form)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	form := dto.NewUserDto("Novo usuário", userAuth)

	form.SetUserId(p.ByName("id"))
	form.LoadFormData(r)

	if validators.ValidateUser(form) && form.User.Save() {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
		return
	}

	helpers.RenderTemplate(w, "layout_admin", "users/form", form)
}

func ViewUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	form := dto.NewUserDto("Editar usuário", userAuth)
	form.SetUserId(p.ByName("id"))
	user := models.GetUser(form.Id)
	if user.Id == 0 {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
		return
	}
	form.User = user
	helpers.RenderTemplate(w, "layout_admin", "users/view", form)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	id := p.ByName("id")
	idInt, _ := strconv.Atoi(id)
	user := models.GetUser(idInt)
	if user.Id > 0 {
		user.Delete()
	}
	http.Redirect(w, r, "/admin/users", http.StatusFound)
}

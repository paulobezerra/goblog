package controllers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/paulobezerra/goblog/src/controllers/forms"
	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/controllers/pages"
	"github.com/paulobezerra/goblog/src/controllers/validators"
	"github.com/paulobezerra/goblog/src/models"
)

func IndexUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	users := models.FindAllUsers()
	page := pages.NewUserIndexPage(userAuth, users)
	helpers.RenderTemplate(w, "layout_admin", "users/index", page)
}

func FormCreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	data := forms.NewUserFormData("Novo usuário", user)
	helpers.RenderTemplate(w, "layout_admin", "users/form", data)
}

func CreateUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	form := forms.NewUserFormData("Novo usuário", userAuth)

	form.LoadFormData(r)

	user := form.GetUser()
	if validators.ValidateUser(form) && user.Create() {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
	}

	helpers.RenderTemplate(w, "layout_admin", "users/form", form)
}

func FormUpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	form := forms.NewUserFormData("Editar usuário", userAuth)
	form.SetUserId(p.ByName("id"))
	user := models.GetUser(form.Id)
	if user.Id == 0 {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
		return
	}
	form.SetUser(user)
	helpers.RenderTemplate(w, "layout_admin", "users/form", form)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	form := forms.NewUserFormData("Novo usuário", userAuth)

	form.SetUserId(p.ByName("id"))
	form.LoadFormData(r)

	user := form.GetUser()
	if validators.ValidateUser(form) && user.Save() {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
		return
	}

	helpers.RenderTemplate(w, "layout_admin", "users/form", form)
}

func ViewUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params, userAuth models.User) {
	form := forms.NewUserFormData("Editar usuário", userAuth)
	form.SetUserId(p.ByName("id"))
	user := models.GetUser(form.Id)
	if user.Id == 0 {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
		return
	}
	form.SetUser(user)
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

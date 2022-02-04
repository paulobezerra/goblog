package controllers

import (
	"net/http"

	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/controllers/validators"
	"github.com/paulobezerra/goblog/src/models"
)

func IndexUsers(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	users := models.FindAllUsers()
	data := map[string]interface{}{
		"User":  user,
		"users": users,
	}
	helpers.RenderTemplate(w, "layout_admin", "users/index", data)
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var (
		username           string            = ""
		firstname          string            = ""
		lastname           string            = ""
		errorMessage       *string           = nil
		validationMessages map[string]string = nil
		valid              bool              = false
	)
	if r.Method == "POST" {
		username = r.FormValue("username")
		firstname = r.FormValue("firstname")
		lastname = r.FormValue("lastname")
		password := r.FormValue("password")

		validationMessages, valid = validators.ValidateUser(username, firstname, lastname, password, false)
		if valid {
			user := models.CreateUser(username, firstname, lastname, password)
			if user.Id > 0 {
				http.Redirect(w, r, "/admin/users", http.StatusFound)
			}
		}
	}
	data := map[string]interface{}{
		"Username":           username,
		"Firstname":          firstname,
		"Lastname":           lastname,
		"ValidationMessages": validationMessages,
		"ErrorMessage":       &errorMessage,
		"User":               user,
		"FormTitle":          "Novo usuário",
	}
	helpers.RenderTemplate(w, "layout_admin", "users/form", data)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var (
		username           string            = ""
		firstname          string            = ""
		lastname           string            = ""
		errorMessage       *string           = nil
		validationMessages map[string]string = nil
		valid              bool              = false
	)
	id := r.URL.Query().Get("id")
	if r.Method == "GET" {
		user := models.GetUser(id)
		if user.Id == 0 {
			http.Redirect(w, r, "/admin/users", http.StatusFound)
		}
		username = user.Username
		firstname = user.Firstname
		lastname = user.Lastname
	} else if r.Method == "POST" {
		username = r.FormValue("username")
		firstname = r.FormValue("firstname")
		lastname = r.FormValue("lastname")
		password := r.FormValue("password")
		validationMessages, valid = validators.ValidateUser(username, firstname, lastname, password, true)
		if valid {
			user := models.UpdateUser(id, username, firstname, lastname, password)
			if user.Id > 0 {
				http.Redirect(w, r, "/admin/users", http.StatusFound)
			}
		}
	}
	data := map[string]interface{}{
		"Username":           username,
		"Firstname":          firstname,
		"Lastname":           lastname,
		"ValidationMessages": validationMessages,
		"ErrorMessage":       &errorMessage,
		"User":               user,
		"FormTitle":          "Novo usuário",
	}
	helpers.RenderTemplate(w, "layout_admin", "users/form", data)
}

func ViewUsers(w http.ResponseWriter, r *http.Request) {
	userAuth, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	id := r.URL.Query().Get("id")

	user := models.GetUser(id)
	if user.Id == 0 {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
	}
	username := user.Username
	firstname := user.Firstname
	lastname := user.Lastname

	data := map[string]interface{}{
		"Id":        id,
		"Username":  username,
		"Firstname": firstname,
		"Lastname":  lastname,
		"User":      userAuth,
	}
	helpers.RenderTemplate(w, "layout_admin", "users/view", data)
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	_, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	id := r.URL.Query().Get("id")

	user := models.GetUser(id)
	if user.Id == 0 || user.Username == "admin" {
		http.Redirect(w, r, "/admin/users", http.StatusFound)
		return
	}

	models.DeleteUser(id)
	http.Redirect(w, r, "/admin/users", http.StatusFound)
}

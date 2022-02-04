package controllers

import (
	"net/http"

	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/models"
	"github.com/paulobezerra/goblog/src/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err == nil && user.Id > 0 {
		http.Redirect(w, r, "/admin/posts", http.StatusFound)
		return
	}
	var message string
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user := models.FindOneUserByUsername(username)

		if utils.CheckPasswordHash(password, user.Password) {
			jwt := helpers.GenerateJWT(user)
			helpers.SaveCookie(helpers.CookieNameWithAuthToken, jwt, w)

			http.Redirect(w, r, "/admin/posts", http.StatusFound)
			return
		}
		message = "Login ou senha inválidos"
	}
	helpers.RenderTemplate(w, "layout", "login", message)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	helpers.DropCookie(helpers.CookieNameWithAuthToken, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

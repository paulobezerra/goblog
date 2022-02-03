package controllers

import (
	"net/http"

	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/models"
)

func IndexPosts(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	posts := models.FindAllPosts()
	data := map[string]interface{}{
		"User":  user,
		"posts": posts,
	}
	helpers.RenderTemplate(w, "layout_admin", "posts/index", data)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var (
		slug               string            = ""
		title              string            = ""
		abstract           string            = ""
		content            string            = ""
		tags               string            = ""
		errorMessage       *string           = nil
		validationMessages map[string]string = nil
		valid              bool              = false
	)
	if r.Method == "POST" {
		slug = r.FormValue("slug")
		title = r.FormValue("title")
		abstract = r.FormValue("abstract")
		content = r.FormValue("content")
		tags = r.FormValue("tags")

		validationMessages, valid = models.ValidatePost("0", slug, title, abstract, content, tags)
		if valid {
			errorMessage = models.CreatePost(slug, title, abstract, content, tags)
			if errorMessage == nil {
				http.Redirect(w, r, "/admin/posts", http.StatusFound)
			}
		}
	}
	data := map[string]interface{}{
		"Slug":               slug,
		"Title":              title,
		"Abstract":           abstract,
		"Content":            content,
		"Tags":               tags,
		"ValidationMessages": validationMessages,
		"ErrorMessage":       &errorMessage,
		"User":               user,
		"FormTitle":          "Novo post",
	}
	helpers.RenderTemplate(w, "layout_admin", "posts/form", data)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var (
		slug               string            = ""
		title              string            = ""
		abstract           string            = ""
		content            string            = ""
		tags               string            = ""
		errorMessage       *string           = nil
		validationMessages map[string]string = nil
		valid              bool              = false
	)
	id := r.URL.Query().Get("id")
	if r.Method == "GET" {
		post := models.GetPost(id)
		if user.Id == 0 {
			http.Redirect(w, r, "/admin/posts", http.StatusFound)
		}
		slug = post.Slug
		title = post.Title
		abstract = post.Abstract
		content = post.Content
		tags = post.Tags
	} else if r.Method == "POST" {
		slug = r.FormValue("slug")
		title = r.FormValue("title")
		abstract = r.FormValue("abstract")
		content = r.FormValue("content")
		tags = r.FormValue("tags")
		validationMessages, valid = models.ValidatePost(id, slug, title, abstract, content, tags)
		if valid {
			errorMessage = models.UpdatePost(id, slug, title, abstract, content, tags)
			if errorMessage == nil {
				http.Redirect(w, r, "/admin/posts", http.StatusFound)
			}
		}
	}
	data := map[string]interface{}{
		"Slug":               slug,
		"Title":              title,
		"Content":            content,
		"Abstract":           abstract,
		"Tags":               tags,
		"ValidationMessages": validationMessages,
		"ErrorMessage":       &errorMessage,
		"User":               user,
		"FormTitle":          "Editar post",
	}
	helpers.RenderTemplate(w, "layout_admin", "posts/form", data)
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	var (
		slug               string            = ""
		title              string            = ""
		content            string            = ""
		abstract           string            = ""
		tags               string            = ""
		errorMessage       *string           = nil
		validationMessages map[string]string = nil
	)
	id := r.URL.Query().Get("id")
	post := models.GetPost(id)
	if user.Id == 0 {
		http.Redirect(w, r, "/admin/posts", http.StatusFound)
	}
	slug = post.Slug
	title = post.Title
	abstract = post.Abstract
	content = post.Content
	tags = post.Tags

	data := map[string]interface{}{
		"Id":                 id,
		"Slug":               slug,
		"Title":              title,
		"Abstract":           abstract,
		"Content":            content,
		"Tags":               tags,
		"ValidationMessages": validationMessages,
		"ErrorMessage":       &errorMessage,
		"User":               user,
	}
	helpers.RenderTemplate(w, "layout_admin", "posts/view", data)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	_, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	id := r.URL.Query().Get("id")

	models.DeletePost(id)
	http.Redirect(w, r, "/admin/posts", http.StatusFound)
}

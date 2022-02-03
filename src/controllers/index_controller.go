package controllers

import (
	"net/http"

	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	posts := models.FindAllPosts()
	data := map[string]interface{}{
		"posts": posts,
	}
	helpers.RenderTemplate(w, "layout", "index", data)
}

func PublicPost(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")
	post := models.FindOnePostBySlug(slug)
	helpers.RenderTemplate(w, "layout", "post", post)
}

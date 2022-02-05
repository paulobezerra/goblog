package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/models"
)

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	posts := models.FindAllPosts()
	data := map[string]interface{}{
		"posts": posts,
	}
	helpers.RenderTemplate(w, "layout", "index", data)
}

func PublicPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	slug := p.ByName("slug")
	post := models.FindOnePostBySlug(slug)
	helpers.RenderTemplate(w, "layout", "post", post)
}

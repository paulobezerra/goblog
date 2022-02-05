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

func IndexPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	posts := models.FindAllPosts()
	data := map[string]interface{}{
		"PostsActive": "active",
		"User":        user,
		"Posts":       posts,
	}
	helpers.RenderTemplate(w, "layout_admin", "posts/index", data)
}

func FormCreatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	data := dto.NewPostDto("Novo post", user)
	helpers.RenderTemplate(w, "layout_admin", "posts/form", data)
}

func CreatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewPostDto("Novo post", user)
	form.LoadFormData(r)
	if validators.ValidatePost(&form) && form.Post.Create() {
		http.Redirect(w, r, "/admin/posts", http.StatusFound)
		return
	}
	helpers.RenderTemplate(w, "layout_admin", "posts/form", form)
}

func FormUpdatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewPostDto("Novo post", user)
	form.SetPostId(p.ByName("id"))
	post := models.GetPost(form.Id)
	if post.Id == 0 {
		http.Redirect(w, r, "/admin/posts", http.StatusFound)
		return
	}
	form.Post = *post
	helpers.RenderTemplate(w, "layout_admin", "posts/form", form)
}

func UpdatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewPostDto("Novo post", user)
	form.SetPostId(p.ByName("id"))
	form.LoadFormData(r)
	if validators.ValidatePost(&form) && form.Post.Save() {
		http.Redirect(w, r, "/admin/posts", http.StatusFound)
		return
	}
	helpers.RenderTemplate(w, "layout_admin", "posts/form", form)
}

func ViewPost(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewPostDto("Dados do post", user)
	form.SetPostId(p.ByName("id"))
	post := models.GetPost(form.Id)
	if post.Id == 0 {
		http.Redirect(w, r, "/admin/posts", http.StatusFound)
		return
	}
	form.Post = *post
	helpers.RenderTemplate(w, "layout_admin", "posts/view", form)
}

func DeletePost(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	id := p.ByName("id")
	idInt, _ := strconv.Atoi(id)
	post := models.GetPost(idInt)
	fmt.Println(id, idInt, post)
	if post.Id > 0 {
		post.Delete()
	}
	http.Redirect(w, r, "/admin/posts", http.StatusFound)
}

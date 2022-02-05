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

func IndexTags(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	tags := models.FindAllTags()
	data := map[string]interface{}{
		"TagsActive": "active",
		"User":       user,
		"Tags":       tags,
	}
	helpers.RenderTemplate(w, "layout_admin", "tags/index", data)
}

func FormCreateTag(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	data := dto.NewTagDto("Nova tag", user)
	helpers.RenderTemplate(w, "layout_admin", "tags/form", data)
}

func CreateTag(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewTagDto("Nova tag", user)
	form.LoadFormData(r)
	if validators.ValidateTag(&form) && form.Tag.Create() {
		http.Redirect(w, r, "/admin/tags", http.StatusFound)
		return
	}
	helpers.RenderTemplate(w, "layout_admin", "tags/form", form)
}

func FormUpdateTag(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewTagDto("Editar tag", user)
	form.SetTagId(p.ByName("id"))
	tag := models.GetTag(form.Id)
	if tag.Id == 0 {
		http.Redirect(w, r, "/admin/tags", http.StatusFound)
		return
	}
	form.Tag = *tag
	helpers.RenderTemplate(w, "layout_admin", "tags/form", form)
}

func UpdateTag(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewTagDto("Editar tag", user)
	form.SetTagId(p.ByName("id"))
	form.LoadFormData(r)
	if validators.ValidateTag(&form) && form.Tag.Save() {
		http.Redirect(w, r, "/admin/tags", http.StatusFound)
		return
	}
	helpers.RenderTemplate(w, "layout_admin", "tags/form", form)
}

func ViewTag(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	form := dto.NewTagDto("Dados da tag", user)
	form.SetTagId(p.ByName("id"))
	tag := models.GetTag(form.Id)
	if tag.Id == 0 {
		http.Redirect(w, r, "/admin/tags", http.StatusFound)
		return
	}
	form.Tag = *tag
	helpers.RenderTemplate(w, "layout_admin", "tags/view", form)
}

func DeleteTag(w http.ResponseWriter, r *http.Request, p httprouter.Params, user models.User) {
	id := p.ByName("id")
	idInt, _ := strconv.Atoi(id)
	tag := models.GetTag(idInt)
	fmt.Println(id, idInt, tag)
	if tag.Id > 0 {
		tag.Delete()
	}
	http.Redirect(w, r, "/admin/tags", http.StatusFound)
}

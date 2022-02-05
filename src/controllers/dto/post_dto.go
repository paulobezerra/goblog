package dto

import (
	"net/http"
	"strconv"

	"github.com/paulobezerra/goblog/src/models"
)

type PostDto struct {
	FormDto
	DashboardDto
	models.Post
	AllTags       []models.Tag
	AllCategories []models.Category
}

func NewPostDto(title string, user models.User) PostDto {
	return PostDto{
		FormDto: FormDto{
			FormTitle:          title,
			ValidationMessages: map[string]string{},
		},
		DashboardDto: DashboardDto{
			PostsActive: "active",
			User:        user,
		},
		Post: models.Post{
			AuthorId: user.Id,
		},
	}
}

func (form *PostDto) SetPostId(id string) {
	idInt, _ := strconv.Atoi(id)
	form.Id = idInt
}

func (form *PostDto) LoadFormData(r *http.Request) {
	if form.Id > 0 {
		post := models.GetPost(form.Id)
		form.Post.UpdatedAt = post.UpdatedAt
		form.Post.Model = post.Model
	}
	form.Slug = r.FormValue("slug")
	form.Title = r.FormValue("title")
	form.Abstract = r.FormValue("abstract")
	form.Content = r.FormValue("content")
	categoryId, _ := strconv.Atoi(r.FormValue("category"))
	form.Category = *models.GetCategory(categoryId)
	form.CategoryId = categoryId
	tagsIds := r.Form["tags"]
	if len(tagsIds) > 0 {
		var tagsIdsInt []int
		for _, id := range tagsIds {
			idInt, _ := strconv.Atoi(id)
			tagsIdsInt = append(tagsIdsInt, idInt)
		}
		form.Tags = models.FindTagsById(tagsIdsInt)
	}
}

func (post PostDto) ContainsTag(id int) bool {
	for _, tag := range post.Tags {
		if tag.Id == id {
			return true
		}
	}
	return false
}

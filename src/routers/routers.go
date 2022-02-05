package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/paulobezerra/goblog/src/controllers"
	"github.com/paulobezerra/goblog/src/controllers/helpers"
	"github.com/paulobezerra/goblog/src/models"
)

var router httprouter.Router

type Handle func(http.ResponseWriter, *http.Request, httprouter.Params, models.User)

func Init() *httprouter.Router {
	router = *httprouter.New()
	router.GET("/", controllers.Index)
	router.GET("/post/:slug", controllers.PublicPost)

	router.GET("/login", controllers.FormLogin)
	router.POST("/login", controllers.Login)
	router.GET("/logout", validateToken(controllers.Logout))

	router.GET("/admin/users", validateToken(controllers.IndexUsers))
	router.GET("/admin/users/new", validateToken(controllers.FormCreateUser))
	router.POST("/admin/users/new", validateToken(controllers.CreateUsers))
	router.GET("/admin/users/view/:id", validateToken(controllers.ViewUsers))
	router.GET("/admin/users/edit/:id", validateToken(controllers.FormUpdateUser))
	router.POST("/admin/users/edit/:id", validateToken(controllers.UpdateUsers))
	router.GET("/admin/users/delete/:id", validateToken(controllers.DeleteUsers))

	router.GET("/admin/categories", validateToken(controllers.IndexCategories))
	router.GET("/admin/categories/new", validateToken(controllers.FormCreateCategory))
	router.POST("/admin/categories/new", validateToken(controllers.CreateCategory))
	router.GET("/admin/categories/view/:id", validateToken(controllers.ViewCategory))
	router.GET("/admin/categories/edit/:id", validateToken(controllers.FormUpdateCategory))
	router.POST("/admin/categories/edit/:id", validateToken(controllers.UpdateCategory))
	router.GET("/admin/categories/delete/:id", validateToken(controllers.DeleteCategory))

	router.GET("/admin/tags", validateToken(controllers.IndexTags))
	router.GET("/admin/tags/new", validateToken(controllers.FormCreateTag))
	router.POST("/admin/tags/new", validateToken(controllers.CreateTag))
	router.GET("/admin/tags/view/:id", validateToken(controllers.ViewTag))
	router.GET("/admin/tags/edit/:id", validateToken(controllers.FormUpdateTag))
	router.POST("/admin/tags/edit/:id", validateToken(controllers.UpdateTag))
	router.GET("/admin/tags/delete/:id", validateToken(controllers.DeleteTag))

	router.GET("/admin/posts", validateToken(controllers.IndexPosts))
	router.GET("/admin/posts/new", validateToken(controllers.FormCreatePost))
	router.POST("/admin/posts/new", validateToken(controllers.CreatePost))
	router.GET("/admin/posts/view/:id", validateToken(controllers.ViewPost))
	router.GET("/admin/posts/edit/:id", validateToken(controllers.FormUpdatePost))
	router.POST("/admin/posts/edit/:id", validateToken(controllers.UpdatePost))
	router.GET("/admin/posts/delete/:id", validateToken(controllers.DeletePost))

	return &router
}

func validateToken(h Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, err := helpers.GetUserByCookieName(helpers.CookieNameWithAuthToken, r)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		h(w, r, ps, *user)
	}
}

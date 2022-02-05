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

	router.GET("/admin/posts", validateToken(controllers.IndexPosts))
	router.GET("/admin/posts/new", validateToken(controllers.FormCreatePost))
	router.POST("/admin/posts/new", validateToken(controllers.CreatePost))
	router.GET("/admin/posts/view/:id", validateToken(controllers.ViewPost))
	router.GET("/admin/posts/edit/:id", validateToken(controllers.FormUpdatePost))
	router.POST("/admin/posts/edit/:id", validateToken(controllers.UpdatePost))
	router.GET("/admin/posts/delete/:id", validateToken(controllers.DeletePost))

	return &router
}

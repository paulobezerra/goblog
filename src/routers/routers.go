package routers

import (
	"net/http"

	"github.com/paulobezerra/goblog/src/controllers"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/post", controllers.PublicPost)

	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/logout", controllers.Logout)

	mux.HandleFunc("/admin/users", controllers.IndexUsers)
	mux.HandleFunc("/admin/users/new", controllers.CreateUsers)
	mux.HandleFunc("/admin/users/view", controllers.ViewUsers)
	mux.HandleFunc("/admin/users/edit", controllers.UpdateUsers)
	mux.HandleFunc("/admin/users/delete", controllers.DeleteUsers)

	mux.HandleFunc("/admin/posts", controllers.IndexPosts)
	mux.HandleFunc("/admin/posts/new", controllers.CreatePost)
	mux.HandleFunc("/admin/posts/view", controllers.ViewPost)
	mux.HandleFunc("/admin/posts/edit", controllers.UpdatePost)
	mux.HandleFunc("/admin/posts/delete", controllers.DeletePost)

	return mux
}

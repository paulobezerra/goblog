package routers

import "net/http"

func InitStaticFilesServer(mux *http.ServeMux) *http.ServeMux {

	fileServer := http.FileServer(http.Dir("./public"))

	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	return mux
}

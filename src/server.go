package main

import (
	"log"
	"net/http"

	"github.com/paulobezerra/goblog/src/models"
	"github.com/paulobezerra/goblog/src/routers"
)

func main() {

	models.InitDB()

	mux := routers.Init()

	mux = routers.InitStaticFilesServer(mux)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

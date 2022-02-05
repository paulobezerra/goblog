package main

import (
	"log"
	"net/http"

	"github.com/paulobezerra/goblog/src/db"
	"github.com/paulobezerra/goblog/src/routers"
)

func main() {
	db.Migrate()
	db.Seed()

	router := routers.Init()
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

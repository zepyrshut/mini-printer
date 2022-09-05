package main

import (
	"github.com/go-chi/chi"
	"github.com/zepyrshut/mini-printer/handlers"
	"net/http"
)

const version = "1.0.0-FINAL.20220905"

func main() {
	router := chi.NewRouter()

	router.Get("/", handlers.Index)
	router.Post("/print", handlers.Print)

	println("Server is running on port 8082")
	println("A Mini Printer Server version " + version)
	err := http.ListenAndServe(":8082", router)
	if err != nil {
		panic(err.Error())
	}
}

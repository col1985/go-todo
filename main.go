package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/col1985/go-todo/db"
	"github.com/col1985/go-todo/router"
)


func main() {
	db.Init()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
	})

	r.Mount("/todos", router.TodoRoutes())

	err := http.ListenAndServe(":8080", r)

	if err != nil {
			panic(err)
	}

	log.Print("Go Todo App is running on http://localhost:8080")
}
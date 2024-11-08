package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/col1985/go-todo/db"
	"github.com/col1985/go-todo/router"
)


func main() {
	db.Init()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
	})

	r.Get("/metrics", promhttp.Handler().ServeHTTP)

	r.Mount("/todos", router.TodoRoutes())

	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	err := http.ListenAndServe(port, r)

	if err != nil {
			panic(err)
	}

	fmt.Printf("Go Todo App is running on http://localhost%s", port)
}
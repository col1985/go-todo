package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/col1985/go-todo/db"
	"github.com/col1985/go-todo/router"
	"github.com/col1985/go-todo/utils"
)

func main() {
	isDev := os.Getenv("IS_DEV")
	if isDev == "true" {
		utils.LoadEnvFile()
	}

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

	host := ""
	if isDev == "true" {
		host = fmt.Sprintf("localhost:%s", os.Getenv("API_PORT"))
	} else {
		host = fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	}

	log.Printf("Starting Todo API is running on %s ...", host)

	err := http.ListenAndServe(host, r)

	if err != nil {
		panic(err)
	}
}
package router

import (
	"encoding/json"
	"net/http"

	"github.com/col1985/go-todo/db"
	"github.com/go-chi/chi/v5"
)

func TodoRoutes() chi.Router {
    r := chi.NewRouter()

    r.Get("/", GetTodoList)
    r.Post("/", CreateTodo)
    r.Get("/{id}", GetTodo)
    r.Put("/{id}", UpdateTodo)
    r.Delete("/{id}", DeleteTodo)
    return r
}

func errorHandler(w http.ResponseWriter, msg string, statusCode int) {
	http.Error(w, msg, statusCode)
  return
}

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	res, err := db.GetTodoList()
	encodingErr := json.NewEncoder(w).Encode(res)

	if err != nil || encodingErr != nil {
		errorHandler(w, "Internal error", http.StatusInternalServerError)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := db.GetTodoById(id)

	encodingErr := json.NewEncoder(w).Encode(res)

	if err != nil || encodingErr != nil {
		errorHandler(w, "Internal error", http.StatusInternalServerError)
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo *db.Todo
	encodingErr := json.NewDecoder(r.Body).Decode(&todo)

	if encodingErr != nil {
		errorHandler(w, encodingErr.Error(), http.StatusBadRequest)
	}

	res, err := db.CreateTodo(todo)

	if err != nil {
		errorHandler(w, "Internal error", http.StatusInternalServerError)
	}

	resErr := json.NewEncoder(w).Encode(res)
	if resErr != nil {
		errorHandler(w, resErr.Error(), http.StatusBadRequest)
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var todo db.Todo
	decodeErr := json.NewDecoder(r.Body).Decode(&todo)

	if decodeErr != nil {
		errorHandler(w, decodeErr.Error(), http.StatusBadRequest)
	}

	dbTodo, err := db.GetTodoById(id)

	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
	}

	dbTodo.Title = todo.Title
	dbTodo.Author = todo.Author
	dbTodo.Completed = todo.Completed

	res, err := db.UpdateTodo(dbTodo)

	resErr := json.NewEncoder(w).Encode(res)
	if resErr != nil {
		errorHandler(w, resErr.Error(), http.StatusBadRequest)
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := db.DeleteTodo(id)

	if err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
	}
}

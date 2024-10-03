package router

import (
	"encoding/json"
	"net/http"

	"github.com/col1985/go-todo/db"
	"github.com/col1985/go-todo/utils"
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

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	res, err := db.GetTodoList()
	encodingErr := json.NewEncoder(w).Encode(res)

	if err != nil || encodingErr != nil {
		utils.HttpErrorHandler(w, "Internal error", http.StatusInternalServerError)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := db.GetTodoById(id)
	encodingErr := json.NewEncoder(w).Encode(res)

	if err != nil || encodingErr != nil {
		utils.HttpErrorHandler(w, "Internal error", http.StatusInternalServerError)
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo *db.Todo
	encodingErr := json.NewDecoder(r.Body).Decode(&todo)

	if encodingErr != nil {
		utils.HttpErrorHandler(w, encodingErr.Error(), http.StatusBadRequest)
	}

	res, err := db.CreateTodo(todo)

	if err != nil {
		utils.HttpErrorHandler(w, "Internal error", http.StatusInternalServerError)
	}

	resErr := json.NewEncoder(w).Encode(res)
	if resErr != nil {
		utils.HttpErrorHandler(w, resErr.Error(), http.StatusBadRequest)
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var todo db.Todo
	decodeErr := json.NewDecoder(r.Body).Decode(&todo)

	if decodeErr != nil {
		utils.HttpErrorHandler(w, decodeErr.Error(), http.StatusBadRequest)
	}

	dbTodo, err := db.GetTodoById(id)

	if err != nil {
		utils.HttpErrorHandler(w, err.Error(), http.StatusInternalServerError)
	}

	dbTodo.Task = todo.Task
	dbTodo.Author = todo.Author
	dbTodo.Completed = todo.Completed
	dbTodo.UpdateDate = utils.GetDateString()

	res, err := db.UpdateTodo(dbTodo)

	resErr := json.NewEncoder(w).Encode(res)
	if resErr != nil {
		utils.HttpErrorHandler(w, resErr.Error(), http.StatusBadRequest)
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := db.DeleteTodo(id)

	if err != nil {
		utils.HttpErrorHandler(w, err.Error(), http.StatusBadRequest)
	}
}

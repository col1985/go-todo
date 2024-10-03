package db

import (
	"fmt"

	"github.com/col1985/go-todo/utils"
	"github.com/google/uuid"
)

func CreateTodo(todo *Todo) (*Todo, error) {
	todo.ID = uuid.New().String()
	todo.Completed = false
	todo.CreatedDate = utils.GetDateString()

	res := db.Create(&todo)

	if res.Error != nil {
		return nil, res.Error
	}

	return todo, nil
}

func GetTodoById(id string) (*Todo, error) {
	var todo Todo

	res := db.First(&todo, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, utils.ErrorHandler(fmt.Sprintf("Todo item with `id` %s is not found", id))
	}

	return &todo, nil
}

func GetTodoList() ([]*Todo, error) {
	var todos []*Todo

	res := db.Find(&todos)
	if res.Error != nil {
			return nil, utils.ErrorHandler("No todo items found")
	}

	return todos, nil
}

func UpdateTodo(todo *Todo) (*Todo, error) {
	var todoToUpdate Todo

	result := db.Model(&todoToUpdate).Where("id = ?", todo.ID).Updates(todo)
	if result.RowsAffected == 0 {
			return &todoToUpdate, utils.ErrorHandler(fmt.Sprintf("Todo item with `id` %s is not updated.", todo.ID))
	}

	return todo, nil
}

func DeleteTodo(id string) error {
	var deleted Todo

	result := db.Where("id = ?", id).Delete(&deleted)
	if result.RowsAffected == 0 {
			return utils.ErrorHandler(fmt.Sprintf("Todo item with `id` %s is not deleted.", id))
	}

	return nil
}

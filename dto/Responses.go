package dto

import "github.com/sobhan/tod/entitys"

type TodoResponseBody struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
	Time string `json:"time"`
}

func ParseFromEntity(todo entitys.TodoLists) TodoResponseBody {
	return TodoResponseBody{
		ID:   todo.ID,
		Name: todo.Name,
		Done: todo.Done,
		Time: todo.Time.String(),
	}
}

func ParseFromEntityList(todos []entitys.TodoLists) []TodoResponseBody {
	var TodoResponseBodys []TodoResponseBody
	for _, todo := range todos {
		TodoResponseBodys = append(TodoResponseBodys, ParseFromEntity(todo))
	}
	return TodoResponseBodys
}

package services

import (
	"github.com/sobhan/tod/Repositories"
	"github.com/sobhan/tod/dto"
	"github.com/sobhan/tod/entitys"
)

type UserService interface {
	AddTask(dto.TodoRequestBody) (entitys.TodoLists, error)
	ListTasks() ([]entitys.TodoLists, error)
	UpdateTask(int, dto.TodoRequestBody) error
	RemoveTask(int) error
	GetTask(int) (entitys.TodoLists, error)
}

type userService struct {
	TodoListsRepository Repositories.TodoListsRepository
}

func NewProductService(TodoListsRepository Repositories.TodoListsRepository) userService {
	return userService{
		TodoListsRepository: TodoListsRepository,
	}
}
func (us userService) AddTask(dto dto.TodoRequestBody) (entitys.TodoLists, error) {
	todoList := entitys.TodoLists{
		Name: dto.Name,
		Time: dto.Time,
	}
	return us.TodoListsRepository.Save(todoList), nil
}

func (us userService) ListTasks() ([]entitys.TodoLists, error) {
	return us.TodoListsRepository.FindAll(), nil
}

func (us userService) UpdateTask(productID int, dto dto.TodoRequestBody) error {
	todoList := entitys.TodoLists{
		Name: dto.Name,
		Time: dto.Time,
		Done: dto.Done,
	}

	us.TodoListsRepository.Update(productID, todoList)
	return nil
}

func (us userService) RemoveTask(productID int) error {
	todoList, err := us.GetTask(productID)
	if err != nil {
		return err
	}
	us.TodoListsRepository.Delete(todoList)
	return nil
}

func (us userService) GetTask(productID int) (entitys.TodoLists, error) {
	return us.TodoListsRepository.FindByID(productID), nil
}

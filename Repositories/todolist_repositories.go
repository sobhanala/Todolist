package Repositories

import (
	"github.com/sobhan/tod/db"
	"github.com/sobhan/tod/entitys"
)

type ProductRepositoryInterface interface {
	FindAll() []entitys.TodoLists
	FindByID(int) entitys.TodoLists
	Save(entitys.TodoLists) entitys.TodoLists
	Delete(int, entitys.TodoLists)
}

type TodoListsRepository struct {
	DB db.Database
}

func NewProductRepostiory(DB db.Database) TodoListsRepository {
	return TodoListsRepository{
		DB: DB,
	}
}

func (p *TodoListsRepository) FindAll() []entitys.TodoLists {
	var products []entitys.TodoLists
	p.DB.DB.Find(&products)

	return products
}

func (p *TodoListsRepository) FindByID(id int) entitys.TodoLists {
	var product entitys.TodoLists
	p.DB.DB.First(&product, id)

	return product
}

func (p *TodoListsRepository) Save(product entitys.TodoLists) entitys.TodoLists {
	p.DB.DB.Save(&product)

	return product
}

func (p *TodoListsRepository) Delete(product entitys.TodoLists) {
	p.DB.DB.Delete(&product)
}
func (p *TodoListsRepository) Update(id int, product entitys.TodoLists) error {
	result := p.DB.DB.Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/sobhan/tod/Repositories"
	"github.com/sobhan/tod/dto"
	_ "github.com/sobhan/tod/entitys"
	"github.com/sobhan/tod/services"
)

type TodolistHandlers interface {
	AddTask(ctx *gin.Context)
	GetTask(ctx *gin.Context)
	ListTasks(ctx *gin.Context)
	RemoveTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
}

type todoListHandler struct {
	ProductService services.UserService
}

func NewHttpHandler(
	productService services.UserService,
) todoListHandler {
	return todoListHandler{
		ProductService: productService,
	}
}

func (services todoListHandler) ListTasks(ctx *gin.Context) {
	products, err := services.ProductService.ListTasks()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productsResponse := dto.ParseFromEntityList(products)

	ctx.JSON(http.StatusOK, productsResponse)

}

func (services todoListHandler) GetTask(ctx *gin.Context) {
	rawProduct_id := ctx.Param("id")
	product_id, err := strconv.Atoi(rawProduct_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	product, err := services.ProductService.GetTask(product_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return

	}

	productResponse := dto.ParseFromEntity(product)
	ctx.JSON(http.StatusOK, productResponse)

}

func (services todoListHandler) RemoveTask(ctx *gin.Context) {
	rawProduct_id := ctx.Param("id")
	product_id, err := strconv.Atoi(rawProduct_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	err = services.ProductService.RemoveTask(product_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())

	}
	ctx.JSON(http.StatusOK, "deleted")

}

func (services todoListHandler) UpdateTask(ctx *gin.Context) {
	rawProductID := ctx.Param("id")
	productID, err := strconv.Atoi(rawProductID)
	print(err, productID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	//todo
	print("salam")
	dtoRequstbody := new(dto.TodoRequestBody)
	if err := ctx.BindJSON(&dtoRequstbody); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = services.ProductService.UpdateTask(productID, *dtoRequstbody)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
	}
	ctx.JSON(http.StatusOK, "updated")
}

func (services todoListHandler) AddTask(ctx *gin.Context) {
	dtoRequstbody := new(dto.TodoRequestBody)
	if err := ctx.BindJSON(&dtoRequstbody); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	Todolist, err := services.ProductService.AddTask(*dtoRequstbody)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	TodolistResponse := dto.ParseFromEntity(Todolist)
	ctx.JSON(http.StatusOK, TodolistResponse)
}

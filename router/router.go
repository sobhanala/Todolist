package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sobhan/tod/Repositories"
	"github.com/sobhan/tod/db"
	handlers "github.com/sobhan/tod/handler"
	"github.com/sobhan/tod/services"
)

func SetupRouter() error {

	database, err := db.NewDatabase()
	if err != nil {
		return err
	}

	todoRepo := Repositories.NewProductRepostiory(*database)
	todoService := services.NewProductService(todoRepo)
	todoHandler := handlers.NewHttpHandler(todoService)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	v1 := r.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.POST("/", todoHandler.AddTask)
			todos.GET("/", todoHandler.ListTasks)
			todos.GET("/:id", todoHandler.GetTask)
			todos.PUT("/:id", todoHandler.UpdateTask)
			todos.DELETE("/:id", todoHandler.RemoveTask)
		}
	}

	r.GET("/", func(c *gin.Context) {
		tasks, err := todoService.ListTasks()
		if err != nil {
			c.String(http.StatusInternalServerError, "Unable to fetch tasks")
			return
		}
		c.HTML(http.StatusOK, "index.html", tasks)
	})
	r.Run(":8080")
	return nil

}

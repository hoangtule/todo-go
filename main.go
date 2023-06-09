package main

import (
	//"new-example/repo/sqlite"
	"new-example/repo/postgresql"
	"new-example/service"
	"new-example/transport/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//
	todoRepo := &postgresql.TodoRepositoryImpl{}
	todoRepo.InitDB()

	//
	service := service.TodoService{
		TodoRepo: todoRepo,
	}

	//usecases
	httpService := &http.TodoService{
		BusinessService: &service,
	}

	router := gin.Default()

	router.GET("/todos", httpService.GetAll)

	router.GET("/todos/:id", httpService.GetByID)

	router.POST("/todos", httpService.Create)

	router.PUT("/todos/:id", httpService.Update)

	router.DELETE("/todos/:id", httpService.Delete)

	router.Run(":8080")
}

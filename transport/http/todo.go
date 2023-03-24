package http

import (
	"net/http"
	"new-example/entity"

	"github.com/gin-gonic/gin"
)

type TodoBusinessService interface {
	InitDB()
	GetByID(id string) (*entity.Todo, error)
	GetAll() ([]*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(id string, todo *entity.Todo) (*entity.Todo, error)
	Delete(id string) error
}

type TodoService struct {
	BusinessService TodoBusinessService
}

func (service *TodoService) GetAll(c *gin.Context) {
	todos, err := service.BusinessService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (service *TodoService) GetByID(c *gin.Context) {
	id := c.Param("id")

	todo, err := service.BusinessService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (service *TodoService) Create(c *gin.Context) {
	var todo entity.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo, err := service.BusinessService.Create(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTodo)
}

func (service *TodoService) Update(c *gin.Context) {
	id := c.Param("id")

	var todo entity.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := service.BusinessService.Update(id, &todo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

func (service *TodoService) Delete(c *gin.Context) {
	id := c.Param("id")

	err := service.BusinessService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

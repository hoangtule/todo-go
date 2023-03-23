package main

import (
	"net/http"

	"new-example/entity"
	"new-example/repo"
	"new-example/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repo.NewTodoRepository()
	service := service.NewTodoService(repo)

	router := gin.Default()

	router.GET("/todos", func(c *gin.Context) {
		todos, err := service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	})

	router.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		todo, err := service.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, todo)
	})

	router.POST("/todos", func(c *gin.Context) {
		var todo entity.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newTodo, err := service.Create(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newTodo)
	})

	router.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		var todo entity.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedTodo, err := service.Update(id, &todo)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedTodo)
	})

	router.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := service.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
	})

	router.Run(":8080")
}

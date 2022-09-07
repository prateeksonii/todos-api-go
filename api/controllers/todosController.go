package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id          int       `json:"id"`
	Todo        string    `json:"todo" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	IsCompleted bool      `json:"is_completed"`
}

var todos = []todo{
	{Id: 1, Todo: "Default todo", CreatedAt: time.Now(), IsCompleted: false},
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func PostTodo(c *gin.Context) {
	var json todo

	err := c.ShouldBindJSON(&json)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	newTodo := todo{
		Id:          len(todos) + 1,
		Todo:        json.Todo,
		CreatedAt:   time.Now(),
		IsCompleted: false,
	}

	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, gin.H{
		"ok":   true,
		"todo": newTodo,
	})

}

func ToggleTodoStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos[i].IsCompleted = !todo.IsCompleted
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

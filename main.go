package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prateeksonii/todos/api-go/api/controllers"
)

func main() {
	app := gin.Default()

	app.GET("/todos", controllers.GetTodos)
	app.POST("/todos", controllers.PostTodo)
	app.PATCH("/todos/:id/toggle", controllers.ToggleTodoStatus)

	app.Run(":4000")
}

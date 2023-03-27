package routes

import (
	"github.com/Ryu-seunghyun/Todo-List/controllers"
	"github.com/gin-gonic/gin"
)

/*
GET /todos				GetTodos()
GET /todos/:item_id		GetTodoById()
POST /todos				CreateTodo()
PATCH /todos/:item_id	UpdateTodoById()
DELETE /todos/:item_id 	DeleteTodoById()
*/

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/todos", controllers.GetTodos)
	r.GET("/todos/:item_id", controllers.GetTodoById)
	r.POST("/todos", controllers.CreateTodo)
	r.PATCH("/todos/:item_id", controllers.UpdateTodoById)
	r.DELETE("/todos/:item_id", controllers.DeleteTodoById)

	return r
}

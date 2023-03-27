package controllers

import (
	"net/http"

	"github.com/Ryu-seunghyun/Todo-List/model"
	"github.com/gin-gonic/gin"
)

/*
GET /todos				GetTodos()
GET /todos/:item_id		GetTodoById()
POST /todos				CreateTodo()
PATCH /todos/:item_id	UpdateTodoById()
DELETE /todos/:item_id 	DeleteTodoById()
*/

// GET /todos  -> 전체 인덱스 조회
func GetTodos(c *gin.Context) {
	todos, err := model.ManyGet()
	if err != nil { // 조회 실패  500
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{ // 조회 성공  200
		"item": todos,
	})
}

// GET /todos/:item_id	->	ID로 특정 인덱스 조회
func GetTodoById(c *gin.Context) { // 없는 ID 조회  404 ( gin이 알아서 처리가 안됌!! )
	id := c.Param("item_id")
	todo, code, err := model.OnlyGet(id)
	if err != nil { // 조회 실패  500
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, gin.H{"item": todo}) // 조회 성공  200
}

// POST /todos	-> 생성
func CreateTodo(c *gin.Context) {
	var t model.Todo
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	todo, code, err := model.OnlyCreate(&t)
	if err != nil { // 런타임 에러  500
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, gin.H{ // 생성 성공  201
		"message": "Success Created!!! ",
		"item":    todo,
	})
}

// PATCH /todos/:item_id  -> ID로 특정 인덱스 수정
func UpdateTodoById(c *gin.Context) { // 없는 ID 수정  404
	var t model.Todo
	id := c.Param("item_id")
	if err := c.ShouldBindJSON(&t); err != nil { // 입력 에러  400
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	todo, code, err := model.OnlyUpdate(id, &t)
	if err != nil { // 런타임 에러  500
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, gin.H{ // 수정 성공  200
		"message": "Success Updated!!!",
		"item":    todo,
	})
}

// DELETE /todos/:item_id  -> ID로 특정 인덱스 삭제
func DeleteTodoById(c *gin.Context) { // 없는 ID 삭제  404
	id := c.Param("item_id")
	code, err := model.OnlyDelete(id)
	if err != nil { // 삭제 실패 (이미 삭제된 인덱스이거나 런타임 에러)  400 or 500
		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(code, gin.H{}) // 삭제 성공  204
}

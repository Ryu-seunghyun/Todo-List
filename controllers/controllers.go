package controllers

import (
	"net/http"

	"github.com/Ryu-seunghyun/Todo-List/model/domain"
	"github.com/gin-gonic/gin"
)

var err error

func (h *Handler) ListTodos(ctx *gin.Context) {
	var todos []domain.Todo
	todos, err = h.Services.Todos.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (h *Handler) GetTodo(ctx *gin.Context) {
	var todo domain.Todo
	id := ctx.Param("item_id")
	todo, err = h.Services.Todos.GetByTodoId(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "해당하는 Todo를 찾을 수 없습니다.")
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (h *Handler) CreateTodo(ctx *gin.Context) {
	var t domain.Todo
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, "잘못된 요청입니다.")
		return
	}
	if t.Description == "" {
		ctx.JSON(http.StatusBadRequest, "Description 값을 입력하지 않았습니다.")
		return
	}
	var todo domain.Todo
	todo, err = h.Services.Todos.Create(ctx, t)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, todo)
}

func (h *Handler) UpdateTodo(ctx *gin.Context) {
	var t, ori domain.Todo
	id := ctx.Param("item_id")
	ori, err = h.Services.Todos.GetByTodoId(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "해당하는 Todo를 찾을 수 없습니다.")
		return
	}
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if t.Description == "" && t.Status == "" {
		ctx.JSON(http.StatusBadRequest, "변경할 값을 입력하세요.")
		return
	}
	if t.Description == ori.Description && t.Status == ori.Status {
		ctx.JSON(http.StatusBadRequest, "변경사항이 없습니다.")
		return
	}

	var todo domain.Todo
	todo, err = h.Services.Todos.Update(ctx, id, t)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func (h *Handler) DeleteTodo(ctx *gin.Context) {
	var t domain.Todo
	id := ctx.Param("item_id")
	t, err = h.Services.Todos.GetDeletedByTodoId(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "해당하는 Todo를 찾을 수 없습니다.")
		return
	}
	if t.DeletedAt.Valid {
		ctx.JSON(http.StatusBadRequest, "이미 삭제된 Todo입니다.")
		return
	}
	if err := h.Services.Todos.Delete(ctx, t); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

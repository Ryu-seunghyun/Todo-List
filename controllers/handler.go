package controllers

import (
	"github.com/Ryu-seunghyun/Todo-List/service"
	"github.com/gin-gonic/gin"
)

// Handler 에 서비스를 담아서 사용

type Handler struct {
	Services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/todos", h.ListTodos)
	router.GET("/todos/:item_id", h.GetTodo)
	router.POST("/todos", h.CreateTodo)
	router.PATCH("/todos/:item_id", h.UpdateTodo)
	router.DELETE("/todos/:item_id", h.DeleteTodo)

	return router
}

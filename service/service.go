package service

import (
	"context"
	"fmt"

	"github.com/Ryu-seunghyun/Todo-List/model"
	"github.com/Ryu-seunghyun/Todo-List/model/dto"
	"github.com/Ryu-seunghyun/Todo-List/repository"
)

type Services struct {
	Todos Todos
}

type Todos interface {
	List(ctx context.Context) ([]dto.Todo, error)
	GetByTodoId(ctx context.Context, todoId string) (dto.Todo, error)
	GetDeletedByTodoId(ctx context.Context, todoId string) (dto.Todo, error)
	Create(ctx context.Context, t dto.Todo) (dto.Todo, error)
	Update(ctx context.Context, id string, t dto.Todo) (dto.Todo, error)
	Delete(ctx context.Context, t dto.Todo) error
}

type TodoService struct {
	todoRepository repository.Todos
}

func NewServices(repository repository.Repositories) *Services {
	todos := NewTodoService(repository.Todos)
	return &Services{
		Todos: todos,
	}
}
func NewTodoService(todoRepository repository.Todos) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

// DTO -> domain  -> DTO

func (s *TodoService) List(ctx context.Context) ([]dto.Todo, error) {
	var todos []dto.Todo
	fromTodos, err := s.todoRepository.FindAll()
	fmt.Println(fromTodos)
	for _, fromTodo := range fromTodos {
		var todo dto.Todo
		todo = dto.NewTodo(fromTodo)
		todos = append(todos, todo)
	}
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetByTodoId(ctx context.Context, todoId string) (dto.Todo, error) {
	var todo dto.Todo
	fromTodo, err := s.todoRepository.FindByTodoId(todoId)
	todo = dto.NewTodo(fromTodo)
	if err != nil {
		return dto.Todo{}, err
	}
	return todo, nil
}
func (s *TodoService) GetDeletedByTodoId(ctx context.Context, todoId string) (dto.Todo, error) {
	var todo dto.Todo
	fromTodo, err := s.todoRepository.FindDeletedByTodoId(todoId)
	todo = dto.NewTodo(fromTodo)
	if err != nil {
		return dto.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Create(ctx context.Context, toTodo dto.Todo) (dto.Todo, error) {
	var todo dto.Todo
	if toTodo.Status == "" {
		toTodo.Status = model.TodoStatusReady
	}
	fromTodo := dto.NewDomain(toTodo)
	fromTodo, err := s.todoRepository.Create(fromTodo)
	todo = dto.NewTodo(fromTodo)
	if err != nil {
		return dto.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Update(ctx context.Context, id string, toTodo dto.Todo) (dto.Todo, error) {
	var todo dto.Todo
	fromTodo := dto.NewDomain(todo)
	fromTodo, err := s.todoRepository.Update(id, fromTodo)
	todo = dto.NewTodo(fromTodo)
	if err != nil {
		return dto.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Delete(ctx context.Context, toTodo dto.Todo) error {
	fromTodo := dto.NewDomain(toTodo)
	if err := s.todoRepository.Delete(fromTodo); err != nil {
		return err
	}
	return nil
}

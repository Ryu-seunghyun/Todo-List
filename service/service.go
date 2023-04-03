package service

import (
	"context"

	"github.com/Ryu-seunghyun/Todo-List/model"
	"github.com/Ryu-seunghyun/Todo-List/model/domain"
	"github.com/Ryu-seunghyun/Todo-List/repository"
)

type Services struct {
	Todos Todos
}

type Todos interface {
	List(ctx context.Context) ([]domain.Todo, error)
	GetByTodoId(ctx context.Context, todoId string) (domain.Todo, error)
	GetDeletedByTodoId(ctx context.Context, todoId string) (domain.Todo, error)
	Create(ctx context.Context, t domain.Todo) (domain.Todo, error)
	Update(ctx context.Context, id string, t domain.Todo) (domain.Todo, error)
	Delete(ctx context.Context, t domain.Todo) error
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

func (s *TodoService) List(ctx context.Context) ([]domain.Todo, error) {
	todos, err := s.todoRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) GetByTodoId(ctx context.Context, todoId string) (domain.Todo, error) {
	todo, err := s.todoRepository.FindByTodoId(todoId)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}
func (s *TodoService) GetDeletedByTodoId(ctx context.Context, todoId string) (domain.Todo, error) {
	todo, err := s.todoRepository.FindDeletedByTodoId(todoId)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Create(ctx context.Context, t domain.Todo) (domain.Todo, error) {
	if t.Status == "" {
		t.Status = model.TodoStatusReady
	}
	todo, err := s.todoRepository.Create(t)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Update(ctx context.Context, id string, t domain.Todo) (domain.Todo, error) {
	todo, err := s.todoRepository.Update(id, t)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Delete(ctx context.Context, t domain.Todo) error {
	if err := s.todoRepository.Delete(t); err != nil {
		return err
	}
	return nil
}

package repository

// repository 의 역할 : DAO 의 접근/저장 함수 구현, 최소 기능 수준의 함수
//  서비스에서 쿼리를 고려할 필요가 없도록

import (
	"errors"

	"github.com/Ryu-seunghyun/Todo-List/model/domain"
	"gorm.io/gorm"
)

type Repositories struct {
	Todos Todos
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Todos: NewTodoRepository(db),
	}
}

type Todos interface {
	FindAll() ([]domain.Todo, error)
	FindByTodoId(todoId string) (domain.Todo, error)
	FindDeletedByTodoId(todoId string) (domain.Todo, error)
	// FindByUserId
	Create(domain.Todo) (domain.Todo, error)
	Update(string, domain.Todo) (domain.Todo, error)
	Delete(domain.Todo) error
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (t *TodoRepository) FindAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	if err := t.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodoRepository) FindByTodoId(todoId string) (domain.Todo, error) {
	var todo domain.Todo
	if err := t.db.Find(&todo, "ID = ?", todoId).Error; err != nil {
		return domain.Todo{}, err
	}
	// Not found resource
	if todo.ID == "" {
		return domain.Todo{}, errors.New("not found")
	}
	return todo, nil
}

func (t *TodoRepository) FindDeletedByTodoId(todoId string) (domain.Todo, error) {
	var todo domain.Todo
	if err := t.db.Unscoped().Find(&todo, "ID = ?", todoId).Error; err != nil {
		return domain.Todo{}, err
	}
	// Not found resource
	if todo.ID == "" {
		return domain.Todo{}, errors.New("not found")
	}
	return todo, nil
}

func (t *TodoRepository) Create(todo domain.Todo) (domain.Todo, error) {
	if err := t.db.Save(&todo).Error; err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

// 조회 후 삭제
func (t *TodoRepository) Update(id string, todo domain.Todo) (domain.Todo, error) {
	var ori_todo domain.Todo
	if err := t.db.Model(&ori_todo).Where("ID = ?").UpdateColumns(&todo).Error; err != nil {
		return domain.Todo{}, err
	}
	return ori_todo, nil
}

func (t *TodoRepository) Delete(todo domain.Todo) error {

	if err := t.db.Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}

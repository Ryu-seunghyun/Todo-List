package dto

// DTO의 역할 : DAO와 연계하여 데이터 처리에 활용할 구조 정의 및 생성자 구현
import (
	"time"

	"github.com/Ryu-seunghyun/Todo-List/model/domain"
	"github.com/google/uuid"
)

// go-playground/validator
type Todo struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"last_updated_at"`
	DeletedAt bool      `json:"-"`
}

func NewDomain(t Todo) domain.Todo {
	if t.Id == "" {
		t.Id = uuid.New().String()
	}
	return domain.Todo{
		ID:          t.Id,
		Description: t.Description,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func NewTodo(todo domain.Todo) Todo {
	return Todo{
		Id:          todo.ID,
		Description: todo.Description,
		Status:      todo.Status,

		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
		DeletedAt: todo.DeletedAt.Valid,
	}
}

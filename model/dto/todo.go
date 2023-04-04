package dto

// DTO의 역할 : DAO와 연계하여 데이터 처리에 활용할 구조 정의 및 생성자 구현
import (
	"time"

	"github.com/Ryu-seunghyun/Todo-List/model/domain"
)

type Todo struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"last_updated_at"`
}

func NewTodoFromDomain(todo domain.Todo) Todo {
	return Todo{
		Id:          todo.ID,
		Description: todo.Description,
		Status:      todo.Status,

		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

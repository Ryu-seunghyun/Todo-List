package domain

// Domain 의 역할  -  DAO 정의 및 생성자 구현

import (
	"github.com/Ryu-seunghyun/Todo-List/model"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	// ID string `gorm:"type:varchar(255)"`
	/* gorm.Model 의 ID (Auto_Increment) 는 제어가 어려움 (POST 에러가 발생해도 자동증가)
	1. ID를 직접 관리하는 방법   (호출되는 모든곳에서 관리해야 함)
	2. uuid 로 관리하는 방법	(랜덤으로 생성되는 uuid를 식별하는 방법) - User 의 Profile

	일단, ID로 기능 구현 먼저 할 것
	*/
	// UserId	string `gorm:"type:varchar(255)"`  이후 사용자별 Todo 서비스 확장 시
	Description string `gorm:"type:varchar(255)"`
	Status      string `gorm:"type:varchar(255)"`

	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt
}

// 생성자는 서비스단에서 로직처리에 활용하기 위한 용도
func newTodo(desc string) Todo { // 올바른 형식으로 요청을 했는지 검증 필요
	return Todo{
		Description: desc,
		Status:      model.TodoStatusReady,
	}
}

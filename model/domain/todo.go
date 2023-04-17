package domain

// Domain 의 역할  -  DAO 정의 및 생성자 구현

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Status      string `gorm:"type:varchar(255)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

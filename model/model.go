package model

import (
	"errors"
	"net/http"
	"time"

	"github.com/Ryu-seunghyun/Todo-List/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Todo struct {
	// gorm.Model
	ID          uint       `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Description string     `json:"description" gorm:"unique"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"last_updated_at"`
	Done        bool       `json:"done"`
	DeletedAt   *time.Time `json:"-" sql:"index"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

/*
GET /todos				GetTodos()
GET /todos/:item_id		GetTodoById()
POST /todos				CreateTodo()
PATCH /todos/:item_id	UpdateTodoById()
DELETE /todos/:item_id 	DeleteTodoById()
*/

// ID 로 받는 것  -  Get, Delete, Update
// *todo 로 (입력) 받을 수 있는 것 - Create, Update

// Get all 은 db.Find(&todos)  								[]Todos
// GetById 는 db.First(&todo,"ID=?",id) 					*Todo
// Create 는 db.Create(&data) 								*Todo		description 을 필수로 입력받도록
// Update 는 db.Model(&todo).Where("ID=?",id).Update(&data)     		description or done 중 하나 이상 입력받도록
// Delete 는 db.Delete(&todo,"ID=?",id) -> Unscoped 					이미 삭제된 인덱스 처리 후 삭제 처리

func OnlyCreate(t *Todo) (*Todo, int, error) { // Call by Reference -> 추가 작업이 가능
	if t.Description == "" {
		return nil, http.StatusBadRequest, errors.New("required (description) key ")
	}
	db.NewRecord(t) // ID pk 는 unique 하여 별도 에러처리는 하지 않음.
	if err := db.Create(&t).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return t, http.StatusCreated, nil
}

func ManyGet() ([]Todo, error) { // Slice ptr
	var todos []Todo
	resp := db.Find(&todos) // 데이터가 없어도 조회 처리
	if resp.Error != nil {
		return nil, resp.Error
	}
	return todos, nil
}

func OnlyGet(id string) (*Todo, int, error) {
	var todo Todo
	resp := db.First(&todo, "ID=?", id) // ID 존재 여부
	switch resp.Error {
	case gorm.ErrRecordNotFound:
		return nil, http.StatusNotFound, resp.Error
	case nil:
		return &todo, 200, nil
	default:
		return nil, http.StatusInternalServerError, resp.Error
	}
	// if resp.Error != nil {
	// 	return nil, resp.Error
	// }
}

func OnlyUpdate(id string, t *Todo) (*Todo, int, error) {
	var todo Todo
	if err := db.First(&todo, "ID=?", id); err != nil {
		return nil, 404, err.Error
	}
	if (t.Description == "" && t.Done == todo.Done) || (t.Description == todo.Description && t.Done == todo.Done) {
		return nil, 400, errors.New("required ( 'description'|'done' )  Or  No Changed")
	}
	resp := db.Model(&todo).Update(&t) // PATCH 로 변경된 부분만 부분수정 가능
	switch resp.Error {
	case gorm.ErrRecordNotFound:
		return nil, http.StatusNotFound, resp.Error
	case nil:
		return &todo, 200, nil
	default:
		return nil, http.StatusInternalServerError, resp.Error
	}
}

func OnlyDelete(id string) (int, error) { // 204 - NoContent
	var todo Todo
	resp := db.Unscoped().First(&todo, "ID=? AND deleted_at IS NOT NULL", id)
	switch {
	case int(todo.ID) == 0: // 없는 ID 삭제 시
		return http.StatusNotFound, gorm.ErrRecordNotFound
	case resp.RowsAffected != 0: // 이미 삭제된 ID
		return http.StatusBadRequest, errors.New("index already deleted!! ")
	}

	resp = db.Delete(&todo, "ID=?", id)
	if resp.Error != nil {
		return http.StatusInternalServerError, resp.Error
	}
	return http.StatusNoContent, nil
}

package database

import (
	"fmt"
	"sync"

	"github.com/Ryu-seunghyun/Todo-List/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var database *gorm.DB

// 구조체로 만들면 직관적임!!
type Database struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

// Connect 생성과 호출의 구분을 명확히함.
func GetConnection(config Database) *gorm.DB {
	once.Do(func() {
		database = newConnection(config)
	})
	return database
}

func newConnection(config Database) *gorm.DB {
	dsn := getDSN(config)
	db, err := gorm.Open(
		mysql.Open(dsn),
	)
	if err != nil {
		panic("Database 연결에 실패하였습니다.")
	}
	return db
}

func getDSN(config Database) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
	)
}

func AutoMigrate() {
	database.AutoMigrate(&domain.Todo{})
}

// DB 구성의 setter 설정 (+ 디폴트)

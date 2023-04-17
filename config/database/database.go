package database

import (
	"fmt"
	"sync"

	"github.com/Ryu-seunghyun/Todo-List/model/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var database *gorm.DB

const (
	filePath = "./config/database/"
	fileName = "app"
	fileType = "json"
)

// 구조체로 만들면 직관적임!!
type Database struct {
	Host         string `mapstructure:"DB_HOST"`
	Port         string `mapstructure:"DB_PORT"`
	User         string `mapstructure:"DB_USER"`
	Password     string `mapstructure:"DB_PASSWORD"`
	DatabaseName string `mapstructure:"DB_NAME"`
}

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
	fileConfig, err := LoadConfig()
	if err != nil {
		panic("Database 구성을 불러올 수 없습니다.")
	}
	config = fileConfig
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
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

func LoadConfig() (config Database, err error) {
	viper.AddConfigPath(filePath)
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)

	// viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}

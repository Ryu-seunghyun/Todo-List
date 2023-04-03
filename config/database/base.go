package database

// DB를 호출하는 곳에서 사용
func SetDefault(db *Database) {
	db.Host = defaultHost
	db.Port = defaultPort
	db.User = defaultUser
	db.Password = defaultPassword
	db.DatabaseName = defaultDatabaseName
}

const (
	defaultHost         = "localhost"
	defaultPort         = 3306
	defaultUser         = "root"
	defaultPassword     = "mysql123"
	defaultDatabaseName = "todo"
)

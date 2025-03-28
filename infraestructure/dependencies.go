package infraestructure

var mysql *MySQL

func InitDependencies() {
	mysql = NewMySQL()
}

func GetMySQL() *MySQL {
	return mysql
}
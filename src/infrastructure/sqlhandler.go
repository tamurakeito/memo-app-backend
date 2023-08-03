package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

// sqlHandlerのインスタンスを生成
func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root:password@tcp(memo-app-db:3306)/memo_app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

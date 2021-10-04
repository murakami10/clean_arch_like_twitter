package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler{
	conn, err := sql.Open("mysql", "user:user@tcp(127.0.0.1:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	SqlHandler := new(SqlHandler)
	SqlHandler.Conn = conn

	err = conn.Ping()
	if err != nil {
		fmt.Println(err)
		fmt.Println("error")
	}else{
		fmt.Println("ok")
	}

	return SqlHandler
}
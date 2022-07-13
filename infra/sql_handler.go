package infra

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SqlHandler struct {
	Conn *sqlx.DB
}

func NewSqlHandler() (*SqlHandler, error) {

	conn, err := sqlx.Open("mysql", "user:user@tcp(127.0.0.1:3306)/test_database")
	if err != nil {
		return nil, err
	}

	return &SqlHandler{Conn: conn}, nil
}

package infra

import (
	"clean_arch/interfaces/database"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "user:user@tcp(127.0.0.1:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	SqlHandler := new(SqlHandler)
	SqlHandler.Conn = conn

	return SqlHandler
}

func (h *SqlHandler) Execute(s string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := h.Execute(s, args)
	if err != nil {
		return res, err
	}
	res.Result = result

	return res, nil
}

func (h *SqlHandler) Query(s string, args ...interface{}) (database.Row, error) {
	rows, err := h.Conn.Query(s, args)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}

package common

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	NamedExec(string, interface{}) (sql.Result, error)
}

type myDB struct {
	sqlxDb sqlx.DB
}

func (db *myDB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return db.sqlxDb.NamedExec(query, arg)
}

func NewDB() DB {
	return &myDB{}
}

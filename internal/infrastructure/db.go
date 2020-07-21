package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DbConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
}

type DB interface {
	NamedExec(string, interface{}) (sql.Result, error)
	Select(interface{}, string) error
}

type myDB struct {
	sqlxDb *sqlx.DB
}

func (db *myDB) Select(dest interface{}, query string) error {
	return db.sqlxDb.Select(dest, query)
}

func (db *myDB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return db.sqlxDb.NamedExec(query, arg)
}

func NewDB(sqlxDb *sqlx.DB) DB {
	return &myDB{
		sqlxDb: sqlxDb,
	}
}

func CreateDB(dbConfig DbConfig) (DB, error) {
	dataSourceString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Host,
		dbConfig.Port,
	)
	db, err := sqlx.Connect("postgres", dataSourceString)
	if err != nil {
		return nil, err
	}

	if pingError := db.Ping(); pingError != nil {
		log.Printf("Ping failed, %s", pingError)
	}

	return NewDB(db), nil
}

package infrastructure

import (
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

func CreateDB(dbConfig DbConfig) (*sqlx.DB, error) {
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

	return db, nil
}

package db

import (
	"database/sql"
	"log"
	"time"

	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type mDB struct {
	db *sql.DB
}

type DBConfig struct {
	User     string
	Passwd   string
	DBName   string
	DBHost   string
	DBPort   string
	Attempts int
}

//InitDb faz a conexão com o banco
func InitDb(dbConfig DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Passwd, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	var dbErr error
	for attempt := 1; attempt < dbConfig.Attempts; attempt++ {
		dbErr = db.Ping()
		if dbErr == nil {
			break
		}
		log.Printf("[WARN][VITRINE] %s \n", dbErr)
		time.Sleep(time.Duration(attempt) * time.Second)
	}
	return db, dbErr
}

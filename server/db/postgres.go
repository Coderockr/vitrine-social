package db

import (
	"database/sql"
	"log"
	"time"

	"fmt"
	"os"

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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Passwd, os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), dbConfig.DBName)
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

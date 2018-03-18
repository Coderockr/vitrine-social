package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"fmt"

	"github.com/jmoiron/sqlx"

	// needed for sqlx to know how to connect
	_ "github.com/lib/pq"
)

type mDB struct {
	db *sql.DB
}

// Config informs the data needed for connecting to the database
type Config struct {
	User     string
	Passwd   string
	DBName   string
	DBHost   string
	DBPort   string
	Attempts int
}

//InitDb faz a conexão com o banco
func InitDb(config Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Passwd, config.DBHost, config.DBPort, config.DBName)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	var dbErr error
	for attempt := 1; attempt < config.Attempts; attempt++ {
		dbErr = db.Ping()
		if dbErr == nil {
			break
		}
		log.Printf("[WARN][VITRINE] %s \n", dbErr)
		time.Sleep(time.Duration(attempt) * time.Second)
	}
	return db, dbErr
}

// GetFromEnv will return the database connection based on the enviroment variables
func GetFromEnv() (*sqlx.DB, error) {
	dbConf := Config{
		User:     os.Getenv("DATABASE_USER"),
		Passwd:   os.Getenv("DATABASE_PASSWORD"),
		DBName:   os.Getenv("DATABASE_NAME"),
		DBHost:   os.Getenv("DATABASE_HOST"),
		DBPort:   os.Getenv("DATABASE_PORT"),
		Attempts: 10,
	}
	return InitDb(dbConf)
}

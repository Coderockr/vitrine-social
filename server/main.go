package main

import (
	"log"
	"net/http"

	"os"

	"github.com/Coderockr/vitrine-social/server/db"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	env := os.Getenv("VITRINESOCIAL_ENV")
	err := godotenv.Load("config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	StartServer()
}

//StartServer rotas e handlers
func StartServer() {
	dbConf := db.DBConfig{
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
	}
	dbConn, err := db.InitDb(dbConf)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
	}

	r := mux.NewRouter()

	http.Handle("/", r)
	http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(http.DefaultServeMux))
}

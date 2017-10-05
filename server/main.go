package main

import (
	"Coderockr/vitrine-social/server/handlers"
	"log"
	"net/http"
	"time"

	"os"

	"github.com/Coderockr/vitrine-social/server/auth"
	"github.com/Coderockr/vitrine-social/server/db"
	"github.com/Coderockr/vitrine-social/server/db/inmemory"
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
		10,
	}
	_, err := db.InitDb(dbConf)
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	muxR := mux.NewRouter()
	options := auth.Options{
		SigningMethod: "RS256",
		PrivateKey:    os.Getenv("VITRINESOCIAL_PRIVATE_KEY"), // $ openssl genrsa -out app.rsa keysize
		PublicKey:     os.Getenv("VITRINESOCIAL_PUBLIC_KEY"),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:    60 * time.Minute,
	}

	// creates the route with Bolt and JWT options
	authRoute := auth.NewAuthRoute(inmemory.NewUserRepository(), options)
	v1 := muxR.PathPrefix("/v1").Subrouter()
	authSub := v1.PathPrefix("/auth").Subrouter()
	authSub.HandleFunc("/login", authRoute.Login)
	v1.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {

	})

	v1.Handle("/need/{id}", handlers.NeedGet())

	err = http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(muxR))
	if err != nil {
		log.Fatal(err)
	}
}

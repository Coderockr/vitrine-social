package main

import (
	"log"
	"net/http"
	"time"

	"os"

	"github.com/Coderockr/vitrine-social/server/auth"
	"github.com/Coderockr/vitrine-social/server/db"
	"github.com/Coderockr/vitrine-social/server/db/inmemory"
	"github.com/codegangsta/negroni"
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
		20,
	}
	_, err := db.InitDb(dbConf)
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	mux := mux.NewRouter()
	options := auth.Options{
		SigningMethod: "RS256",
		PrivateKey:    os.Getenv("VITRINESOCIAL_PRIVATE_KEY"), // $ openssl genrsa -out app.rsa keysize
		PublicKey:     os.Getenv("VITRINESOCIAL_PUBLIC_KEY"),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:    60 * time.Minute,
	}

	// creates the route with Bolt and JWT options
	authRoute := auth.NewAuthRoute(inmemory.NewUserRepository(), options)
	app := negroni.Classic()
	v1 := mux.PathPrefix("/v1").Subrouter()
	authRoute := v1.PathPrefix("/auth").Subrouter()
	authRoute.HandleFunc("/login", authRoute.Login)
	// authRoute.HandleFunc("/signin", authRoute.Signin)
	v1.HandleFunc("/search", negroni.New(
		authRoute.AuthMiddleware,
		func(w http.ResponseWriter, req *http.Request) {

		}))

	http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(mux))
}

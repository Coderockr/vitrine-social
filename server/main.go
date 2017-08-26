package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"os"

	"github.com/Coderockr/vitrine-social/server/auth"
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
		10,
	}
	dbx, err := db.InitDb(dbConf)
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	mux := mux.NewRouter()
	privatekey, err := ioutil.ReadFile(os.Getenv("VITRINESOCIAL_PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Error getting the jwt private key (%s): %v\n", os.Getenv("VITRINESOCIAL_PRIVATE_KEY"), err)
	}
	publickey, err := ioutil.ReadFile(os.Getenv("VITRINESOCIAL_PUBLIC_KEY"))
	if err != nil {
		log.Fatalf("Error getting the jwt public key (%s): %v\n", os.Getenv("VITRINESOCIAL_PUBLIC_KEY"), err)
	}
	println(string(privatekey))
	options := auth.Options{
		SigningMethod: "RS256",
		PrivateKey:    privatekey,
		PublicKey:     publickey,
		Expiration:    60 * time.Minute,
	}

	// creates the route with Bolt and JWT options
	authRoute := auth.NewAuthRoute(db.NewUserRepository(dbx), options)
	v1 := mux.PathPrefix("/v1").Subrouter()
	authSub := v1.PathPrefix("/auth").Subrouter()
	authSub.HandleFunc("/login", authRoute.Login)
	v1.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {

	})
	err = http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(mux))
	if err != nil {
		log.Fatal(err)
	}
}

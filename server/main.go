package main

import (
	"log"
	"net/http"
	"time"

	"os"

	"github.com/Coderockr/vitrine-social/server/auth"
	"github.com/Coderockr/vitrine-social/server/db"
	"github.com/Coderockr/vitrine-social/server/db/inmemory"
	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("VITRINESOCIAL_ENV")
	err := godotenv.Load("server/config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading file ", "server/config/"+env+".env")
	}
	StartServer()
}

//StartServer rotas e handlers
func StartServer() {
	dbConf := db.DBConfig{
		User:     os.Getenv("DATABASE_USER"),
		Passwd:   os.Getenv("DATABASE_PASSWORD"),
		DBName:   os.Getenv("DATABASE_NAME"),
		DBHost:   os.Getenv("DATABASE_HOST"),
		DBPort:   os.Getenv("DATABASE_PORT"),
		Attempts: 10,
	}
	conn, err := db.InitDb(dbConf)
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	oR := repo.NewOrganizationRepository(conn)
	nR := repo.NewNeedRepository(conn)

	mux := mux.NewRouter()
	options := auth.Options{
		SigningMethod: "RS256",
		PrivateKey:    os.Getenv("VITRINESOCIAL_PRIVATE_KEY"), // $ openssl genrsa -out app.rsa keysize
		PublicKey:     os.Getenv("VITRINESOCIAL_PUBLIC_KEY"),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:    60 * time.Minute,
	}

	// creates the route with Bolt and JWT options
	authRoute := auth.NewAuthRoute(inmemory.NewUserRepository(), options)

	v1 := mux.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/auth/login", authRoute.Login)

	v1.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {})

	organizationRoute := handlers.NewOrganizationHandler(oR)
	v1.HandleFunc("/organization/{id:[0-9]+}", organizationRoute.Get)

	needRoute := handlers.NewNeedHandler(nR, oR)
	v1.Handle("/need/{id}", needRoute.NeedGet())

	err = http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(mux))
	if err != nil {
		log.Fatal(err)
	}
}

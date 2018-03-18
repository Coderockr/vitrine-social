// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Coderockr/vitrine-social/server/db"
	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/middlewares"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Iniciar o servidor HTTP",
	Run:   withEnvironment(serveCmdFunc),
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func getJWTOptions() handlers.JWTOptions {
	return handlers.JWTOptions{
		SigningMethod: "HS256",
		PrivateKey:    []byte(os.Getenv("VITRINESOCIAL_PRIVATE_KEY")), // $ openssl genrsa -out app.rsa keysize
		PublicKey:     []byte(os.Getenv("VITRINESOCIAL_PUBLIC_KEY")),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:    60 * time.Minute,
	}
}

func serveCmdFunc(cmd *cobra.Command, args []string) {

	conn, err := db.GetFromEnv()
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	oR := repo.NewOrganizationRepository(conn)
	nR := repo.NewNeedRepository(conn)

	mux := mux.NewRouter()

	v1 := mux.PathPrefix("/v1").Subrouter()
	options := getJWTOptions()

	AuthHandler := handlers.AuthHandler{
		UserGetter: &repo.UserRepository{
			DB: conn,
		},
		TokenManager: &handlers.JWTManager{OP: options},
	}
	v1.HandleFunc("/auth/login", AuthHandler.Login)

	v1.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {})

	organizationRoute := handlers.NewOrganizationHandler(oR)
	v1.HandleFunc("/organization/{id:[0-9]+}", organizationRoute.Get)

	needRoute := handlers.NewNeedHandler(nR, oR)
	v1.Handle("/need/{id}", needRoute.NeedGet())

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(middlewares.Cors))

	// router goes last
	n.UseHandler(mux)

	log.Printf("Listening at :%s", os.Getenv("API_PORT"))
	err = http.ListenAndServe(":"+os.Getenv("API_PORT"), context.ClearHandler(n))
	if err != nil {
		log.Fatal(err)
	}
}

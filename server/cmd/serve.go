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

func serveCmdFunc(cmd *cobra.Command, args []string) {

	conn, err := db.GetFromEnv()
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	oR := repo.NewOrganizationRepository(conn)
	nR := repo.NewNeedRepository(conn)

	needResponseRepo := repo.NewNeedResponseRepository(conn)

	mux := mux.NewRouter()

	v1 := mux.PathPrefix("/v1").Subrouter()
	options := getJWTOptions()

	AuthHandler := handlers.AuthHandler{
		UserGetter:   oR,
		TokenManager: &handlers.JWTManager{OP: options},
	}

	authMiddleware := negroni.New()
	authMiddleware.UseFunc(AuthHandler.AuthMiddleware)

	v1.HandleFunc("/auth/login", AuthHandler.Login)

	v1.Path("/auth/update-password").Handler(authMiddleware.With(
		negroni.WrapFunc(handlers.UpdatePasswordHandler(oR)),
	)).Methods("POST")

	v1.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {})

	v1.HandleFunc("/organization/{id:[0-9]+}", handlers.GetOrganizationHandler(oR.Get)).Methods("GET")

	v1.Path("/organization/{id:[0-9]+}").Handler(authMiddleware.With(
		negroni.WrapFunc(handlers.UpdateOrganizationHandler(oR)),
	)).Methods("PUT")

	v1.HandleFunc("/need/{id}", handlers.GetNeedHandler(nR, oR)).Methods("GET")

	v1.Path("/need").Handler(authMiddleware.With(
		negroni.WrapFunc(handlers.CreateNeedHandler(nR.Create)),
	)).Methods("POST")

	v1.HandleFunc("/need/{id}/response", handlers.NeedResponse(nR, needResponseRepo)).
		Methods("POST")

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

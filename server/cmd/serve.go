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
	"os"
	"strings"

	"github.com/Coderockr/vitrine-social/server/server"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Iniciar o servidor HTTP",
	Run:   serveCmdFunc,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP("env", "e", os.Getenv("VITRINESOCIAL_ENV"), "Informe qual ambiente deve ser iniciado (dev ou production)")
}

func serveCmdFunc(cmd *cobra.Command, args []string) {
	env := strings.ToLower(cmd.Flag("env").Value.String())

	err := godotenv.Load("./config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading file ", "./config/"+env+".env")
	}

	server.StartServer()
}

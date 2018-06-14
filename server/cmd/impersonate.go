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
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/spf13/cobra"
)

// impersonateCmd represents the impersonate command
var impersonateCmd = &cobra.Command{
	Use:   "impersonate [organizationId]",
	Short: "Returs a JWT for the organization informed",
	Run: withEnvironment(func(cmd *cobra.Command, args []string) {
		options := getJWTOptions()

		manager := handlers.JWTManager{OP: options}

		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		token, err := manager.CreateToken(
			model.User{
				ID: id,
			},
			nil,
		)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Print(token)
	}),
}

func init() {
	rootCmd.AddCommand(impersonateCmd)
}

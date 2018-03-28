// Copyright © 2018 NAME HERE <EMAIL createRESS>
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
	"time"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/model"

	"github.com/Coderockr/vitrine-social/server/db"
	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new organization (ORG) and returns a URL for confirmation",
	Run:   withEnvironment(createCmdFunc),
}

var (
	email   string
	name    string
	logo    string
	address string
	phone   string
	resume  string
	video   string
	slug    string
)

func init() {
	orgCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&email, "email", "e", "", "organization's e-mail")
	createCmd.Flags().StringVarP(&name, "name", "n", "", "organization's name")
	createCmd.Flags().StringVarP(&logo, "logo", "l", "", "organization's logo")
	createCmd.Flags().StringVarP(&address, "address", "a", "", "organization's address")
	createCmd.Flags().StringVarP(&phone, "phone", "p", "", "organization's phone")
	createCmd.Flags().StringVarP(&slug, "slug", "s", "", "organization's slug")
	createCmd.Flags().StringVarP(&resume, "resume", "r", "", "organization's resume")
	createCmd.Flags().StringVarP(&video, "video", "v", "", "organization's video")

	createCmd.MarkFlagRequired("email")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("logo")
	createCmd.MarkFlagRequired("address")
	createCmd.MarkFlagRequired("phone")
	createCmd.MarkFlagRequired("slug")
}

func createCmdFunc(cmd *cobra.Command, args []string) {
	conn, err := db.GetFromEnv()
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	oR := repo.NewOrganizationRepository(conn)

	o, err := oR.Create(model.Organization{
		User: model.User{
			Email:    email,
			Password: "",
		},
		Name:    name,
		Logo:    logo,
		Address: address,
		Phone:   phone,
		Slug:    slug,
		Resume:  resume,
		Video:   video,
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	options := getJWTOptions()
	options.Expiration = 24 * 3 * time.Hour

	manager := handlers.JWTManager{OP: options}

	token, err := manager.CreateToken(o.User)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("%s/complete-register/%s", os.Getenv("FRONTEND_URL"), token)
}

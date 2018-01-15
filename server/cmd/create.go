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
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new organization (ONG) and returns a URL for confirmation",
	Run:   withEnvironment(createCmdFunc),
}

func init() {
	ongCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("email", "e", "", "organization's e-mail")
	createCmd.Flags().StringP("name", "n", "", "organization's name")
	createCmd.Flags().StringP("logo", "l", "", "organization's logo")
	createCmd.Flags().StringP("address", "a", "", "organization's address")
	createCmd.Flags().StringP("phone", "p", "", "organization's phone")
	createCmd.Flags().StringP("resume", "r", "", "organization's resume")
	createCmd.Flags().StringP("video", "v", "", "organization's video")
	createCmd.Flags().StringP("slug", "s", "", "organization's slug")
}

func createCmdFunc(cmd *cobra.Command, args []string) {

}

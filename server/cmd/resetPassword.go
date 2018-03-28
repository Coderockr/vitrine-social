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
	"strconv"
	"strings"

	"github.com/Coderockr/vitrine-social/server/db"
	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/spf13/cobra"
)

// resetPasswordCmd represents the resetPassword command
var resetPasswordCmd = &cobra.Command{
	Aliases: []string{"resetPassword"},
	Use:     "reset-password organizationEmailOrId [newPassword]",
	Short:   "Resets the password for a organization",
	Args:    cobra.RangeArgs(1, 2),
	Run:     withEnvironment(resetPasswordCmdFunc),
}

func init() {
	orgCmd.AddCommand(resetPasswordCmd)
}

func resetPasswordCmdFunc(cmd *cobra.Command, args []string) {
	var o *model.Organization

	conn, err := db.GetFromEnv()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	r := repo.NewOrganizationRepository(conn)

	if id, err := strconv.ParseInt(args[0], 10, 64); err == nil {
		o, err = r.Get(id)
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}
	}

	if o == nil {
		o, err = r.GetByEmail(args[0])
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}
	}

	var pass string
	if len(args) == 2 {
		pass = strings.TrimSpace(args[1])
	}

	if len(pass) == 0 {
		pass = randSeq(10)
		println(pass)
	}

	r.ResetPasswordTo(o, pass)
}

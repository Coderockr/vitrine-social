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
	"os"

	"github.com/spf13/cobra"
)

var env = os.Getenv("VITRINESOCIAL_ENV")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vitrine-social",
	Short: "Uma vitrine para projetos sociais divulgarem suas necessidades",
	Long:  `Uma vitrine para projetos sociais divulgarem suas necessidades de doação e possibilidade de volutariado`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vitrine-social.yaml)")
	rootCmd.Flags().StringVarP(&env, "env", "e", os.Getenv("VITRINESOCIAL_ENV"), "Informe qual ambiente deve ser iniciado (dev ou production)")
}

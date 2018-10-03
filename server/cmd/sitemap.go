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

	"github.com/Coderockr/vitrine-social/server/db"
	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/spf13/cobra"
	"github.com/thecodenation/go-sitemap-generator/stm"
)

// sitemapCmd represents the sitemap command
var sitemapCmd = &cobra.Command{
	Aliases: []string{"sitemap"},
	Use:     "sitemap-generate",
	Short:   "Generate sitemap",
	Run:     withEnvironment(sitemapFunc),
}

func init() {
	rootCmd.AddCommand(sitemapCmd)
}

func sitemapFunc(cmd *cobra.Command, args []string) {
	var o []*model.Organization
	sm := generateSitemap()

	conn, err := db.GetFromEnv()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	rO := repo.NewOrganizationRepository(conn)
	rN := repo.NewNeedRepository(conn)
	o, err = rO.GetAll()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	for _, v := range o {
		sm.Add(stm.URL{"loc": fmt.Sprintf("entidade/%d", v.ID), "changefreq": "hourly", "priority": 0.8})
		n, err := rN.GetOrganizationNeeds(v.ID, "id", "asc")
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}
		for _, k := range n {
			sm.Add(stm.URL{"loc": fmt.Sprintf("%s/v1/need/%d/share", os.Getenv("API_URL"), k.ID), "changefreq": "hourly", "priority": 0.8})
		}
	}
	err = saveSitemap(sm)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func generateSitemap() *stm.Sitemap {
	sm := stm.NewSitemap()
	sm.SetDefaultHost("https://www.vitrinesocial.org/")
	sm.Add(stm.URL{"loc": "", "changefreq": "hourly", "priority": 1.0})
	sm.Add(stm.URL{"loc": "sobre", "changefreq": "hourly"})
	sm.Add(stm.URL{"loc": "contato", "changefreq": "hourly"})

	return sm
}

func saveSitemap(sm *stm.Sitemap) error {
	f, err := os.Create("../frontend/public/sitemap.xml")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(sm.XMLContent())
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}

	log.Println("Generated sitemap")
	return nil
}

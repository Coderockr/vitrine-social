package storage

import (
	"os"

	"github.com/graymeta/stow"
	// support Google storage
	google "github.com/graymeta/stow/google"
	// support local storage
	local "github.com/graymeta/stow/local"
	// support s3 storage
	s3 "github.com/graymeta/stow/s3"
)

// Dial dials stow storage.
// See stow.Dial for more information.
func Dial() (stow.Location, error) {
	var kind string
	var config stow.ConfigMap

	if os.Getenv("STORAGE") == "local" {
		kind = "local"
		config = stow.ConfigMap{
			local.ConfigKeyPath: os.Getenv("STORAGE_PATH"),
		}
	}

	if os.Getenv("STORAGE") == "s3" {
		kind = "s3"
		config = stow.ConfigMap{
			s3.ConfigAccessKeyID: "246810",
			s3.ConfigSecretKey:   "abc123",
			s3.ConfigRegion:      "eu-west-1",
		}
	}

	if os.Getenv("STORAGE") == "google" {
		kind = "google"
		config = stow.ConfigMap{
			google.ConfigJSON:      "json",
			google.ConfigProjectId: "project_id",
			google.ConfigScopes:    "scopes",
		}
	}

	location, err := stow.Dial(kind, config)
	if err != nil {
		return nil, err
	}
	defer location.Close()

	return location, err
}

//Container cria o diretório ou bucket para salvar o arquivo
func Container(location stow.Location, id string) (stow.Container, error) {
	container, err := location.Container(os.Getenv("STORAGE_PATH") + "/" + id)
	if err != nil && err.Error() == "not found" {
		container, err = location.CreateContainer(id)
		if err != nil {
			return nil, err
		}
	}
	return container, nil
}

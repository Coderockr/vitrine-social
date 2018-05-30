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
			local.ConfigKeyPath: os.Getenv("STORAGE_LOCAL_PATH"),
		}
	}

	if os.Getenv("STORAGE") == "s3" {
		kind = "s3"
		config = stow.ConfigMap{
			s3.ConfigAccessKeyID: os.Getenv("STORAGE_S3_CONFIG_ACCESS_KEY_ID"),
			s3.ConfigSecretKey:   os.Getenv("STORAGE_S3_CONFIG_SECRET_KEY"),
			s3.ConfigRegion:      os.Getenv("STORAGE_S3_CONFIG_REGION"),
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
	var basePath string

	if os.Getenv("STORAGE") == "local" {
		basePath = os.Getenv("STORAGE_LOCAL_PATH")

	}
	if os.Getenv("STORAGE") == "s3" {
		basePath = os.Getenv("STORAGE_S3_PATH")

	}
	if os.Getenv("STORAGE") == "google" {
		basePath = os.Getenv("STORAGE_GOOGLE_PATH")

	}

	container, err := location.Container(basePath + "/" + id)
	if err != nil && err.Error() == "not found" {
		container, err = location.CreateContainer(basePath + "/" + id)
		if err != nil {
			return nil, err
		}
	}
	return container, nil
}

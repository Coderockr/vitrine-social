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

// Connect to storage and return the container
func Connect() (stow.Container, error) {
	var kind string
	var config stow.ConfigMap
	var containerName string

	if os.Getenv("STORAGE") == "local" {
		kind = "local"
		config = stow.ConfigMap{
			local.ConfigKeyPath: os.Getenv("STORAGE_LOCAL_PATH"),
		}
		containerName = os.Getenv("STORAGE_LOCAL_PATH")
	}

	if os.Getenv("STORAGE") == "s3" {
		kind = "s3"
		config = stow.ConfigMap{
			s3.ConfigAccessKeyID: os.Getenv("STORAGE_S3_CONFIG_ACCESS_KEY_ID"),
			s3.ConfigSecretKey:   os.Getenv("STORAGE_S3_CONFIG_SECRET_KEY"),
			s3.ConfigRegion:      os.Getenv("STORAGE_S3_CONFIG_REGION"),
		}
		containerName = os.Getenv("STORAGE_S3_BUCKET_NAME")
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

	container, err := Container(location, containerName)
	if err != nil {
		return nil, err
	}

	return container, err
}

//Container search or create directory/bucket to store files
func Container(location stow.Location, containerName string) (stow.Container, error) {
	container, err := location.Container(containerName)
	if err != nil && err.Error() == "not found" {
		container, err = location.CreateContainer(containerName)
		if err != nil {
			return nil, err
		}
	}
	return container, err
}

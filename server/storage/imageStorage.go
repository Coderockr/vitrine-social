package storage

import (
	"bytes"
	"fmt"
	"log"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gobuffalo/uuid"
	"github.com/graymeta/stow"
)

type needImageRepo interface {
	Get(int64) (*model.Need, error)
	CreateImage(model.NeedImage) (model.NeedImage, error)
	GetNeedsImages(model.Need) ([]model.NeedImage, error)
	DeleteImage(imageID int64, needID int64) error
}

type orgImageRepo interface {
	Get(int64) (*model.Organization, error)
	CreateImage(model.OrganizationImage) (model.OrganizationImage, error)
	DeleteImage(imageID int64, organizationID int64) error
}

// ImageStorage will save files into the storage and reference then into the database
type ImageStorage struct {
	Container              stow.Container
	NeedRepository         needImageRepo
	OrganizationRepository orgImageRepo
}

// DeleteNeedImage removes the image from a need, if user has permission
func (s *ImageStorage) DeleteNeedImage(t *model.Token, needID, imageID int64) error {

	n, err := s.NeedRepository.Get(needID)
	if err != nil {
		return fmt.Errorf("there is no need with the id %d", needID)
	}

	if t.UserID != n.OrganizationID {
		return fmt.Errorf("need %d does not belong to organization %d", needID, t.UserID)
	}

	images, err := s.NeedRepository.GetNeedsImages(*n)
	if err != nil {
		return err
	}

	for _, i := range images {
		if i.ID != imageID {
			continue
		}

		s.Container.RemoveItem(i.URL)
		s.NeedRepository.DeleteImage(i.ID, i.NeedID)
		return nil
	}

	return fmt.Errorf("there is no image with id %d at the need %d", imageID, needID)
}

// CreateNeedImage storages and links the uploaded file with the Need
func (s *ImageStorage) CreateNeedImage(t *model.Token, needID int64, fr *bytes.Reader) (*model.NeedImage, error) {
	n, err := s.NeedRepository.Get(needID)
	if err != nil {
		return nil, fmt.Errorf("there is no need with the id %d", needID)
	}

	if t.UserID != n.OrganizationID {
		return nil, fmt.Errorf("need %d does not belong to organization %d", needID, t.UserID)
	}

	i, err := s.createImage(fr, fmt.Sprintf("need-%d", needID))
	if err != nil {
		return nil, err
	}

	image := model.NeedImage{
		Image:  *i,
		NeedID: needID,
	}

	image, err = s.NeedRepository.CreateImage(image)
	if err != nil {
		s.Container.RemoveItem(i.URL)
		return nil, err
	}

	return &image, nil
}

// DeleteOrganizationImage removes the image from a need, if user has permission
func (s *ImageStorage) DeleteOrganizationImage(t *model.Token, imageID int64) error {
	o, err := s.OrganizationRepository.Get(t.UserID)
	if err != nil {
		return err
	}

	for _, i := range o.Images {
		if i.ID != imageID {
			continue
		}

		s.Container.RemoveItem(i.URL)
		s.OrganizationRepository.DeleteImage(i.ID, i.OrganizationID)
		return nil
	}

	return fmt.Errorf("there is no image with id %d at the organization %d", imageID, t.UserID)
}

// CreateOrganizationImage storages and link the image with the organization
func (s *ImageStorage) CreateOrganizationImage(t *model.Token, fr *bytes.Reader) (*model.OrganizationImage, error) {
	i, err := s.createImage(fr, fmt.Sprintf("organization-%d", t.UserID))
	if err != nil {
		return nil, err
	}

	image := model.OrganizationImage{
		Image:          *i,
		OrganizationID: t.UserID,
	}

	image, err = s.OrganizationRepository.CreateImage(image)
	if err != nil {
		s.Container.RemoveItem(i.URL)
		return nil, err
	}

	return &image, nil
}

func (s *ImageStorage) createImage(fr *bytes.Reader, folder string) (*model.Image, error) {
	fileName := uuid.Must(uuid.NewV4()).String() + ".png"
	filePath := fmt.Sprintf("%s/%s", folder, fileName)

	item, err := s.Container.Put(
		filePath,
		fr,
		fr.Size(),
		nil,
	)

	if err != nil {
		log.Printf("[ImageStorage] Error uploading file to container: %#v", err)
		return nil, fmt.Errorf("there was a problem saving the file")
	}

	i := model.Image{
		Name: fileName,
		URL:  item.ID(),
	}

	return &i, nil
}

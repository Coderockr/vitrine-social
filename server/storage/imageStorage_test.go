package storage_test

import (
	"errors"
	"mime/multipart"
	"testing"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/Coderockr/vitrine-social/server/storage"
	"github.com/stretchr/testify/require"
)

func TestCreateNeedImageShouldFail(t *testing.T) {

	type test struct {
		token  *model.Token
		needId int64
		fh     *multipart.FileHeader
		err    string
	}

	tests := map[string]test{
		"when_need_not_exists": test{
			token:  &model.Token{},
			needId: 404,
			fh:     nil,
			err:    "there is no need with the id 404",
		},
		"when_org_does_not_own_need": test{
			token:  &model.Token{UserID: 403},
			needId: 405,
			fh:     nil,
			err:    "need 405 does not belong to organization 403",
		},
		"when_fails_to_load_to_container": test{
			token:  &model.Token{UserID: 888},
			needId: 405,
			fh:     &multipart.FileHeader{Filename: "upload.png"},
			err:    "there was a problem with the file upload.png",
		},
	}

	iS := storage.ImageStorage{
		NeedRepository: &needRepositoryMock{
			GetFN: func(id int64) (*model.Need, error) {
				if id == 404 {
					return nil, errors.New("not found")
				}

				n := &model.Need{
					ID:             id,
					OrganizationID: 888,
				}

				return n, nil

			},
		},
	}

	for n, p := range tests {
		t.Run(n, func(t *testing.T) {
			_, err := iS.CreateNeedImage(p.token, p.needId, p.fh)
			require.Equal(t, err.Error(), p.err)
		})

	}
}

type needRepositoryMock struct {
	GetFN            func(int64) (*model.Need, error)
	CreateImageFN    func(model.NeedImage) (model.NeedImage, error)
	GetNeedsImagesFN func(model.Need) ([]model.NeedImage, error)
	DeleteImageFN    func(imageID, needID int64) error
}

func (m *needRepositoryMock) Get(id int64) (*model.Need, error) {
	return m.GetFN(id)
}

func (m *needRepositoryMock) CreateImage(ni model.NeedImage) (model.NeedImage, error) {
	return m.CreateImageFN(ni)
}

func (m *needRepositoryMock) GetNeedsImages(n model.Need) ([]model.NeedImage, error) {
	return m.GetNeedsImages(n)
}

func (m *needRepositoryMock) DeleteImage(imageID, needID int64) error {
	return m.DeleteImageFN(imageID, needID)
}

type orgRepositoryMock struct {
	GetFN         func(int64) (*model.Organization, error)
	CreateImageFN func(model.OrganizationImage) (model.OrganizationImage, error)
	DeleteImageFN func(imageID, orgID int64) error
}

func (m *orgRepositoryMock) Get(id int64) (*model.Organization, error) {
	return m.GetFN(id)
}

func (m *orgRepositoryMock) CreateImage(ni model.OrganizationImage) (model.OrganizationImage, error) {
	return m.CreateImageFN(ni)
}

func (m *orgRepositoryMock) DeleteImage(imageID, orgID int64) error {
	return m.DeleteImageFN(imageID, orgID)
}

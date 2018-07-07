package storage_test

import (
	"errors"
	"io"
	"mime/multipart"
	"net/url"
	"testing"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/Coderockr/vitrine-social/server/storage"
	"github.com/Coderockr/vitrine-social/server/testutils"
	"github.com/graymeta/stow"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDeleteOrganizationImage(t *testing.T) {

	repo := &orgRepositoryMock{}
	c := &containerMock{}

	oi := model.OrganizationImage{
		OrganizationID: 888,
		Image: model.Image{
			ID:  405,
			URL: "http://localhost/405.png",
		},
	}

	o := &model.Organization{
		User:   model.User{ID: oi.ID},
		Images: []model.OrganizationImage{oi},
	}

	repo.On("Get", o.ID).
		Once().
		Return(o, nil)

	repo.On("DeleteImage", oi.ID, oi.OrganizationID).
		Once().
		Return(nil)

	c.On("RemoveItem", oi.URL).
		Once().
		Return(nil)

	iS := storage.ImageStorage{
		OrganizationRepository: repo,
		Container:              c,
	}

	err := iS.DeleteOrganizationImage(
		&model.Token{UserID: o.ID},
		oi.ID,
	)

	require.Empty(t, err, "should've not fail")

	c.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestDeleteOrganizationImageShouldFail(t *testing.T) {
	type test struct {
		repo    *orgRepositoryMock
		token   *model.Token
		imageID int64
		err     string
	}

	tests := map[string]test{
		"when_org_does_not_exists": test{
			token: &model.Token{UserID: 888},
			repo: func() *orgRepositoryMock {
				r := &orgRepositoryMock{}
				r.On("Get", int64(888)).
					Return(nil, errors.New("not found"))
				return r
			}(),
			err: "not found",
		},
		"when_image_need_does_not_exist": test{
			token: &model.Token{UserID: 888},
			repo: func() *orgRepositoryMock {
				r := &orgRepositoryMock{}
				o := &model.Organization{
					User: model.User{ID: 888},
					Images: []model.OrganizationImage{
						model.OrganizationImage{
							Image: model.Image{
								ID:  405,
								URL: "http://localhost/image.png",
							},
						},
					},
				}

				r.On("Get", o.ID).Return(o, nil)

				return r
			}(),
			imageID: 404,
			err:     "there is no image with id 404 at the organization 888",
		},
	}

	for n, params := range tests {
		t.Run(n, func(t *testing.T) {

			iS := storage.ImageStorage{OrganizationRepository: params.repo}

			err := iS.DeleteOrganizationImage(
				params.token,
				params.imageID,
			)

			require.Equal(t, err.Error(), params.err)

			params.repo.AssertExpectations(t)
		})
	}
}

func TestDeleteNeedImage(t *testing.T) {

	repo := &needRepositoryMock{}
	c := &containerMock{}

	n := &model.Need{
		ID:             405,
		OrganizationID: 888,
	}

	repo.On("Get", int64(405)).
		Once().
		Return(n, nil)

	ni := model.NeedImage{
		NeedID: n.ID,
		Image: model.Image{
			ID:  305,
			URL: "http://localhost/305.png",
		},
	}

	repo.On("GetNeedsImages", *n).
		Once().
		Return([]model.NeedImage{ni}, nil)

	repo.On("DeleteImage", ni.ID, ni.NeedID).
		Once().
		Return(nil)

	c.On("RemoveItem", ni.URL).
		Once().
		Return(nil)

	iS := storage.ImageStorage{
		NeedRepository: repo,
		Container:      c,
	}

	err := iS.DeleteNeedImage(
		&model.Token{
			UserID: 888,
		},
		ni.NeedID,
		ni.ID,
	)

	require.Empty(t, err, "should've not fail")

	c.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestDeleteNeedImageShouldFail(t *testing.T) {
	type test struct {
		repo    *needRepositoryMock
		token   *model.Token
		needID  int64
		imageID int64
		err     string
	}

	tests := map[string]test{
		"when_need_does_not_exists": test{
			token: &model.Token{UserID: 888},
			repo: func() *needRepositoryMock {
				r := &needRepositoryMock{}
				r.On("Get", int64(404)).
					Return(nil, errors.New("not found"))
				return r
			}(),
			needID: 404,
			err:    "there is no need with the id 404",
		},
		"when_org_does_not_own_need": test{
			token: &model.Token{UserID: 888},
			repo: func() *needRepositoryMock {
				r := &needRepositoryMock{}
				r.On("Get", int64(405)).
					Return(
						&model.Need{
							ID:             405,
							OrganizationID: 777,
						},
						nil,
					)
				return r
			}(),
			needID: 405,
			err:    "need 405 does not belong to organization 888",
		},
		"when_fail_to_select_images": test{
			token: &model.Token{UserID: 888},
			repo: func() *needRepositoryMock {
				r := &needRepositoryMock{}
				n := &model.Need{
					ID:             405,
					OrganizationID: 888,
				}
				r.On("Get", int64(405)).
					Return(n, nil)

				r.On("GetNeedsImages", *n).
					Return([]model.NeedImage{}, errors.New("failed here"))

				return r
			}(),
			needID:  405,
			imageID: 404,
			err:     "failed here",
		},
		"when_image_need_does_not_exist": test{
			token: &model.Token{UserID: 888},
			repo: func() *needRepositoryMock {
				r := &needRepositoryMock{}
				n := &model.Need{
					ID:             405,
					OrganizationID: 888,
				}
				r.On("Get", int64(405)).
					Return(n, nil)

				r.On("GetNeedsImages", *n).
					Return([]model.NeedImage{model.NeedImage{}}, nil)

				return r
			}(),
			needID:  405,
			imageID: 404,
			err:     "there is no image with id 404 at the need 405",
		},
	}

	for n, params := range tests {
		t.Run(n, func(t *testing.T) {

			iS := storage.ImageStorage{NeedRepository: params.repo}

			err := iS.DeleteNeedImage(
				params.token,
				params.needID,
				params.imageID,
			)

			require.Equal(t, err.Error(), params.err)

			params.repo.AssertExpectations(t)
		})
	}
}

func TestCreateOrganizationImage(t *testing.T) {

	repo := &orgRepositoryMock{}
	c := &containerMock{}

	iM := &itemMock{}
	c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			iM.url = args.String(0)
			iM.md = args.Get(3).(map[string]interface{})

			require.Regexp(t, "^organization-888/.*\\.go$", iM.url)
		}).
		Return(iM, nil)

	call := repo.On("CreateImage", mock.Anything)
	call.Run(func(args mock.Arguments) {
		nI := args.Get(0).(model.OrganizationImage)
		nI.ID = 333
		call.Return(nI, nil)
	})

	iS := storage.ImageStorage{
		Container:              c,
		OrganizationRepository: repo,
	}

	r := testutils.NewFileUploadRequest(
		"/test",
		"POST",
		make(map[string]string),
		map[string]string{"images": "imageStorage_test.go"},
	)

	r.ParseMultipartForm(32 << 20)

	nI, err := iS.CreateOrganizationImage(
		&model.Token{UserID: 888},
		r.MultipartForm.File["images"][0],
	)

	require.Empty(t, err, "should've not fail")

	require.Equal(t, int64(333), nI.ID)
	require.Equal(t, "imageStorage_test", nI.Name)
	require.Equal(t, int64(888), nI.OrganizationID)
	require.Regexp(t, "^organization-888/.*\\.go$", nI.URL)

	c.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestCreateOrganizationImageShouldFail(t *testing.T) {
	type test struct {
		token     *model.Token
		fh        *multipart.FileHeader
		err       string
		container *containerMock
		repo      *orgRepositoryMock
	}

	r := testutils.NewFileUploadRequest(
		"/test",
		"POST",
		make(map[string]string),
		map[string]string{
			"to_fail":     "imageStorage_test.go",
			"not_to_fail": "imageStorage.go",
		},
	)

	r.ParseMultipartForm(32 << 20)

	tests := map[string]test{
		"when_fails_to_process_file": test{
			token: &model.Token{UserID: 888},
			fh:    &multipart.FileHeader{Filename: "upload.png"},
			err:   "there was a problem with the file upload.png",
		},
		"when_fails_to_load_to_container": test{
			token: &model.Token{UserID: 888},
			fh:    r.MultipartForm.File["to_fail"][0],
			err:   "there was a problem saving the file imageStorage_test.go",
			container: func() *containerMock {
				c := &containerMock{}
				c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New("fail to save it"))
				return c
			}(),
		},
		"when_fails_to_save_to_database": test{
			token: &model.Token{UserID: 888},
			fh:    r.MultipartForm.File["not_to_fail"][0],
			err:   "it have failed to save",
			container: func() *containerMock {
				c := &containerMock{}

				i := &itemMock{}
				c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Run(func(args mock.Arguments) {
						i.url = args.String(0)
						i.md = args.Get(3).(map[string]interface{})

						c.On("RemoveItem", i.url).Once().Return(nil)
					}).
					Return(i, nil)

				return c
			}(),
			repo: func() *orgRepositoryMock {
				repo := &orgRepositoryMock{}
				call := repo.On("CreateImage", mock.Anything)
				call.Run(func(args mock.Arguments) {
					call.Return(args.Get(0).(model.OrganizationImage), errors.New("it have failed to save"))
				})
				return repo
			}(),
		},
	}

	for n, p := range tests {
		t.Run(n, func(t *testing.T) {
			iS := storage.ImageStorage{
				Container:              p.container,
				OrganizationRepository: p.repo,
			}

			_, err := iS.CreateOrganizationImage(p.token, p.fh)
			require.Equal(t, err.Error(), p.err)

			if p.container != nil {
				p.container.AssertExpectations(t)
			}

			if p.repo != nil {
				p.repo.AssertExpectations(t)
			}
		})

	}
}

func TestCreateNeedImage(t *testing.T) {

	repo := &needRepositoryMock{}
	c := &containerMock{}

	repo.On("Get", int64(405)).
		Return(
			&model.Need{
				ID:             405,
				OrganizationID: 888,
			},
			nil,
		)

	iM := &itemMock{}
	c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			iM.url = args.String(0)
			iM.md = args.Get(3).(map[string]interface{})
		}).
		Return(iM, nil)

	call := repo.On("CreateImage", mock.Anything)
	call.Run(func(args mock.Arguments) {
		nI := args.Get(0).(model.NeedImage)
		nI.ID = 333
		call.Return(nI, nil)
	})

	iS := storage.ImageStorage{
		Container:      c,
		NeedRepository: repo,
	}

	r := testutils.NewFileUploadRequest(
		"/test",
		"POST",
		make(map[string]string),
		map[string]string{"images": "imageStorage_test.go"},
	)

	r.ParseMultipartForm(32 << 20)

	nI, err := iS.CreateNeedImage(
		&model.Token{
			UserID: 888,
		},
		405,
		r.MultipartForm.File["images"][0],
	)

	require.Empty(t, err, "should've not fail")

	require.Equal(t, int64(333), nI.ID)
	require.Equal(t, "imageStorage_test", nI.Name)
	require.Equal(t, int64(405), nI.NeedID)
	require.Regexp(t, "^need-405/.*\\.go$", nI.URL)

	c.AssertExpectations(t)
	repo.AssertExpectations(t)
}

func TestCreateNeedImageShouldFail(t *testing.T) {
	type test struct {
		token     *model.Token
		needID    int64
		fh        *multipart.FileHeader
		err       string
		container *containerMock
		needRepo  *needRepositoryMock
	}

	r := testutils.NewFileUploadRequest(
		"/test",
		"POST",
		make(map[string]string),
		map[string]string{
			"to_fail":     "imageStorage_test.go",
			"not_to_fail": "imageStorage.go",
		},
	)

	r.ParseMultipartForm(32 << 20)

	genNeedRepoMock := func(id int64) *needRepositoryMock {
		repo := &needRepositoryMock{}
		n := &model.Need{
			OrganizationID: 888,
			ID:             id,
		}
		repo.On("Get", id).
			Once().
			Return(n, nil)

		return repo
	}

	tests := map[string]test{
		"when_need_not_exists": test{
			token:  &model.Token{},
			needID: 404,
			fh:     nil,
			err:    "there is no need with the id 404",
			needRepo: func() *needRepositoryMock {
				repo := &needRepositoryMock{}
				repo.On("Get", mock.AnythingOfType("int64")).
					Return(nil, errors.New("not found"))
				return repo
			}(),
		},
		"when_org_does_not_own_need": test{
			token:    &model.Token{UserID: 403},
			needID:   405,
			fh:       nil,
			err:      "need 405 does not belong to organization 403",
			needRepo: genNeedRepoMock(405),
		},
		"when_fails_to_process_file": test{
			token:    &model.Token{UserID: 888},
			needID:   405,
			fh:       &multipart.FileHeader{Filename: "upload.png"},
			err:      "there was a problem with the file upload.png",
			needRepo: genNeedRepoMock(405),
		},
		"when_fails_to_load_to_container": test{
			token:  &model.Token{UserID: 888},
			needID: 405,
			fh:     r.MultipartForm.File["to_fail"][0],
			err:    "there was a problem saving the file imageStorage_test.go",
			container: func() *containerMock {
				c := &containerMock{}
				c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New("fail to save it"))
				return c
			}(),
			needRepo: genNeedRepoMock(405),
		},
		"when_fails_to_save_to_database": test{
			token:  &model.Token{UserID: 888},
			needID: 405,
			fh:     r.MultipartForm.File["not_to_fail"][0],
			err:    "it have failed to save",
			container: func() *containerMock {
				c := &containerMock{}

				i := &itemMock{}
				c.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Run(func(args mock.Arguments) {
						i.url = args.String(0)
						i.md = args.Get(3).(map[string]interface{})

						c.On("RemoveItem", i.url).Once().Return(nil)
					}).
					Return(i, nil)

				return c
			}(),
			needRepo: func() *needRepositoryMock {
				repo := &needRepositoryMock{}

				repo.On("Get", int64(405)).
					Return(
						&model.Need{
							ID:             405,
							OrganizationID: 888,
						},
						nil,
					)

				call := repo.On("CreateImage", mock.Anything)
				call.Run(func(args mock.Arguments) {
					call.Return(args.Get(0).(model.NeedImage), errors.New("it have failed to save"))
				})
				return repo
			}(),
		},
	}

	for n, p := range tests {
		t.Run(n, func(t *testing.T) {
			iS := storage.ImageStorage{
				Container:      p.container,
				NeedRepository: p.needRepo,
			}

			_, err := iS.CreateNeedImage(p.token, p.needID, p.fh)
			require.Equal(t, err.Error(), p.err)

			if p.container != nil {
				p.container.AssertExpectations(t)
			}
			p.needRepo.AssertExpectations(t)
		})

	}
}

type needRepositoryMock struct {
	mock.Mock
}

func (n *needRepositoryMock) Get(id int64) (*model.Need, error) {
	args := n.Called(id)
	if need := args.Get(0); need != nil {
		return need.(*model.Need), args.Error(1)
	}
	return nil, args.Error(1)
}

func (n *needRepositoryMock) CreateImage(ni model.NeedImage) (model.NeedImage, error) {
	args := n.Called(ni)
	return args.Get(0).(model.NeedImage), args.Error(1)
}

func (n *needRepositoryMock) GetNeedsImages(need model.Need) ([]model.NeedImage, error) {
	args := n.Called(need)
	return args.Get(0).([]model.NeedImage), args.Error(1)
}

func (n *needRepositoryMock) DeleteImage(imageID, needID int64) error {
	args := n.Called(imageID, needID)
	return args.Error(0)
}

type orgRepositoryMock struct {
	mock.Mock
}

func (m *orgRepositoryMock) Get(id int64) (*model.Organization, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Organization), args.Error(1)
}

func (m *orgRepositoryMock) CreateImage(ni model.OrganizationImage) (model.OrganizationImage, error) {
	args := m.Called(ni)
	return args.Get(0).(model.OrganizationImage), args.Error(1)
}

func (m *orgRepositoryMock) DeleteImage(imageID, orgID int64) error {
	args := m.Called(imageID, orgID)
	return args.Error(0)
}

type containerMock struct {
	mock.Mock
}

func (c *containerMock) ID() string {
	args := c.Called()
	return args.String(0)
}

func (c *containerMock) Name() string {
	args := c.Called()
	return args.String(0)
}

func (c *containerMock) Item(id string) (stow.Item, error) {
	args := c.Called(id)
	return args.Get(0).(stow.Item), args.Error(1)
}

func (c *containerMock) Items(prefix, cursor string, count int) ([]stow.Item, string, error) {
	args := c.Called(prefix, cursor, count)
	return args.Get(0).([]stow.Item), args.String(1), args.Error(3)
}

func (c *containerMock) RemoveItem(id string) error {
	args := c.Called(id)
	return args.Error(0)
}

func (c *containerMock) Put(name string, r io.Reader, size int64, metadata map[string]interface{}) (stow.Item, error) {
	args := c.Called(name, r, size, metadata)
	if i := args.Get(0); i == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(stow.Item), args.Error(1)
}

type itemMock struct {
	url string
	md  map[string]interface{}
}

func (i itemMock) ID() string {
	return i.url
}
func (i itemMock) Name() string {
	return i.ID()
}

func (i itemMock) URL() *url.URL {
	return nil
}
func (i itemMock) Size() (int64, error) {
	return 0, nil
}
func (i itemMock) Open() (io.ReadCloser, error) {
	return nil, nil
}
func (i itemMock) ETag() (string, error) {
	return i.ID(), nil
}

func (i itemMock) LastMod() (time.Time, error) {
	return time.Time{}, nil
}

func (i itemMock) Metadata() (map[string]interface{}, error) {
	return i.md, nil
}

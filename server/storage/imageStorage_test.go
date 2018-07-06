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
	testutil "github.com/Coderockr/vitrine-social/server/testutils"
	"github.com/graymeta/stow"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateNeedImage(t *testing.T) {
	assert := assert.New(t)

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

	r := testutil.NewFileUploadRequest(
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

	assert.Empty(err, "should've not fail")

	assert.Equal(int64(333), nI.ID)
	assert.Equal("imageStorage_test", nI.Name)
	assert.Equal(int64(405), nI.NeedID)
	assert.Regexp("^need-405/.*\\.go$", nI.URL)

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

	r := testutil.NewFileUploadRequest(
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

package graphql_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Coderockr/vitrine-social/server/graphql"
	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/context"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func getGraphqlRequest(t *model.Token, query string, vars *map[string]interface{}) *http.Request {
	params := map[string]interface{}{"query": query}

	if vars != nil {
		params["variables"] = vars
	}

	b, _ := json.Marshal(params)
	r := httptest.NewRequest("POST", "/graphql", bytes.NewBuffer(b))
	context.Set(r, handlers.TokenKey, t)
	return r
}

func TestAllCategoriesQuery(t *testing.T) {

	type testCase struct {
		repo     *catRepoMock
		query    string
		response string
	}

	tests := map[string]testCase{
		"when_success": testCase{
			query: `query {
				allCategories{
					name, slug
				}
			}`,
			response: `{"data": { "allCategories": [
					{"name":"Category1", "slug": "c1"},
					{"name":"Category2", "slug": "c2"}
				]
			}}`,
			repo: func() *catRepoMock {
				cMock := &catRepoMock{}
				cMock.On("GetAll").Once().
					Return(
						[]model.Category{
							model.Category{
								Name: "Category1",
								Slug: "c1",
							},
							model.Category{
								Name: "Category2",
								Slug: "c2",
							},
						},
						nil,
					)
				return cMock
			}(),
		},
		"when_fail": testCase{
			query: `query { allCategories{ name, slug } }`,
			response: `{"data": { "allCategories": null }, "errors":[
			{"message": "fail to query", "locations":[]}
			] }`,
			repo: func() *catRepoMock {
				cMock := &catRepoMock{}
				cMock.On("GetAll").Once().
					Return([]model.Category{}, errors.New("fail to query"))
				return cMock
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			h := graphql.NewHandler(
				&needRepoMock{},
				&orgRepoMock{},
				&tokenManagerMock{},
				test.repo,
				&searchRepoMock{},
				&imageStorageMock{},
			)

			w := httptest.NewRecorder()

			h.ServeHTTP(w, getGraphqlRequest(nil, test.query, nil))

			body, _ := ioutil.ReadAll(w.Result().Body)
			require.JSONEq(t, test.response, string(body))

			test.repo.AssertExpectations(t)
		})
	}
}

type needRepoMock struct {
	mock.Mock
}

func (m *needRepoMock) Get(id int64) (*model.Need, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Need), args.Error(1)
}

func (m *needRepoMock) GetNeedsImages(n model.Need) ([]model.NeedImage, error) {
	args := m.Called(n)
	return args.Get(0).([]model.NeedImage), args.Error(1)
}

func (m *needRepoMock) Create(n model.Need) (model.Need, error) {
	args := m.Called(n)
	return args.Get(0).(model.Need), args.Error(1)
}

func (m *needRepoMock) Update(n model.Need) (model.Need, error) {
	args := m.Called(n)
	return args.Get(0).(model.Need), args.Error(1)
}

type orgRepoMock struct {
	mock.Mock
}

func (m *orgRepoMock) Get(id int64) (*model.Organization, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Organization), args.Error(1)
}

func (m *orgRepoMock) Update(o model.Organization) (model.Organization, error) {
	args := m.Called(o)
	return args.Get(0).(model.Organization), args.Error(1)
}

func (m *orgRepoMock) GetUserByEmail(s string) (model.User, error) {
	args := m.Called(s)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *orgRepoMock) ChangePassword(o model.Organization, cPassword, nPassword string) (model.Organization, error) {
	args := m.Called(o, cPassword, nPassword)
	return args.Get(0).(model.Organization), args.Error(1)
}

func (m *orgRepoMock) ResetPasswordTo(o *model.Organization, s string) error {
	args := m.Called(o, s)
	return args.Error(1)
}

type catRepoMock struct {
	mock.Mock
}

func (m *catRepoMock) Get(id int64) (*model.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *catRepoMock) GetAll() ([]model.Category, error) {
	args := m.Called()
	return args.Get(0).([]model.Category), args.Error(1)
}

func (m *catRepoMock) GetNeedsCount(c *model.Category) (int64, error) {
	args := m.Called(c)
	return args.Get(0).(int64), args.Error(1)
}

type searchRepoMock struct {
	mock.Mock
}

func (m *searchRepoMock) Search(
	text string,
	categoriesID []int,
	organizationsID int64,
	status string,
	orderBy string,
	order string,
	page int,
) (results []model.SearchNeed, count int, err error) {
	args := m.Called(text, categoriesID, organizationsID, status, orderBy, order, page)
	return args.Get(0).([]model.SearchNeed), args.Int(1), args.Error(2)
}

type tokenManagerMock struct {
	mock.Mock
}

func (m *tokenManagerMock) CreateToken(u model.User, ps *[]string) (string, error) {
	args := m.Called(u, ps)
	return args.String(0), args.Error(1)
}

type imageStorageMock struct {
	mock.Mock
}

func (m *imageStorageMock) CreateNeedImage(t *model.Token, needID int64, fh *multipart.FileHeader) (*model.NeedImage, error) {
	args := m.Called(t, needID, fh)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.NeedImage), args.Error(1)
}
func (m *imageStorageMock) CreateOrganizationImage(t *model.Token, fh *multipart.FileHeader) (*model.OrganizationImage, error) {
	args := m.Called(t, fh)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.OrganizationImage), args.Error(1)
}
func (m *imageStorageMock) DeleteNeedImage(t *model.Token, needID, imageID int64) error {
	args := m.Called(t, needID, imageID)
	return args.Error(0)
}

func (m *imageStorageMock) DeleteOrganizationImage(t *model.Token, imageID int64) error {
	args := m.Called(t, imageID)
	return args.Error(0)
}

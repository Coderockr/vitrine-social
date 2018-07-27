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
	"github.com/Coderockr/vitrine-social/server/security"
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

func TestOpenQueries(t *testing.T) {
	type testCase struct {
		catRepo  *catRepoMock
		query    string
		response string
	}

	tests := map[string]testCase{
		"allCategories/when_success": testCase{
			query: `query {
				allCategories{
					name, slug
				}
			}`,
			response: `{"data": { "allCategories": [
				{"name":"Category1", "slug": "c1"},
				{"name":"Category2", "slug": "c2"}
			]}}`,
			catRepo: func() *catRepoMock {
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
		"allCategories/when_fail": testCase{
			query: `query { allCategories{ name, slug } }`,
			response: `{"data": { "allCategories": null }, "errors":[
				{"message": "fail to query", "locations":[]}
			] }`,
			catRepo: func() *catRepoMock {
				cMock := &catRepoMock{}
				cMock.On("GetAll").Once().
					Return([]model.Category{}, errors.New("fail to query"))
				return cMock
			}(),
		},
		"category/with_simple_query": testCase{
			query: `query {
				category(id: 1){
					name, slug
				}
			}`,
			response: `{"data": {
				"category": {"name":"Category1", "slug": "c1"}
			}}`,
			catRepo: func() *catRepoMock {
				cMock := &catRepoMock{}
				cMock.On("Get", int64(1)).Once().
					Return(
						&model.Category{
							Name: "Category1",
							Slug: "c1",
						},
						nil,
					)
				return cMock
			}(),
		},
		"category/with_needsCount": testCase{
			query: `query {
				category(id: 1){
					name, slug, needsCount
				}
			}`,
			response: `{"data": {
				"category": {"name":"Category1", "slug": "c1", "needsCount": 3}
			}}`,
			catRepo: func() *catRepoMock {
				cMock := &catRepoMock{}
				c := &model.Category{
					Name: "Category1",
					Slug: "c1",
				}
				cMock.On("Get", int64(1)).Once().
					Return(c, nil)

				cMock.On("GetNeedsCount", c).Once().
					Return(int64(3), nil)
				return cMock
			}(),
		},
		"category/when_fail": testCase{
			query: `query { category(id: 1){ name, slug } }`,
			response: `{"data": { "category": null }, "errors":[
				{"message": "fail to query", "locations":[]}
			] }`,
			catRepo: func() *catRepoMock {
				cMock := &catRepoMock{}
				cMock.On("Get", int64(1)).Once().
					Return(nil, errors.New("fail to query"))
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
				test.catRepo,
				&searchRepoMock{},
				&imageStorageMock{},
			)

			w := httptest.NewRecorder()

			h.ServeHTTP(w, getGraphqlRequest(nil, test.query, nil))

			body, _ := ioutil.ReadAll(w.Result().Body)
			require.JSONEq(t, test.response, string(body))

			test.catRepo.AssertExpectations(t)
		})
	}
}

func TestMutations(t *testing.T) {
	type testCase struct {
		tmMock      *tokenManagerMock
		orgRepoMock *orgRepoMock
		mutation    string
		response    string
	}

	tests := map[string]testCase{
		"login/when_success": testCase{
			mutation: `mutation {
				login(email:"admin@coderockr.com", password: "1234567"){
					token
					organization{id, name}
				}
			}`,
			response: `{"data": { "login": {
				"token": "tokengeradovalido",
				"organization": { "id": 1, "name": "Coderockr" }
			}}}`,
			orgRepoMock: func() *orgRepoMock {
				m := &orgRepoMock{}
				p, _ := security.GenerateHash("1234567")
				m.On("GetUserByEmail", "admin@coderockr.com").Once().
					Return(
						model.User{
							Email:    "admin@coderockr.com",
							ID:       1,
							Password: p,
						},
						nil,
					)

				m.On("Get", int64(1)).Once().
					Return(
						&model.Organization{
							User: model.User{ID: 1},
							Name: "Coderockr",
						},
						nil,
					)
				return m
			}(),
			tmMock: func() *tokenManagerMock {
				tm := &tokenManagerMock{}
				var pms *[]string
				tm.On("CreateToken", mock.Anything, pms).Once().
					Return(
						"tokengeradovalido",
						nil,
					)
				return tm
			}(),
		},
		"login/when_email_does_not_exists": testCase{
			mutation: `mutation {
				login(email:"admin@coderockr.com", password: "1234567"){
					token
					organization{id, name}
				}
			}`,
			response: `{"data": { "login": null }, "errors": [
				{"message": "email does not exist", "locations":[]}
			]}`,
			orgRepoMock: func() *orgRepoMock {
				m := &orgRepoMock{}
				m.On("GetUserByEmail", "admin@coderockr.com").Once().
					Return(model.User{}, errors.New("email does not exist"))

				return m
			}(),
		},
		"login/when_password_is_invalid": testCase{
			mutation: `mutation {
				login(email:"admin@coderockr.com", password: "1234"){
					token
					organization{id, name}
				}
			}`,
			response: `{"data": { "login": null }, "errors": [
				{"message": "password does not match", "locations":[]}
			]}`,
			orgRepoMock: func() *orgRepoMock {
				m := &orgRepoMock{}
				p, _ := security.GenerateHash("1234567")
				m.On("GetUserByEmail", "admin@coderockr.com").Once().
					Return(
						model.User{
							Email:    "admin@coderockr.com",
							ID:       1,
							Password: p,
						},
						nil,
					)

				return m
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			h := graphql.NewHandler(
				&needRepoMock{},
				test.orgRepoMock,
				test.tmMock,
				&catRepoMock{},
				&searchRepoMock{},
				&imageStorageMock{},
			)

			w := httptest.NewRecorder()

			h.ServeHTTP(w, getGraphqlRequest(nil, test.mutation, nil))

			body, _ := ioutil.ReadAll(w.Result().Body)
			require.JSONEq(t, test.response, string(body))

			if test.orgRepoMock != nil {
				test.orgRepoMock.AssertExpectations(t)
			}
			if test.tmMock != nil {
				test.tmMock.AssertExpectations(t)
			}
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

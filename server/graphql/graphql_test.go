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
	"github.com/lucassabreu/graphql-multipart-middleware/testutil"
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
		needRepo *needRepoMock
		orgRepo  *orgRepoMock
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
		"need/with_simple_query": testCase{
			query: `query { need(id: 444) { id, title } }`,
			response: `{"data": { "need": {
				"id": 444, "title": "a need"
			}}}`,
			needRepo: func() *needRepoMock {
				m := &needRepoMock{}
				m.On("Get", int64(444)).Once().
					Return(&model.Need{ID: 444, Title: "a need"}, nil)
				return m
			}(),
		},
		"need/with_full_query": testCase{
			query: `query { need(id: 444) { id, title, images { name, url }, category{slug}, organization{id, name} } }`,
			response: `{"data": { "need": {
				"id": 444, "title": "a need",
				"images": [ { "name": "a image", "url": "http://localhost/a-image.jpg" } ],
				"category": { "slug": "higiene" },
				"organization": {"id": 5, "name": "some org"}
			}}}`,
			needRepo: func() *needRepoMock {
				m := &needRepoMock{}
				m.On("Get", int64(444)).Once().
					Return(
						&model.Need{
							ID: 444, Title: "a need",
							Images: []model.NeedImage{
								model.NeedImage{
									Image: model.Image{
										Name: "a image",
										URL:  "http://localhost/a-image.jpg",
									},
								},
							},
							Category: model.Category{
								Slug: "higiene",
							},
							OrganizationID: 5,
						},
						nil,
					)
				return m
			}(),
			orgRepo: func() *orgRepoMock {
				m := &orgRepoMock{}
				m.On("Get", int64(5)).Once().
					Return(
						&model.Organization{
							User: model.User{ID: 5},
							Name: "some org",
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
				test.needRepo,
				test.orgRepo,
				&tokenManagerMock{},
				test.catRepo,
				&searchRepoMock{},
				&imageStorageMock{},
			)

			w := httptest.NewRecorder()

			h.ServeHTTP(w, getGraphqlRequest(nil, test.query, nil))

			body, _ := ioutil.ReadAll(w.Result().Body)
			require.JSONEq(t, test.response, string(body))

			if test.catRepo != nil {
				test.catRepo.AssertExpectations(t)
			}

			if test.needRepo != nil {
				test.needRepo.AssertExpectations(t)
			}

			if test.orgRepo != nil {
				test.orgRepo.AssertExpectations(t)
			}
		})
	}
}

func TestMutations(t *testing.T) {
	type testCase struct {
		tmMock       *tokenManagerMock
		orgRepoMock  *orgRepoMock
		needRepoMock *needRepoMock
		imageStorage *imageStorageMock
		token        *model.Token
		mutation     string
		response     string
	}

	dOrg := &model.Organization{
		User: model.User{ID: 1},
		Name: "Coderockr",
	}
	dToken := &model.Token{UserID: 1}

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
		"needCreate/when_success": testCase{
			token: dToken,
			mutation: `mutation {
				viewer {
					needCreate(input: {
						title: "new need",
						description: "a need",
						unit: "box",
						categoryId: 1,
						dueDate: "2018-07-01"
					}) {
						need { id, title, requiredQuantity }
					}
				}
			}`,
			response: `{"data":{  "viewer": { "needCreate": { "need" : {
				"id": 333, "title": "new need", "requiredQuantity": 0
			}}}}}`,
			needRepoMock: func() *needRepoMock {
				m := &needRepoMock{}
				call := m.On("Create", mock.AnythingOfType("Need"))
				call.Once().
					Run(func(args mock.Arguments) {
						n := args.Get(0).(model.Need)
						n.ID = 333
						call.Return(n, nil)
					})
				return m
			}(),
		},
		"needCreate/when_fail": testCase{
			token: dToken,
			mutation: `mutation {
				viewer {
					needCreate(input: {
						title: "new need",
						description: "a need",
						unit: "box",
						categoryId: 1,
					}) {
						need { id, title, requiredQuantity }
					}
				}
			}`,
			response: `{"data":{  "viewer": { "needCreate": null }}, "errors": [
				{"message": "fail to save", "locations":[] }
			]}`,
			needRepoMock: func() *needRepoMock {
				m := &needRepoMock{}
				call := m.On("Create", mock.AnythingOfType("Need"))
				call.Once().
					Run(func(args mock.Arguments) {
						call.Return(args.Get(0).(model.Need), errors.New("fail to save"))
					})
				return m
			}(),
		},
		"needImageDelete/when_sucess": testCase{
			token: dToken,
			mutation: `mutation {
				viewer {
					needImageDelete(input: {
						needId: 1,
						needImageId: 333,
					}) {
						need { id, title }
					}
				}
			}`,
			response: `{"data":{  "viewer": { "needImageDelete": { "need": {
				"title": "old need", "id": 1
			}}}}}`,
			imageStorage: func() *imageStorageMock {
				m := &imageStorageMock{}
				m.On("DeleteNeedImage", dToken, int64(1), int64(333)).Once().
					Return(nil)
				return m
			}(),
			needRepoMock: func() *needRepoMock {
				m := &needRepoMock{}
				m.On("Get", int64(1)).Once().
					Return(&model.Need{ID: 1, Title: "old need"}, nil)
				return m
			}(),
		},
		"needImageDelete/when_fail": testCase{
			token: dToken,
			mutation: `mutation {
				viewer {
					needImageDelete(input: {
						needId: 1,
						needImageId: 333,
					}) {
						need { id, title }
					}
				}
			}`,
			response: `{"data":{  "viewer": { "needImageDelete": null}}, "errors": [
				{"message": "it is not your image", "locations": []}
			]}`,
			imageStorage: func() *imageStorageMock {
				m := &imageStorageMock{}
				m.On("DeleteNeedImage", dToken, int64(1), int64(333)).Once().
					Return(errors.New("it is not your image"))
				return m
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.orgRepoMock == nil {
				test.orgRepoMock = &orgRepoMock{}
				test.orgRepoMock.On("Get", dToken.UserID).
					Return(dOrg, nil)
			}

			h := graphql.NewHandler(
				test.needRepoMock,
				test.orgRepoMock,
				test.tmMock,
				&catRepoMock{},
				&searchRepoMock{},
				test.imageStorage,
			)

			w := httptest.NewRecorder()

			h.ServeHTTP(w, getGraphqlRequest(test.token, test.mutation, nil))

			body, _ := ioutil.ReadAll(w.Result().Body)
			require.JSONEq(t, test.response, string(body))

			if test.orgRepoMock != nil {
				test.orgRepoMock.AssertExpectations(t)
			}

			if test.tmMock != nil {
				test.tmMock.AssertExpectations(t)
			}

			if test.needRepoMock != nil {
				test.needRepoMock.AssertExpectations(t)
			}

			if test.imageStorage != nil {
				test.imageStorage.AssertExpectations(t)
			}
		})
	}
}

func createUploadRequest(
	t *model.Token,
	mutation string,
	vars map[string]interface{},
	fileMap map[string][]string,
	files map[string]string,
) *http.Request {
	jsonFileMap, _ := json.Marshal(fileMap)

	if vars == nil {
		vars = make(map[string]interface{}, 0)
	}

	operation, _ := json.Marshal(map[string]interface{}{
		"query":     mutation,
		"variables": vars,
	})

	r := testutil.NewGraphQLFileUploadRequest(
		"/graphql",
		map[string]string{
			"operations": string(operation),
			"map":        string(jsonFileMap),
		},
		files,
	)
	context.Set(r, handlers.TokenKey, t)

	return r
}

func TestMutationsWithUpload(t *testing.T) {
	type testCase struct {
		orgRepoMock  *orgRepoMock
		imageStorage *imageStorageMock
		token        *model.Token
		mutation     string
		vars         map[string]interface{}
		fileMap      map[string][]string
		files        map[string]string
		response     string
	}

	dOrg := &model.Organization{
		User: model.User{ID: 1},
		Name: "Coderockr",
	}
	dToken := &model.Token{UserID: 1}

	tests := map[string]testCase{
		"needCreate/when_success": testCase{
			mutation: `mutation ($file: Upload!) {
				viewer {
					needImageCreate(input: { needId: 1, file: $file }){
						needImage { name, url }
					}
				}
			}`,
			vars:    map[string]interface{}{"file": nil},
			fileMap: map[string][]string{"graphql": []string{"variables.file"}},
			files:   map[string]string{"graphql": "graphql.go"},
			response: `{"data": { "viewer": { "needImageCreate": { "needImage": {
				"name": "graphql", "url": "http://localhost/need-1/aaaa.go"
			}}}}}`,
			token: dToken,
			imageStorage: func() *imageStorageMock {
				m := &imageStorageMock{}
				m.On("CreateNeedImage", dToken, int64(1), mock.AnythingOfType("*multipart.FileHeader")).Once().
					Return(
						&model.NeedImage{
							Image: model.Image{
								ID:   1,
								Name: "graphql",
								URL:  "http://localhost/need-1/aaaa.go",
							},
						},
						nil,
					)
				return m
			}(),
		},
		"needCreate/when_fail": testCase{
			mutation: `mutation ($file: Upload!) {
				viewer {
					needImageCreate(input: { needId: 1, file: $file }){
						needImage { name, url }
					}
				}
			}`,
			vars:    map[string]interface{}{"file": nil},
			fileMap: map[string][]string{"graphql": []string{"variables.file"}},
			files:   map[string]string{"graphql": "graphql.go"},
			response: `{"data": { "viewer": { "needImageCreate": null }}, "errors": [
				{"message": "no more space", "locations":[]}
			]}`,
			token: dToken,
			imageStorage: func() *imageStorageMock {
				m := &imageStorageMock{}
				m.On("CreateNeedImage", dToken, int64(1), mock.AnythingOfType("*multipart.FileHeader")).Once().
					Return(nil, errors.New("no more space"))
				return m
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.orgRepoMock == nil {
				test.orgRepoMock = &orgRepoMock{}
				test.orgRepoMock.On("Get", int64(1)).
					Return(dOrg, nil)
			}

			h := graphql.NewHandler(
				&needRepoMock{},
				test.orgRepoMock,
				&tokenManagerMock{},
				&catRepoMock{},
				&searchRepoMock{},
				test.imageStorage,
			)

			w := httptest.NewRecorder()

			h.ServeHTTP(
				w,
				createUploadRequest(
					test.token,
					test.mutation,
					test.vars,
					test.fileMap,
					test.files,
				),
			)

			body, _ := ioutil.ReadAll(w.Result().Body)
			require.JSONEq(t, test.response, string(body))

			if test.orgRepoMock != nil {
				test.orgRepoMock.AssertExpectations(t)
			}

			if test.imageStorage != nil {
				test.imageStorage.AssertExpectations(t)
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

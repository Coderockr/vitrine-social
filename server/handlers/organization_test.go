package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type (
	organizationRepositoryMock struct {
		GetFN         func(id int64) (*model.Organization, error)
		UpdateFN      func(o model.Organization) (model.Organization, error)
		DeleteImageFN func(imageID int64, organizationID int64) error
	}
)

func TestUpdateOrganizationHandler(t *testing.T) {
	type params struct {
		userID         int64
		organizationID string
		repository     handlers.OrganizationRepository
	}

	tests := map[string]struct {
		body     string
		status   int
		response string
		params   params
	}{
		"should fail beacuse trying to update another organization": {
			body:     ``,
			status:   http.StatusBadRequest,
			response: ``,
			params: params{
				userID:         2,
				organizationID: "1",
				repository: &organizationRepositoryMock{
					GetFN: func(id int64) (*model.Organization, error) {
						organization := model.Organization{}
						return &organization, nil
					},
					UpdateFN: func(o model.Organization) (model.Organization, error) {
						return o, nil
					},
				},
			},
		},
		"should success beacuse the right values were sent": {
			body: `{
				"name": "Novo Nome",
				"logo": "Novo Logo",
				"address": "Novo Endereço",
				"phone": "123123",
				"resume": "Nova Descrição detalhada da ONG",
				"video": "Novo Link do video",
				"email": "teste@coderockr.com.br"
			}`,
			status:   http.StatusNoContent,
			response: ``,
			params: params{
				userID:         1,
				organizationID: "1",
				repository: &organizationRepositoryMock{
					GetFN: func(id int64) (*model.Organization, error) {
						organization := model.Organization{
							User: model.User{
								Email:    "test@coderockr",
								Password: "",
								ID:       1,
							},
							Name:    "",
							Logo:    "",
							Address: "",
							Phone:   "",
							Resume:  "",
							Video:   "",
						}
						return &organization, nil
					},
					UpdateFN: func(o model.Organization) (model.Organization, error) {
						return o, nil
					},
				},
			},
		},
	}

	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := http.NewRequest("PUT", "/v1/organization/"+v.params.organizationID, strings.NewReader(v.body))
			r = mux.SetURLVars(r, map[string]string{"id": v.params.organizationID})
			context.Set(r, handlers.UserKey, v.params.userID)

			resp := httptest.NewRecorder()

			handlers.UpdateOrganizationHandler(v.params.repository)(resp, r)

			result := resp.Result()
			body, _ := ioutil.ReadAll(result.Body)

			if len(v.response) > 0 {
				require.JSONEq(t, v.response, string(body))
			}
			require.Equal(t, v.status, resp.Code)
		})
	}
}

func TestDeleteOrganizationImageHandler(t *testing.T) {
	type params struct {
		userID         int64
		organizationID string
		imageID        string
		repository     handlers.OrganizationRepository
	}

	tests := map[string]struct {
		body     string
		status   int
		response string
		params   params
	}{
		"should fail beacuse trying to remove imagem from another organization": {
			body:     ``,
			status:   http.StatusBadRequest,
			response: ``,
			params: params{
				userID:         2,
				organizationID: "1",
				imageID:        "2",
				repository: &organizationRepositoryMock{
					GetFN: func(id int64) (*model.Organization, error) {
						organization := model.Organization{}
						return &organization, nil
					},
					DeleteImageFN: func(imageID int64, organizationID int64) error {
						return nil
					},
				},
			},
		},
		"should success beacuse the right values were sent": {
			body:     ``,
			status:   http.StatusNoContent,
			response: ``,
			params: params{
				userID:         1,
				organizationID: "1",
				imageID:        "1",
				repository: &organizationRepositoryMock{
					GetFN: func(id int64) (*model.Organization, error) {
						organization := model.Organization{
							User: model.User{
								Email:    "test@coderockr",
								Password: "",
								ID:       1,
							},
							Name:    "",
							Logo:    "",
							Address: "",
							Phone:   "",
							Resume:  "",
							Video:   "",
						}
						return &organization, nil
					},
					DeleteImageFN: func(imageID int64, organizationID int64) error {
						return nil
					},
				},
			},
		},
	}

	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := http.NewRequest("DELETE", "/v1/organization/"+v.params.organizationID+"/image/"+v.params.imageID, strings.NewReader(v.body))
			r = mux.SetURLVars(r, map[string]string{"id": v.params.organizationID, "image_id": v.params.imageID})
			context.Set(r, handlers.UserKey, v.params.userID)

			resp := httptest.NewRecorder()

			handlers.DeleteOrganizationImageHandler(v.params.repository)(resp, r)

			result := resp.Result()
			body, _ := ioutil.ReadAll(result.Body)

			if len(v.response) > 0 {
				require.JSONEq(t, v.response, string(body))
			}
			require.Equal(t, v.status, resp.Code)
		})
	}
}

func (r *organizationRepositoryMock) Get(id int64) (*model.Organization, error) {
	return r.GetFN(id)
}

func (r *organizationRepositoryMock) Update(o model.Organization) (model.Organization, error) {
	return r.UpdateFN(o)
}

func (r *organizationRepositoryMock) DeleteImage(imageID int64, organizationID int64) error {
	return r.DeleteImageFN(imageID, organizationID)
}

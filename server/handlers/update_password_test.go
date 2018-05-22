package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/Coderockr/vitrine-social/server/security"
	"github.com/gorilla/context"
	"github.com/stretchr/testify/require"
)

type (
	UpdatePasswordOrganizationRepositoryMock struct {
		GetFN             func(id int64) (*model.Organization, error)
		ResetPasswordToFN func(o *model.Organization, password string) error
	}
)

func TestUpdatePasswordHandler(t *testing.T) {
	type params struct {
		userID     int64
		repository handlers.UpdatePasswordOrganizationRepository
	}

	tests := map[string]struct {
		body     string
		status   int
		response string
		params   params
	}{
		"should fail because invalid password": {
			body:     `{ "currentPassword": "test1", "newPassword": "newtest" }`,
			status:   http.StatusUnauthorized,
			response: `{ "code": 401, "message":"Senha inválida" }`,
			params: params{
				userID: 1,
				repository: &UpdatePasswordOrganizationRepositoryMock{
					GetFN: func(id int64) (*model.Organization, error) {
						password, err := security.GenerateHash("test")

						organization := &model.Organization{
							User: model.User{
								Email:    "test@coderockr",
								Password: password,
								ID:       1,
							},
						}
						return organization, err
					},
					ResetPasswordToFN: func(*model.Organization, string) error {
						return nil
					},
				},
			},
		},
		"should success beacuse the right values were sent": {
			body:     `{ "currentPassword": "test", "newPassword": "newtest" }`,
			status:   http.StatusOK,
			response: ``,
			params: params{
				userID: 1,
				repository: &UpdatePasswordOrganizationRepositoryMock{
					GetFN: func(id int64) (*model.Organization, error) {
						password, err := security.GenerateHash("test")

						organization := &model.Organization{
							User: model.User{
								Email:    "test@coderockr",
								Password: password,
								ID:       1,
							},
						}
						return organization, err
					},
					ResetPasswordToFN: func(*model.Organization, string) error {
						return nil
					},
				},
			},
		},
	}

	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := http.NewRequest("POST", "/v1/update-password", strings.NewReader(v.body))
			resp := httptest.NewRecorder()
			context.Set(r, handlers.UserKey, v.params.userID)

			handlers.UpdatePasswordHandler(v.params.repository)(resp, r)

			result := resp.Result()
			body, _ := ioutil.ReadAll(result.Body)

			if len(v.response) > 0 {
				require.JSONEq(t, v.response, string(body))
			}
			require.Equal(t, v.status, resp.Code)
		})
	}
}

func (r *UpdatePasswordOrganizationRepositoryMock) Get(id int64) (*model.Organization, error) {
	return r.GetFN(id)
}

func (r *UpdatePasswordOrganizationRepositoryMock) ResetPasswordTo(o *model.Organization, password string) error {
	return r.ResetPasswordToFN(o, password)
}

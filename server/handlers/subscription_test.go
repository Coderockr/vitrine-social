package handlers_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type (
	subscriptionRepositoryMock struct {
		CreateFN func(model.Subscription) (model.Subscription, error)
	}
)

func TestCreateSubscriptionHandler(t *testing.T) {
	type params struct {
		organizationID string
		repository     handlers.SubscriptionRepository
	}

	tests := map[string]struct {
		body     string
		status   int
		response string
		params   params
	}{
		"should fail beacuse trying to create without parameters": {
			body:     ``,
			status:   http.StatusBadRequest,
			response: ``,
			params: params{
				organizationID: "1",
				repository: &subscriptionRepositoryMock{
					CreateFN: func(model.Subscription) (model.Subscription, error) {
						s := model.Subscription{}
						return s, errors.New("Deve ser informado um nome para a Inscrição")
					},
				},
			},
		},
		"should fail beacuse trying to create with no valid organization": {
			body:     ``,
			status:   http.StatusBadRequest,
			response: ``,
			params: params{
				organizationID: "5",
				repository: &subscriptionRepositoryMock{
					CreateFN: func(model.Subscription) (model.Subscription, error) {
						s := model.Subscription{}
						return s, fmt.Errorf("Não foi encontrada Organização com ID: 5")
					},
				},
			},
		},
		"should success beacuse the right values were sent": {
			body: `{
				"name": "Coderockr Test",
				"email": "test@coderockr.com",
				"phone": "(54) 99999-9999"
			}`,
			status: http.StatusOK,
			response: `{
				"id": 1
			}`,
			params: params{
				organizationID: "1",
				repository: &subscriptionRepositoryMock{
					CreateFN: func(model.Subscription) (model.Subscription, error) {
						s := model.Subscription{
							ID:    1,
							Name:  "Coderockr Test",
							Email: "test@coderockr.com",
							Phone: "(54) 99999-9999",
						}
						return s, nil
					},
				},
			},
		},
	}

	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := http.NewRequest("POST", "/v1/organization/"+v.params.organizationID+"/subscribe", strings.NewReader(v.body))
			r = mux.SetURLVars(r, map[string]string{"id": v.params.organizationID})

			resp := httptest.NewRecorder()

			handlers.CreateSubscriptionHandler(v.params.repository)(resp, r)

			result := resp.Result()
			body, _ := ioutil.ReadAll(result.Body)

			if len(v.response) > 0 {
				require.JSONEq(t, v.response, string(body))
			}
			require.Equal(t, v.status, resp.Code)
		})
	}
}

func (r *subscriptionRepositoryMock) Create(s model.Subscription) (model.Subscription, error) {
	return r.CreateFN(s)
}

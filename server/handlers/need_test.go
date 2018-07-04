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
	needRepositoryMock struct {
		GetFN         func(id int64) (*model.Need, error)
		UpdateFN      func(o model.Need) (model.Need, error)
		CreateImageFN func(i model.NeedImage) (model.NeedImage, error)
	}
)

func TestUpdateNeedHandler(t *testing.T) {
	type params struct {
		userID     int64
		needID     string
		repository handlers.NeedRepository
	}

	tests := map[string]struct {
		body     string
		status   int
		response string
		params   params
	}{
		"should fail beacuse trying to update not owned need": {
			body: `{
				"category": 2,
				"title": "Need Title",
				"description": "Need Description",
				"requiredQuantity": 20,
				"reachedQuantity": 10,
				"dueDate": "2018-05-30",
				"unit": "KG",
				"status": "ACTIVE"
			  }`,
			status:   http.StatusBadRequest,
			response: ``,
			params: params{
				userID: 2,
				needID: "1",
				repository: &needRepositoryMock{
					GetFN: func(id int64) (*model.Need, error) {
						need := model.Need{
							OrganizationID: 1,
						}
						return &need, nil
					},
					UpdateFN: func(n model.Need) (model.Need, error) {
						return n, nil
					},
				},
			},
		},
		"should success beacuse the right values were sent": {
			body: `{
				"category": 2,
				"title": "Need Title",
				"description": "Need Description",
				"requiredQuantity": 20,
				"reachedQuantity": 10,
				"dueDate": "2018-05-30",
				"unit": "KG",
				"status": "ACTIVE"
			  }`,
			status:   http.StatusNoContent,
			response: ``,
			params: params{
				userID: 1,
				needID: "1",
				repository: &needRepositoryMock{
					GetFN: func(id int64) (*model.Need, error) {
						need := model.Need{
							OrganizationID: 1,
						}
						return &need, nil
					},
					UpdateFN: func(n model.Need) (model.Need, error) {
						return n, nil
					},
					CreateImageFN: func(i model.NeedImage) (model.NeedImage, error) {
						return i, nil
					},
				},
			},
		},
	}

	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := http.NewRequest("PUT", "/v1/need/"+v.params.needID, strings.NewReader(v.body))
			r = mux.SetURLVars(r, map[string]string{"id": v.params.needID})
			context.Set(r, handlers.TokenKey, &model.Token{UserID: v.params.userID})

			resp := httptest.NewRecorder()

			handlers.UpdateNeedHandler(v.params.repository)(resp, r)

			result := resp.Result()
			body, _ := ioutil.ReadAll(result.Body)

			if len(v.response) > 0 {
				require.JSONEq(t, v.response, string(body))
			}
			require.Equal(t, v.status, resp.Code)
		})
	}
}

func (r *needRepositoryMock) Get(id int64) (*model.Need, error) {
	return r.GetFN(id)
}

func (r *needRepositoryMock) Update(n model.Need) (model.Need, error) {
	return r.UpdateFN(n)
}

func (r *needRepositoryMock) CreateImage(i model.NeedImage) (model.NeedImage, error) {
	return r.CreateImageFN(i)
}

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
	"github.com/gorilla/context"

	"github.com/stretchr/testify/require"
)

func TestCreateNeedHandler(t *testing.T) {
	shouldNotCall := func(n model.Need) (model.Need, error) {
		t.Errorf("Should try to save !")
		t.FailNow()
		return n, nil
	}

	tests := map[string]struct {
		body     string
		userID   int64
		status   int
		create   func(model.Need) (model.Need, error)
		response string
	}{
		"should fail to read": {
			body:   "",
			status: http.StatusBadRequest,
			create: shouldNotCall,
		},
		"should fail to when body is incomplete": {
			body:     `{"organization": 2}`,
			userID:   1,
			status:   http.StatusBadRequest,
			create:   shouldNotCall,
			response: `{"code":400, "message":"você não possui permissão para criar uma necessidade para a orginização 2"}`,
		},
		"should fail because of business error": {
			body:   `{"organization": 1}`,
			userID: 1,
			status: http.StatusBadRequest,
			create: func(n model.Need) (model.Need, error) {
				return n, errors.New("failed because i want")
			},
			response: `{"code":400, "message":"failed because i want"}`,
		},
		"right values were sent": {
			body:   `{"organization": 1, "category": 99, "title": "test 1", "description": "test 2","requiredQuantity":3,"dueDate":"2017-10-01","unit":"KG"}`,
			userID: 1,
			status: http.StatusOK,
			create: func(n model.Need) (model.Need, error) {
				if n.OrganizationID != 1 ||
					n.CategoryID != 99 ||
					n.Title != "test 1" ||
					n.Description != "test 2" ||
					n.RequiredQuantity != 3 ||
					n.ReachedQuantity != 0 ||
					n.DueDate == nil ||
					n.DueDate.Format("2006-01-02") != "2017-10-01" {
					return n, fmt.Errorf("some values are not matching, values sent: %#v", n)
				}

				n.ID = 1
				return n, nil

			},
			response: `{"id":1}`,
		},
		"right values were sent (duedate is nil)": {
			body:   `{"organization": 1, "category": 99, "title": "test 1", "description": "test 2","requiredQuantity":3,"unit":"KG"}`,
			userID: 1,
			status: http.StatusOK,
			create: func(n model.Need) (model.Need, error) {
				if n.OrganizationID != 1 ||
					n.CategoryID != 99 ||
					n.Title != "test 1" ||
					n.Description != "test 2" ||
					n.RequiredQuantity != 3 ||
					n.ReachedQuantity != 0 ||
					n.DueDate != nil {
					return n, fmt.Errorf("some values are not matching, values sent: %#v", n)
				}

				n.ID = 1
				return n, nil

			},
			response: `{"id":1}`,
		},
	}

	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := http.NewRequest("POST", "/v1/need", strings.NewReader(v.body))
			resp := httptest.NewRecorder()
			context.Set(r, handlers.UserKey, v.userID)

			handlers.CreateNeedHandler(v.create)(resp, r)

			result := resp.Result()
			body, _ := ioutil.ReadAll(result.Body)

			if len(v.response) > 0 {
				require.JSONEq(t, v.response, string(body))
			}
			require.Equal(t, v.status, resp.Code)
		})
	}
}

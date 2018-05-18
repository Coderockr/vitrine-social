package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/gorilla/context"
	"github.com/stretchr/testify/require"
)

func TestUpdatePasswordHandler(t *testing.T) {
	tests := map[string]struct {
		body     string
		userID   int64
		status   int
		response string
	}{
		"should fail because invalid password": {
			body:     ``,
			status:   http.StatusBadRequest,
			response: "Senha inválida",
		},
		"right values were sent": {
			body:     `{ "currentPassword": "teste3", "newPassword": "teste3" }`,
			userID:   1,
			status:   http.StatusOK,
			response: "Senha atualizada com sucesso",
		},
	}

	for name, v := range tests {
		t.Run(name, func(t *testing.T) {
			r, _ := http.NewRequest("POST", "/v1/update-password", strings.NewReader(v.body))
			resp := httptest.NewRecorder()
			context.Set(r, handlers.UserKey, v.userID)

			//handlers.UpdatePasswordHandler(v.create)(resp, r)

			result := resp.Result()
			body, _ := ioutil.ReadAll(result.Body)

			if len(v.response) > 0 {
				require.JSONEq(t, v.response, string(body))
			}
			require.Equal(t, v.status, resp.Code)
		})
	}
}

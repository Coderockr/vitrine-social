package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Coderockr/vitrine-social/server/db/inmemory"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthRoute_Login(t *testing.T) {
	password, err := bcrypt.GenerateFromPassword([]byte("this is my password"), bcrypt.DefaultCost)
	userStorage := &inmemory.UserRepository{
		Storage: map[string]model.User{
			"jhon_doe": model.User{Email: "jhon_doe@gmail.com", ID: 1554, Password: string(password)},
		},
	}
	require.NoError(t, err)
	type fields struct {
		userStore UserRepository
		options   JWTOptions
	}
	type args struct {
		req  *http.Request
		resp string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"email_invalido",
			fields{
				userStore: userStorage,
				options: JWTOptions{
					Expiration:    time.Hour,
					PrivateKey:    []byte("this is my secure key"),
					SigningMethod: "HS256",
				},
			},
			args{
				req: httptest.NewRequest("POST", "http://vitrine/login",
					bytes.NewReader([]byte(`{"email": "jhon@gmail.com", "password": "this is my password"}`))),
				resp: `{"code": 401, "message": "Email não encontrado"}`,
			},
		},
		{
			"senha_invalida",
			fields{
				userStore: userStorage,
				options: JWTOptions{
					Expiration:    time.Hour,
					PrivateKey:    []byte("this is my secure key"),
					SigningMethod: "HS256",
				},
			},
			args{
				req: httptest.NewRequest("POST", "http://vitrine/login",
					bytes.NewReader([]byte(`{"email": "jhon_doe@gmail.com", "password": "this s my password"}`))),
				resp: `{"code": 401, "message": "Senha inválida"}`,
			},
		},
		{
			"email_senha_valido",
			fields{
				userStore: userStorage,
				options: JWTOptions{
					Expiration:    time.Hour,
					PrivateKey:    []byte("this is my secure key"),
					SigningMethod: "HS256",
				},
			},
			args{
				req: httptest.NewRequest("POST", "http://vitrine/login",
					bytes.NewReader([]byte(`{"email": "jhon_doe@gmail.com", "password": "this is my password"}`))),
				resp: `?`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthRoute{
				userStore: tt.fields.userStore,
				options:   tt.fields.options,
			}
			w := httptest.NewRecorder()
			a.Login(w, tt.args.req)
			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)
			require.JSONEq(t, tt.args.resp, string(body))
		})
	}
}

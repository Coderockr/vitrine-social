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

type (
	tokenManagerMock struct {
		CreateTokenFN   func(model.User) (string, error)
		ValidateTokenFN func(string) (int64, error)
	}
)

func TestAuthHandler_Login(t *testing.T) {
	password, err := bcrypt.GenerateFromPassword([]byte("this is my password"), bcrypt.DefaultCost)
	userStorage := &inmemory.UserRepository{
		Storage: map[string]model.User{
			"jhon_doe": {Email: "jhon_doe@gmail.com", ID: 1554, Password: string(password)},
		},
	}
	require.NoError(t, err)
	type fields struct {
		tokenManager TokenManager
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
				tokenManager: &tokenManagerMock{},
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
				tokenManager: &tokenManagerMock{},
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
				tokenManager: &tokenManagerMock{
					CreateTokenFN: func(u model.User) (string, error) {
						return "this-is-my-token", nil
					},
				},
			},
			args{
				req: httptest.NewRequest("POST", "http://vitrine/login",
					bytes.NewReader([]byte(`{"email": "jhon_doe@gmail.com", "password": "this is my password"}`))),
				resp: `{"token": "this-is-my-token"}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthHandler{
				UserGetter:   userStorage,
				TokenManager: tt.fields.tokenManager,
			}
			w := httptest.NewRecorder()
			a.Login(w, tt.args.req)
			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)
			require.JSONEq(t, tt.args.resp, string(body))
		})
	}
}

func TestJWTManager_CanReadItsOwnTokens(t *testing.T) {
	jwtOptions := map[string]JWTOptions{
		"using HS256": JWTOptions{
			SigningMethod: "HS256",
			PrivateKey:    []byte("ThisIsASecretSalt"),
			Expiration:    time.Hour,
		},
		// todo add tests to other supported methods
	}

	u := model.User{ID: 333}
	for name, opt := range jwtOptions {
		t.Run(name, func(t *testing.T) {
			tm := JWTManager{OP: opt}
			token, err := tm.CreateToken(u)
			if err != nil {
				t.Fatalf("Should not fail with: %s", err.Error())
				t.FailNow()
			}

			id, err := tm.ValidateToken(token)
			if err != nil {
				t.Fatalf("Should not fail with: %s", err.Error())
				t.FailNow()
			}

			require.Equal(t, id, u.ID)
		})
	}

}

func (t *tokenManagerMock) CreateToken(user model.User) (string, error) {
	return t.CreateTokenFN(user)
}

func (t *tokenManagerMock) ValidateToken(token string) (int64, error) {
	return t.ValidateTokenFN(token)
}

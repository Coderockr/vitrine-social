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
		CreateTokenFN   func(model.User, *[]string) (string, error)
		ValidateTokenFN func(string) (*model.Token, error)
	}
)

func TestAuthHandler_Login(t *testing.T) {
	password, err := bcrypt.GenerateFromPassword([]byte("this is my password"), bcrypt.DefaultCost)
	organizationStorage := &inmemory.OrganizationRepository{
		Storage: map[string]model.Organization{
			"jhon_doe": {
				User: model.User{
					Email:    "jhon_doe@gmail.com",
					ID:       1554,
					Password: string(password),
				},
				Name:  "Jhon Doe",
				Logo:  "Logo",
				Slug:  "jhon_doe",
				Phone: "123",
			},
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
					CreateTokenFN: func(u model.User, p *[]string) (string, error) {
						return "this-is-my-token", nil
					},
				},
			},
			args{
				req: httptest.NewRequest("POST", "http://vitrine/login",
					bytes.NewReader([]byte(`{"email": "jhon_doe@gmail.com", "password": "this is my password"}`))),
				resp: `
					{
						"organization": {
							"id": 1554,
							"name": "Jhon Doe",
							"logo": "Logo",
							"slug": "jhon_doe",
							"phone": ""
						},
						"token": "this-is-my-token"
					}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthHandler{
				OrganizationGetter: organizationStorage,
				TokenManager:       tt.fields.tokenManager,
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
			token, err := tm.CreateToken(u, nil)
			if err != nil {
				t.Fatalf("Should not fail with: %s", err.Error())
				t.FailNow()
			}

			tk, err := tm.ValidateToken(token)
			if err != nil {
				t.Fatalf("Should not fail with: %s", err.Error())
				t.FailNow()
			}

			require.Equal(t, tk.UserID, u.ID)
		})
	}

}

func (t *tokenManagerMock) CreateToken(user model.User, ps *[]string) (string, error) {
	return t.CreateTokenFN(user, ps)
}

func (t *tokenManagerMock) ValidateToken(token string) (*model.Token, error) {
	return t.ValidateTokenFN(token)
}

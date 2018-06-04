package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/Coderockr/vitrine-social/server/security"
	"github.com/gorilla/context"
)

const (
	// TokenKey is the context key for the JWT token
	TokenKey = "token"
	// UserKey is the context key for the autheticathed user
	UserKey = "user"
)

type (
	// OrganizationGetter represent the operations for retrieve some organization.
	OrganizationGetter interface {
		GetByEmail(email string) (*model.Organization, error)
	}

	// TokenManager represet operations for application tokens.
	TokenManager interface {
		CreateToken(model.User) (string, error)
		ValidateToken(string) (int64, error)
	}

	// AuthHandler represent all the handler endpoints and middlewares.
	AuthHandler struct {
		OrganizationGetter OrganizationGetter
		TokenManager       TokenManager
	}

	loginJSON struct {
		Organization baseOrganizationJSON `json:"organization"`
		Token        string               `json:"token"`
	}
)

// Login process the login requests, returning a JWT token and organization data
func (a *AuthHandler) Login(w http.ResponseWriter, req *http.Request) {
	var authForm map[string]string

	err := requestToJSONObject(req, &authForm)
	if err != nil {
		HandleHTTPError(w, http.StatusBadRequest, err)
		return
	}

	email := authForm["email"]
	pass := authForm["password"]

	organization, err := a.OrganizationGetter.GetByEmail(email)
	user := organization.User
	if err != nil {
		log.Printf("[INFO][Auth Handler] %s", err.Error())
		HandleHTTPError(w, http.StatusUnauthorized, errors.New("Email não encontrado"))
		return
	}
	err = security.CompareHashAndPassword(user.Password, pass)
	if err != nil {
		HandleHTTPError(w, http.StatusUnauthorized, errors.New("Senha inválida"))
		return
	}

	token, err := a.TokenManager.CreateToken(user)
	if err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, errors.New("Error while Signing Token"))
		return
	}

	json := loginJSON{
		Token: token,
		Organization: baseOrganizationJSON{
			ID:   organization.ID,
			Name: organization.Name,
			Logo: organization.Logo,
			Slug: organization.Slug,
		},
	}

	HandleHTTPSuccess(w, json)
}

// GetUserID retorna o id do usuário logado.
func GetUserID(r *http.Request) int64 {
	return context.Get(r, UserKey).(int64)
}

// GetToken retorna o token do usuário logado
func GetToken(r *http.Request) string {
	return context.Get(r, TokenKey).(string)
}

// AuthMiddleware valida o token e filtra usuários não logados corretamente
func (a *AuthHandler) AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := r.Header.Get("Authorization")
	if token == "" {
		HandleHTTPError(w, http.StatusUnauthorized, errors.New("Error no token is provided"))
		return
	}
	userID, err := a.TokenManager.ValidateToken(token)

	if err != nil {
		HandleHTTPError(w, http.StatusUnauthorized, err)
		return
	}

	context.Set(r, TokenKey, token)
	context.Set(r, UserKey, userID)
	next(w, r)
	context.Clear(r)
}

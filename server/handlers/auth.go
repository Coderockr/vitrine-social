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
	tokenKey = "token"
	userKey  = "user"
)

type (
	// UserGetter represent the operations for retrieve some application user.
	UserGetter interface {
		GetUserByEmail(email string) (model.User, error)
	}

	// TokenManager represet operations for application tokens.
	TokenManager interface {
		CreateToken(model.User) (string, error)
		ValidateToken(string) (int64, error)
	}

	// AuthHandler represent all the handler endpoints and middlewares.
	AuthHandler struct {
		UserGetter   UserGetter
		TokenManager TokenManager
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

	user, err := a.UserGetter.GetUserByEmail(email)
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

	HandleHTTPSuccess(w, map[string]string{"token": token})
}

// GetUserID retorna o id do usuário logado.
func GetUserID(r *http.Request) int64 {
	return context.Get(r, userKey).(int64)
}

// GetToken retorna o token do usuário logado
func GetToken(r *http.Request) string {
	return context.Get(r, tokenKey).(string)
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

	context.Set(r, tokenKey, token)
	context.Set(r, userKey, userID)
	next(w, r)
	context.Clear(r)

}

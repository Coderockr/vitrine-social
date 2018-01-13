package handlers

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenKey = "token"
	userKey  = "user"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
}

type AuthRoute struct {
	userStore UserRepository
	options   JWTOptions
}

func (a *AuthRoute) Login(w http.ResponseWriter, req *http.Request) {
	var authForm map[string]string

	err := requestToJSONObject(req, &authForm)
	if err != nil {
		HandleHTTPError(w, http.StatusBadRequest, err)
		return
	}

	email := authForm["email"]
	pass := authForm["password"]

	user, err := a.userStore.GetUserByEmail(email)
	if err != nil {
		HandleHTTPError(w, http.StatusUnauthorized, errors.New("Email não encontrado"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		HandleHTTPError(w, http.StatusUnauthorized, errors.New("Senha inválida"))
		return
	}

	token, err := generateJWTToken(100, a.options)
	if err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, errors.New("Error while Signing Token"))
		return
	}

	HandleHTTPSuccess(w, map[string]string{"token": token})
}

func (a *AuthRoute) authenticate(w http.ResponseWriter, r *http.Request) (int64, string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return 0, "", errors.New("Error no token is provided")
	}
	userID, token, err := validateToken(r, a.options)
	if err != nil {
		return 0, "", err
	}
	return userID, token, nil
}

// GetUserID retorna o id do usuário logado.
func GetUserID(r *http.Request) string {
	return context.Get(r, userKey).(string)
}

// GetToken retorna o token do usuário logado
func GetToken(r *http.Request) string {
	return context.Get(r, tokenKey).(string)
}

// AuthMiddleware valida o token e filtra usuários não logados corretamente
func (a *AuthRoute) AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userID, token, err := a.authenticate(w, r)

	if err != nil {
		HandleHTTPError(w, http.StatusUnauthorized, err)
		return
	}

	context.Set(r, tokenKey, token)
	context.Set(r, userKey, userID)
	next(w, r)
	context.Clear(r)

}

func requestToJSONObject(req *http.Request, jsonDoc interface{}) error {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(jsonDoc)
}

func generateRandomKey(strength int) []byte {
	k := make([]byte, strength)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}

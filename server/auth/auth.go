package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/context"
)

const (
	TokenKey = "token"
	UserKey  = "user"
)

type AuthRoute struct {
	userStore UserRepository
	options   Options
}

func NewAuthRoute(store UserRepository, opt Options) *AuthRoute {
	return &AuthRoute{
		userStore: store,
		options:   opt,
	}
}

func (a *AuthRoute) Login(w http.ResponseWriter, req *http.Request) {
	var authForm map[string]string

	err := requestToJsonObject(req, &authForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := authForm["username"]
	pass := authForm["password"]

	userID, err := a.userStore.Login(email, pass)
	if err != nil {
		http.Error(w, "Username or Password Invalid", http.StatusUnauthorized)
		return
	}

	token, err := generateJWTToken(userID, a.options)
	if err != nil {
		http.Error(w, "Error while Signing Token :S", http.StatusInternalServerError)
		return
	}

	jtoken, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		http.Error(w, "Error marshalling the token to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jtoken)
}

func (a *AuthRoute) RefreshToken(w http.ResponseWriter, req *http.Request) {
	userID, _, err := a.authenticate(w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := generateJWTToken(userID, a.options)
	if err != nil {
		http.Error(w, "Error while Signing Token :S", http.StatusInternalServerError)
		return
	}

	jtoken, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		http.Error(w, "Error marshalling the token to json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jtoken)
}

func (a *AuthRoute) authenticate(w http.ResponseWriter, r *http.Request) (int64, string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return 0, "", errors.New("Error no token is provided")
	}
	userID, token, err := validateToken(r, a.options.PublicKey)
	if err != nil {
		return 0, "", err
	}
	return userID, token, nil
}

// Get User from the context
func GetuserID(r *http.Request) string {
	return context.Get(r, UserKey).(string)
}

// Get Token from the context
func GetToken(r *http.Request) string {
	return context.Get(r, TokenKey).(string)
}

// Auth middleware for negroni
func (a *AuthRoute) AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userID, token, err := a.authenticate(w, r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	context.Set(r, TokenKey, token)
	context.Set(r, UserKey, userID)
	next(w, r)
	context.Clear(r)

}

// Auth Handler for net/http
func (a *AuthRoute) AuthHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, token, err := a.authenticate(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		context.Set(r, TokenKey, token)
		context.Set(r, UserKey, userID)
		h.ServeHTTP(w, r)
		context.Clear(r)
	})
}

func (a *AuthRoute) AuthHandlerFunc(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.AuthMiddleware(w, r, next)
	})
}

func requestToJsonObject(req *http.Request, jsonDoc interface{}) error {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(jsonDoc)
	if err != nil {
		return err
	}
	return nil
}

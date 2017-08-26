package middlewares

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Auth struct {
	session sessions.Store
}

func NewAuth(s sessions.Store) *Auth {
	return &Auth{
		session: s,
	}
}

//IsAuthenticated verifica se o usuário tem uma sessão
func (a *Auth) IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	session, err := a.session.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["profile"]; !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		next(w, r)
	}
}

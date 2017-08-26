package middlewares

import (
	"net/http"
)

//Cors adiciona os headers para suportar o CORS nos navegadores
func Cors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	if r.Method == "OPTIONS" {
		return
	}
	next(w, r)
}

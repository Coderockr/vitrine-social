package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//Crawlers return html with metatags to crawlers like facebook
func Crawlers(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userAgent := r.Header.Get("User-Agent")

	switch userAgent {
	case
		"Googlebot",
		"facebookexternalhit":
		fmt.Printf("user agent is a robot: %s \n", userAgent)

		file, _ := ioutil.ReadFile("./index.html")
		html := string(file)
		html = strings.Replace(html, "__META_URL__", "teste", 1)
		html = strings.Replace(html, "__META_TITLE__", "teste", 1)
		html = strings.Replace(html, "__META_DESCRIPTION__", "teste", 1)
		html = strings.Replace(html, "__META_IMAGE__", "teste", 1)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)

		return
	}

	next(w, r)
}

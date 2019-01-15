package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// ShareNeedHandler return html with metatags for share
func ShareNeedHandler(repo NeedRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		need, _ := repo.Get(id)
		if need == nil {
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Necessidade não encontrada"))
			return
		}

		apiURL := os.Getenv("API_URL")
		frontendURL := os.Getenv("FRONTEND_URL")
		imageURL := os.Getenv("IMAGE_STORAGE_BASE_URL")

		const html = `<!doctype html>
			<html>
				<head>
					<meta http-equiv="refresh" content="0;url={{.RedirectURL}}">
					<meta property="og:locale" content="pt_BR">
					<meta property="og:site_name" content="Vitrine Social">
					<meta property="og:type" content="website">
					<meta property="og:url" content="{{.MetaURL}}">
					<meta property="og:title" content="{{.MetaTitle}}">
					<meta property="og:description" content="{{.MetaDescription}}">
					<meta property="og:image" content="{{.MetaImage}}">
					<meta property="og:image:type" content="image/jpeg">
				</head>
				<body></body>
			</html>`

		t, err := template.New("share").Parse(html)

		data := struct {
			RedirectURL     string
			MetaURL         string
			MetaTitle       string
			MetaDescription string
			MetaImage       string
		}{
			RedirectURL:     frontendURL + "/detalhes/" + vars["id"],
			MetaURL:         apiURL + r.URL.String(),
			MetaTitle:       need.Title,
			MetaDescription: need.Description.String,
			MetaImage:       imageURL + "general/share.jpg",
		}

		_ = t.Execute(w, data)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")

		return
	}
}

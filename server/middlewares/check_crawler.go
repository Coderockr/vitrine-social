package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	// CheckCrawler is a implementation middleware
	CheckCrawler struct {
		needRepo NeedRepository
	}

	// NeedRepository represet operations for need repository.
	NeedRepository interface {
		Get(id int64) (*model.Need, error)
	}
)

// NewCheckCrawler creates a new check crawler middleware
func NewCheckCrawler(nR NeedRepository) *CheckCrawler {
	return &CheckCrawler{
		needRepo: nR,
	}
}

//CheckCrawler return html with metatags to crawlers like facebook
func (c *CheckCrawler) CheckCrawler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userAgent := r.Header.Get("User-Agent")

	switch userAgent {
	case
		"Googlebot",
		"facebookexternalhit":

		s := strings.Split(r.URL.String(), "/")
		idString := s[len(s)-1]
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			handlers.HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", idString))
			return
		}

		need, _ := c.needRepo.Get(id)

		url := os.Getenv("FRONTEND_URL")
		imageURL := os.Getenv("IMAGE_STORAGE_BASE_URL")

		file, _ := ioutil.ReadFile("./share.html")
		html := string(file)
		html = strings.Replace(html, "__META_URL__", url+"/detalhes/"+idString, 1)
		html = strings.Replace(html, "__META_TITLE__", need.Title, 1)
		html = strings.Replace(html, "__META_DESCRIPTION__", need.Description.String, 1)
		html = strings.Replace(html, "__META_IMAGE__", imageURL+"general/share.jpg", 1)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)

		return
	}

	next(w, r)
}

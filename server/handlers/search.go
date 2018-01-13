package handlers

import (
	"errors"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/index"
)

// SearchHandler handles requests about organizations
type SearchHandler struct {
	indexService index.Service
	needRepo     *repo.NeedRepository
}

//NewSearchHandler search handler
func NewSearchHandler(indexService index.Service, needRepo *repo.NeedRepository) *SearchHandler {
	return &SearchHandler{
		indexService: indexService,
		needRepo:     needRepo,
	}
}

//Get handler
func (sR *SearchHandler) Get(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["query"]

	if !ok || len(keys) < 1 {
		err := errors.New("Invalid parameters")
		HandleHttpError(w, http.StatusBadRequest, err)
		return
	}
	query := keys[0]
	docs, err := sR.indexService.Search(query)
	if err != nil {
		return nil, err
	}
	var needs []*needJSON
	for _, j := range docs.Hits {
		p, err := sr.needRepo.Get(j.ID)
		if err != nil {
			continue
		}

		if p != nil {
			needs = append(needs, p)
		}
	}

	return needs, nil
}

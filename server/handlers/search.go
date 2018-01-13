package handlers

import (
	"errors"
	"net/http"

	"github.com/coderockr/vitrine-social/server/index"
)

// SearchHandler handles requests about organizations
type SearchHandler struct {
	indexService index.Service
}

//NewSearchHandler search handler
func NewSearchHandler(indexService index.Service) *SearchHandler {
	return &SearchHandler{
		indexService: indexService,
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
	docs, err := sR.indexService.Search(q)
	if err != nil {
		return nil, err
	}
	var pos []*Position
	for _, j := range docs.Hits {
		p, err := s.Find(bson.ObjectIdHex(j.ID))
		if err != nil {
			continue
		}

		if p != nil {
			pos = append(pos, p)
		}
	}

	return pos, nil
}

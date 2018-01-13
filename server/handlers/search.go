package handlers

import (
	"errors"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/index"
	"github.com/Coderockr/vitrine-social/server/model"
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
		HandleHTTPError(w, http.StatusBadRequest, err)
		return
	}
	query := keys[0]
	docs, err := sR.indexService.Search(query)
	if err != nil {
		HandleHTTPError(w, http.StatusBadRequest, err)
	}
	var needs []*needJSON
	for _, j := range docs.Hits {
		p, err := sR.needRepo.Get(j.ID)
		if err != nil {
			continue
		}

		if p != nil {
			needs = append(needs, needToJSON(p))
		}
	}

	HandleHTTPSuccess(w, needs)
}

func needToJSON(n *model.Need) *needJSON {
	nJSON := &needJSON{
		ID:               n.ID,
		Title:            n.Title,
		Description:      n.Description,
		RequiredQuantity: n.RequiredQuantity,
		ReachedQuantity:  n.ReachedQuantity,
		Unity:            n.Unity,
		Category: categoryJSON{
			ID:   n.Category.ID,
			Name: n.Category.Name,
			Icon: n.Category.Icon,
		},
		// Organization: baseOrganizationJSON{
		// 	ID:   o.ID,
		// 	Name: o.Name,
		// 	Logo: o.Logo,
		// 	Slug: o.Slug,
		// },
		Images: needImagesToImageJSON(n.Images),
		Status: n.Status,
	}
	return nJSON
}

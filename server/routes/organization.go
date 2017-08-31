package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

type baseOrganizationJSON struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Slug string `json:"slug"`
}

type organizationJSON struct {
	baseOrganizationJSON
	Address string      `json:"address"`
	Phone   string      `json:"phone"`
	Resume  string      `json:"resume"`
	Video   string      `json:"video"`
	Email   string      `json:"email"`
	Needs   []needJSON  `json:"needs"`
	Images  []imageJSON `json:"images"`
}

type categoryJSON struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type imageJSON struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type needJSON struct {
	ID               int64                `json:"id"`
	Category         categoryJSON         `json:"category"`
	Organization     baseOrganizationJSON `json:"organization"`
	Images           []imageJSON          `json:"images"`
	Title            string               `json:"title"`
	Description      string               `json:"description"`
	RequiredQuantity int                  `json:"requiredQuantity"`
	ReachedQuantity  int                  `json:"reachedQuantity"`
	Unity            string               `json:"unity"`
	DueDate          *string              `json:"dueDate"`
}

// OrganizationRepository has the commands needed for this route
type OrganizationRepository interface {
	Get(id int64) (*model.Organization, error)
}

// OrganizationRoute handles requests about organizations
type OrganizationRoute struct {
	repo OrganizationRepository
}

// NewOrganizationRoute creates a new OrganizationRoute
func NewOrganizationRoute(repo OrganizationRepository) *OrganizationRoute {
	return &OrganizationRoute{
		repo: repo,
	}
}

// Get will retrive the data from a organization
func (oR *OrganizationRoute) Get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		HandleHttpError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
		return
	}

	o, err := oR.repo.Get(id)

	switch {
	case err == sql.ErrNoRows:
		HandleHttpError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Organização com ID: %d", id))
		return
	case err != nil:
		HandleHttpError(w, http.StatusForbidden, err)
		return
	}

	oJSON := &organizationJSON{
		baseOrganizationJSON: baseOrganizationJSON{
			ID:   o.ID,
			Name: o.Name,
			Logo: o.Logo,
			Slug: o.Slug,
		},
		Address: o.Address,
		Phone:   o.Phone,
		Resume:  o.Resume,
		Video:   o.Video,
		Email:   o.Email,
		Images:  orgImagesToImageJSON(o.Images),
	}

	oJSON.Needs = make([]needJSON, 0, len(o.Needs))
	catMap := make(map[int64]categoryJSON)

	for _, n := range o.Needs {
		if _, ok := catMap[n.CategoryID]; ok == false {
			catMap[n.CategoryID] = categoryJSON{
				ID:   n.Category.ID,
				Name: n.Category.Name,
				Icon: n.Category.Icon,
			}
		}

		oJSON.Needs = append(oJSON.Needs, needJSON{
			ID:               n.ID,
			Title:            n.Title,
			Description:      n.Description,
			RequiredQuantity: n.RequiredQuantity,
			ReachedQuantity:  n.ReachedQuantity,
			Unity:            n.Unity,
			DueDate:          n.DueDate,
			Category:         catMap[n.CategoryID],
			Organization:     oJSON.baseOrganizationJSON,
			Images:           needImagesToImageJSON(n.Images),
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(oJSON)
}

func needImagesToImageJSON(images []model.NeedImage) []imageJSON {
	imagesJSON := make([]imageJSON, 0, len(images))
	for _, i := range images {
		imagesJSON = append(imagesJSON, imageJSON{
			ID:   i.ID,
			Name: i.Name,
			URL:  i.URL,
		})
	}
	return imagesJSON
}

func orgImagesToImageJSON(images []model.OrganizationImage) []imageJSON {
	imagesJSON := make([]imageJSON, 0, len(images))
	for _, i := range images {
		imagesJSON = append(imagesJSON, imageJSON{
			ID:   i.ID,
			Name: i.Name,
			URL:  i.URL,
		})
	}
	return imagesJSON
}

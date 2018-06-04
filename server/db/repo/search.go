package repo

import (
	"fmt"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

// SearchRepository is an implementation for Postgres
type SearchRepository struct {
	db *sqlx.DB
}

// NewSearchRepository creates a new repository
func NewSearchRepository(db *sqlx.DB) *SearchRepository {
	return &SearchRepository{
		db: db,
	}
}

type dbSearch struct {
	model.Need
	OrganizationName string `db:"organization_name"`
	CategoryID       int64  `db:"category_id"`
	CategoryName     string `db:"category_name"`
	CategoryIcon     string `db:"category_icon"`
}

type searchOrganization struct {
	ID   int64
	Name string
}

// SearchResult formato do resultado da busca
type SearchResult struct {
	ID               int64
	Title            string
	Description      string
	RequiredQuantity int
	ReachedQuantity  int
	Unity            string
	DueDate          *time.Time
	Category         model.Category
	Organization     searchOrganization
}

// Search needs by text, category or organization
func (r *SearchRepository) Search(text string, categoriesID []int, organizationsID int64, page int64) ([]SearchResult, error) {
	var filter string
	var numParams int

	args := []interface{}{
		"%" + text + "%",
		(page - 1) * 10,
	}

	if organizationsID > 0 {
		filter += "and n.organization_id = $3"
		args = append(args, organizationsID)
		numParams++
	}

	if len(categoriesID) > 0 {
		var binds string
		for i := range categoriesID {
			binds += fmt.Sprintf("$%d,", i+numParams+3)
			args = append(args, categoriesID[i])
		}
		binds = binds[0 : len(binds)-1]
		filter = fmt.Sprintf("%s and n.category_id IN (%s) ", filter, binds)
	}

	sql := fmt.Sprintf(`
		SELECT n.*, o.name as organization_name, c.name as category_name, c.icon as category_icon
		FROM needs n
			INNER JOIN organizations o on (o.id = n.organization_id)
			INNER JOIN categories c on (c.id = n.category_id)
		WHERE (LOWER(n.title) LIKE $1 OR LOWER(n.description) LIKE $1)
			%s
		LIMIT 10 OFFSET $2
	`, filter)

	dbNeeds := []dbSearch{}
	err := r.db.Select(&dbNeeds, sql, args...)

	return convertDBToNeed(dbNeeds), err
}

func convertDBToNeed(dbSearch []dbSearch) []SearchResult {
	var need []SearchResult
	need = make([]SearchResult, len(dbSearch))
	for i, s := range dbSearch {
		need[i] = SearchResult{
			ID:               s.ID,
			Title:            s.Title,
			Description:      s.Description,
			RequiredQuantity: s.RequiredQuantity,
			ReachedQuantity:  s.ReachedQuantity,
			Unity:            s.Unity,
			DueDate:          s.DueDate,
			Category: model.Category{
				ID:   s.CategoryID,
				Name: s.CategoryName,
				Icon: s.CategoryIcon,
			},
			Organization: searchOrganization{
				ID:   s.OrganizationID,
				Name: s.OrganizationName,
			},
		}
	}

	return need
}

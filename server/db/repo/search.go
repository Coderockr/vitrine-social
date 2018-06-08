package repo

import (
	"fmt"

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

// Search needs by text, category or organization
func (r *SearchRepository) Search(text string, categoriesID []int, organizationsID int64, orderBy []string, page int64) ([]model.SearchNeed, error) {
	var filter string
	numParams := 3

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
			binds += fmt.Sprintf("$%d,", numParams)
			args = append(args, categoriesID[i])
			numParams++
		}
		binds = binds[0 : len(binds)-1]
		filter = fmt.Sprintf("%s and n.category_id IN (%s) ", filter, binds)
	}

	if len(orderBy) > 0 {
		switch orderBy[0] {
		case
			"id",
			"updated_at":
			break
		default:
			orderBy[0] = "created_at"
		}

		if len(orderBy) == 2 {
			if orderBy[1] != "asc" && orderBy[1] != "desc" {
				return nil, fmt.Errorf("Método de ordenação não reconhecido")
			}
		} else {
			orderBy = append(orderBy, "asc")
		}

		filter = fmt.Sprintf("%s ORDER BY %s %s ", filter, "n."+orderBy[0], orderBy[1])
	}

	sql := fmt.Sprintf(`
		SELECT n.*, o.name as organization_name, o.logo as organization_logo, o.slug as organization_slug,
					 c.name as category_name, c.icon as category_icon
		FROM needs n
			INNER JOIN organizations o on (o.id = n.organization_id)
			INNER JOIN categories c on (c.id = n.category_id)
		WHERE (LOWER(n.title) LIKE $1 OR LOWER(n.description) LIKE $1)
			%s
		LIMIT 10 OFFSET $2
	`, filter)

	dbNeeds := []model.SearchNeed{}
	err := r.db.Select(&dbNeeds, sql, args...)

	return dbNeeds, err
}

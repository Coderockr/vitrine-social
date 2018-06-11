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
func (r *SearchRepository) Search(text string, categoriesID []int, organizationsID int64, orderBy string, order string, page int64) ([]model.SearchNeed, error) {
	var filter string

	args := []interface{}{
		(page - 1) * 10,
	}

	if len(text) > 0 {
		filter += " and (LOWER(n.title) LIKE $2 OR LOWER(n.description) LIKE $2)"
		args = append(args, "%"+text+"%")
	}

	if organizationsID > 0 {
		filter += fmt.Sprintf(" and n.organization_id = $%d", len(args)+1)
		args = append(args, organizationsID)
	}

	if len(categoriesID) > 0 {
		var binds string
		for i := range categoriesID {
			binds += fmt.Sprintf("$%d,", len(args)+1)
			args = append(args, categoriesID[i])
		}
		binds = binds[0 : len(binds)-1]
		filter = fmt.Sprintf("%s and n.category_id IN (%s) ", filter, binds)
	}

	if len(orderBy) > 0 {
		switch orderBy {
		case
			"id",
			"updated_at":
			break
		default:
			orderBy = "created_at"
		}

		if len(order) > 0 {
			if order != "asc" && order != "desc" {
				return nil, fmt.Errorf("Método de ordenação não reconhecido")
			}
		} else {
			order = "asc"
		}

		filter = fmt.Sprintf("%s ORDER BY %s %s ", filter, "n."+orderBy, order)
	}

	sql := fmt.Sprintf(`
		SELECT n.*, o.name as organization_name, o.logo as organization_logo, o.slug as organization_slug,
					 c.name as category_name, c.slug as category_slug
		FROM needs n
			INNER JOIN organizations o on (o.id = n.organization_id)
			INNER JOIN categories c on (c.id = n.category_id)
		WHERE n.status = 'ACTIVE'
			%s
		LIMIT 10 OFFSET $1
	`, filter)

	dbNeeds := []model.SearchNeed{}
	err := r.db.Select(&dbNeeds, sql, args...)

	return dbNeeds, err
}

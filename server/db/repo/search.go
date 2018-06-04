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

// DBSearch estrutura de busca de necessidade
type DBSearch struct {
	model.Need
	OrganizationName string `db:"organization_name"`
	OrganizationLogo string `db:"organization_logo"`
	OrganizationSlug string `db:"organization_slug"`
	CategoryID       int64  `db:"category_id"`
	CategoryName     string `db:"category_name"`
	CategoryIcon     string `db:"category_icon"`
}

// Search needs by text, category or organization
func (r *SearchRepository) Search(text string, categoriesID []int, organizationsID int64, page int64) ([]DBSearch, error) {
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
		SELECT n.*, o.name as organization_name, o.logo as organization_logo, o.slug as organization_slug,
					 c.name as category_name, c.icon as category_icon
		FROM needs n
			INNER JOIN organizations o on (o.id = n.organization_id)
			INNER JOIN categories c on (c.id = n.category_id)
		WHERE (LOWER(n.title) LIKE $1 OR LOWER(n.description) LIKE $1)
			%s
		LIMIT 10 OFFSET $2
	`, filter)

	dbNeeds := []DBSearch{}
	err := r.db.Select(&dbNeeds, sql, args...)

	return dbNeeds, err
}

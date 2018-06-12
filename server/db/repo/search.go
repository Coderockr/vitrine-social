package repo

import (
	"fmt"
	"strings"

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
func (r *SearchRepository) Search(text string, categoriesID []int, organizationsID int64, status string, orderBy string, order string, page int) (needs []model.SearchNeed, count int, err error) {
	var filter string
	var args []interface{}

	if len(text) > 0 {
		filter += " and (LOWER(n.title) LIKE $2 OR LOWER(n.description) LIKE $2)"
		args = append(args, "%"+text+"%")
	}

	if len(status) > 0 {
		if status != "active" && status != "inactive" {
			return nil, 0, fmt.Errorf("O status informado é inválido")
		}

		filter += fmt.Sprintf(" and n.status = $%d", len(args)+1)
		args = append(args, strings.ToUpper(status))
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

	sqlCount := fmt.Sprintf(`
		SELECT count(n.*)
		FROM needs n
			INNER JOIN organizations o on (o.id = n.organization_id)
			INNER JOIN categories c on (c.id = n.category_id)
		WHERE n.id > 0
			%s
	`, filter)

	dbNeedsCount := 0
	err = r.db.Get(&dbNeedsCount, sqlCount, args...)

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
				return nil, 0, fmt.Errorf("Método de ordenação não reconhecido")
			}
		} else {
			order = "asc"
		}

		filter = fmt.Sprintf("%s ORDER BY %s %s ", filter, "n."+orderBy, order)
	}

	args = append(args, (page-1)*10)

	sql := fmt.Sprintf(`
		SELECT n.*, o.name as organization_name, o.logo as organization_logo, o.slug as organization_slug,
					 c.name as category_name, c.slug as category_slug
		FROM needs n
			INNER JOIN organizations o on (o.id = n.organization_id)
			INNER JOIN categories c on (c.id = n.category_id)
		WHERE n.id > 0
			%s
		LIMIT 10 OFFSET $%d
	`, filter, len(args))

	dbNeeds := []model.SearchNeed{}
	err = r.db.Select(&dbNeeds, sql, args...)

	dbNeeds, err = r.getNeedsImages(dbNeeds)

	return dbNeeds, dbNeedsCount, err
}

func (r *SearchRepository) getNeedsImages(dbNeeds []model.SearchNeed) ([]model.SearchNeed, error) {
	var filter string
	var binds string
	var args []interface{}

	for i, s := range dbNeeds {
		binds += fmt.Sprintf("$%d,", i+1)
		args = append(args, s.ID)
	}
	binds = binds[0 : len(binds)-1]
	filter = fmt.Sprintf("need_id IN (%s)", binds)

	sqlImages := fmt.Sprintf("SELECT * FROM needs_images WHERE %s", filter)

	dbNeedsImages := []model.NeedImage{}
	err := r.db.Select(&dbNeedsImages, sqlImages, args...)
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar imagens")
	}

	for i, need := range dbNeeds {
		for _, image := range dbNeedsImages {
			if image.NeedID == need.ID {
				dbNeeds[i].Images = append(dbNeeds[i].Images, image)
			}
		}
	}

	return dbNeeds, nil
}

package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

// NeedRepository is a implementation for Postgres
type NeedRepository struct {
	db      *sqlx.DB
	orgRepo *OrganizationRepository
	catRepo *CategoryRepository
}

// NewNeedRepository creates a new repository
func NewNeedRepository(db *sqlx.DB, orgRepo *OrganizationRepository) *NeedRepository {
	return &NeedRepository{
		db:      db,
		orgRepo: orgRepo,
		catRepo: NewCategoryRepository(db),
	}
}

// Get one Need from database
func (r *NeedRepository) Get(id int64) (*model.Need, error) {
	n := &model.Need{}
	err := r.db.Get(n, "SELECT * FROM needs WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	n.Category, err = r.catRepo.Get(n.CategoryID)
	return n, nil
}

// GetNeedImages without the need data
func (r *NeedRepository) getNeedImages(n *model.Need) ([]model.NeedImage, error) {
	images := []model.NeedImage{}
	err := r.db.Select(&images, "SELECT * FROM needs_images WHERE need_id = $1", n.ID)
	if err != nil {
		return nil, err
	}

	return images, nil
}

// Create creates a new need based on the struct
func (r *NeedRepository) Create(n model.Need) (model.Need, error) {
	n.Title = strings.TrimSpace(n.Title)
	if len(n.Title) == 0 {
		return n, errors.New("Deve ser informado um título para a Necessidade")
	}

	n.Description = strings.TrimSpace(n.Description)
	if len(n.Description) == 0 {
		return n, errors.New("Deve ser informada uma descrição para a Necessidade")
	}

	_, err := r.catRepo.Get(n.CategoryID)
	switch {
	case err == sql.ErrNoRows:
		return n, fmt.Errorf("Não foi encontrada categoria com ID: %d", n.CategoryID)
	case err != nil:
		return n, err
	}

	_, err = r.orgRepo.Get(n.OrganizationID)
	switch {
	case err == sql.ErrNoRows:
		return n, fmt.Errorf("Não foi encontrada Organização com ID: %d", n.OrganizationID)
	case err != nil:
		return n, err
	}

	n.Status = model.NeedStatusActive

	err = r.db.QueryRow(
		`INSERT INTO needs (category_id, organization_id, title, description, required_qtd, reached_qtd, due_date, status, unity)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id
		`,
		n.CategoryID,
		n.OrganizationID,
		n.Title,
		n.Description,
		n.RequiredQuantity,
		n.ReachedQuantity,
		n.DueDate,
		n.Status,
		n.Unity,
	).Scan(&n.ID)

	if err != nil {
		return n, err
	}

	return n, nil
}

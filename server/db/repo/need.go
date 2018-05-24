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
func NewNeedRepository(db *sqlx.DB) *NeedRepository {
	return &NeedRepository{
		db:      db,
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

// getNeedImages without the need data
func getNeedImages(db *sqlx.DB, n *model.Need) ([]model.NeedImage, error) {
	images := []model.NeedImage{}
	err := db.Select(&images, "SELECT * FROM needs_images WHERE need_id = $1", n.ID)
	if err != nil {
		return nil, err
	}

	return images, nil
}

// Create creates a new need based on the struct
func (r *NeedRepository) Create(n model.Need) (model.Need, error) {
	n, err := validate(r, n)

	if err != nil {
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

// Update - Receive a Need and update it in the database, returning the updated Need or error if failed
func (r *NeedRepository) Update(n model.Need) (model.Need, error) {
	n, err := validate(r, n)

	if err != nil {
		return n, err
	}

	_, err = r.db.Exec(
		`UPDATE needs SET
			category_id = $1,
			title = $2,
			description = $3,
			required_qtd = $4,
			reached_qtd = $5,
			due_date = $6,
			unity = $7
		WHERE id = $8
		`,
		n.CategoryID,
		n.Title,
		n.Description,
		n.RequiredQuantity,
		n.ReachedQuantity,
		n.DueDate,
		n.Unity,
		n.ID,
	)

	if err != nil {
		return n, err
	}

	return n, nil
}

func validate(r *NeedRepository, n model.Need) (model.Need, error) {
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

	_, err = getBaseOrganization(r.db, n.OrganizationID)
	switch {
	case err == sql.ErrNoRows:
		return n, fmt.Errorf("Não foi encontrada Organização com ID: %d", n.OrganizationID)
	case err != nil:
		return n, err
	}

	return n, nil
}

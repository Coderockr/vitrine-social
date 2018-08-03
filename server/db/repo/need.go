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
		orgRepo: NewOrganizationRepository(db),
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

	n.Images, _ = getNeedImages(r.db, n)

	c, _ := r.catRepo.Get(n.CategoryID)
	n.Category = *c

	o, err := r.orgRepo.GetBaseOrganization(n.OrganizationID)
	n.Organization = *o
	return n, nil
}

// GetNeedsImages retrive the images of a Need
func (r *NeedRepository) GetNeedsImages(n model.Need) ([]model.NeedImage, error) {
	return getNeedImages(r.db, &n)
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

	c, err := r.catRepo.Get(n.CategoryID)
	if err != nil {
		return n, err
	}

	if c == nil {
		return n, fmt.Errorf("category with id %d not found", n.CategoryID)
	}

	n.Category = *c
	n.Status = model.NeedStatusActive

	err = r.db.QueryRow(
		`INSERT INTO needs (category_id, organization_id, title, description, required_qtd, reached_qtd, due_date, status, unit, updated_at)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
		n.Unit,
		n.UpdatedAt,
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
			unit = $7,
			status = $8,
			updated_at = now()
		WHERE id = $9
		`,
		n.CategoryID,
		n.Title,
		n.Description,
		n.RequiredQuantity,
		n.ReachedQuantity,
		n.DueDate,
		n.Unit,
		n.Status,
		n.ID,
	)

	if err != nil {
		return n, err
	}

	return n, nil
}

// CreateImage creates a new need image based on the struct
func (r *NeedRepository) CreateImage(i model.NeedImage) (model.NeedImage, error) {
	err := r.db.QueryRow(
		`INSERT INTO needs_images (need_id, name, url)
			VALUES($1, $2, $3)
			RETURNING id
		`,
		i.NeedID,
		i.Image.Name,
		i.Image.URL,
	).Scan(&i.ID)

	if err != nil {
		return i, err
	}

	return i, nil
}

// DeleteImage delete a image from a need
func (r *NeedRepository) DeleteImage(imageID, needID int64) error {
	_, err := r.db.Exec(`DELETE FROM needs_images WHERE id = $1 AND need_id = $2`, imageID, needID)
	return err
}

func validate(r *NeedRepository, n model.Need) (model.Need, error) {
	n.Title = strings.TrimSpace(n.Title)
	if len(n.Title) == 0 {
		return n, errors.New("Deve ser informado um título para a Necessidade")
	}

	_, err := r.catRepo.Get(n.CategoryID)
	switch {
	case err == sql.ErrNoRows:
		return n, fmt.Errorf("Não foi encontrada categoria com ID: %d", n.CategoryID)
	case err != nil:
		return n, err
	}

	_, err = r.orgRepo.GetBaseOrganization(n.OrganizationID)
	switch {
	case err == sql.ErrNoRows:
		return n, fmt.Errorf("Não foi encontrada Organização com ID: %d", n.OrganizationID)
	case err != nil:
		return n, err
	}

	return n, nil
}

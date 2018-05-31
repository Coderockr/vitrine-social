package repo

import (
	"github.com/golang/sync/syncmap"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

// CategoryRepository to access database
type CategoryRepository struct {
	db *sqlx.DB
}

var categoryCache = syncmap.Map{}

// NewCategoryRepository create a new repository
func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

// Get a category from database using its id
func (r *CategoryRepository) Get(id int64) (*model.Category, error) {
	if c, ok := categoryCache.Load(id); ok {
		c := c.(*model.Category)
		return c, nil
	}

	c := &model.Category{}
	err := r.db.Get(c, "SELECT * FROM categories WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	categoryCache.Store(id, c)
	return c, nil
}

// GetAll return all categories
func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	c := []model.Category{}
	err := r.db.Select(&c, `
		SELECT categories.*, COUNT(needs.id) as count_need
		FROM categories
		LEFT JOIN needs ON needs.category_id = categories.id
			AND needs.status = 'ACTIVE'
		GROUP BY categories.id
		ORDER BY name ASC
	`)
	if err != nil {
		return []model.Category{}, err
	}

	return c, nil
}

package repo

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

// NeedRepository is a implementation for Postgres
type NeedRepository struct {
	db      *sqlx.DB
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

// GetNeedImages without the need data
func (r *NeedRepository) getNeedImages(n *model.Need) ([]model.NeedImage, error) {
	images := []model.NeedImage{}
	err := r.db.Select(&images, "SELECT * FROM needs_images WHERE need_id = $1", n.ID)
	if err != nil {
		return nil, err
	}

	return images, nil
}

//Insert insert a need
func (r *NeedRepository) Insert(n *model.Need) (int64, error) {
	query := `insert into need (title, description, required_qtd, reached_qtd, unity, due_date, category_id, organization_id) values (?,?,?,?,?,?,?,?)`
	res := r.db.MustExec(query, n.Title, n.Description, n.RequiredQuantity, n.ReachedQuantity, n.Unity, n.DueDate, n.CategoryID, n.OrganizationID)
	return res.LastInsertId()
}

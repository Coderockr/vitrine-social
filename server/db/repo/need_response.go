package repo

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

// NeedResponseRepository is a implementation for Postgres
type NeedResponseRepository struct {
	db *sqlx.DB
}

// NewNeedResponseRepository creates a new repository
func NewNeedResponseRepository(db *sqlx.DB) *NeedResponseRepository {
	return &NeedResponseRepository{db: db}
}

// CreateResponse create NeedResponse in database
func (r *NeedResponseRepository) CreateResponse(nr *model.NeedResponse) (int64, error) {
	var newID int64
	var query = `INSERT INTO need_response 
	(email, name, phone, address, message, need_id)
	 VALUES ($1, $2, $3, $4, $5, $6) returning id;`
	err := r.db.QueryRow(query, nr.Email, nr.Name, nr.Phone, nr.Address, nr.Message, nr.NeedID).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

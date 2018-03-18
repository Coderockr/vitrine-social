package repo

import (
	"database/sql"

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
func (r *NeedResponseRepository) CreateResponse(nr *model.NeedResponse) (sql.Result, error) {
	id, err := r.db.NamedExec(`INSERT INTO need_response 
		(email, name, phone, address, message, need_id)
		 VALUES (:email, :name, :phone, :address, :message, :need_id)`, nr)
	if err != nil {
		return nil, err
	}

	return id, nil
}

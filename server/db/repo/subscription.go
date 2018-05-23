package repo

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

// SubscriptionRepository is a implementation for Postgres
type SubscriptionRepository struct {
	db      *sqlx.DB
	orgRepo *OrganizationRepository
}

// NewSubscriptionRepository creates a new repository
func NewSubscriptionRepository(db *sqlx.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db:      db,
		orgRepo: NewOrganizationRepository(db),
	}
}

// Create new subscription
func (r *SubscriptionRepository) Create(s model.Subscription) (model.Subscription, error) {
	row := r.db.QueryRow(
		`INSERT INTO subscriptions (organization_id, name, email, phone)
			VALUES($1, $2, $3, $4)
			RETURNING id
		`,
		s.OrganizationID,
		s.Name,
		s.Email,
		s.Phone,
	)

	err := row.Scan(&s.ID)

	if err != nil {
		return s, err
	}

	return s, nil
}

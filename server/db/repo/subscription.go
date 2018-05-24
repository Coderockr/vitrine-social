package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	s, err := validate(r, s)

	if err != nil {
		return s, err
	}

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

	err = row.Scan(&s.ID)

	if err != nil {
		return s, err
	}

	return s, nil
}

func validate(r *SubscriptionRepository, s model.Subscription) (model.Subscription, error) {
	s.Name = strings.TrimSpace(s.Name)
	if len(s.Name) == 0 {
		return s, errors.New("Deve ser informado um nome para a Inscrição")
	}

	s.Email = strings.TrimSpace(s.Email)
	if len(s.Email) == 0 {
		return s, errors.New("Deve ser informado um email para a Inscrição")
	}

	s.Phone = strings.TrimSpace(s.Phone)
	if len(s.Phone) == 0 {
		return s, errors.New("Deve ser informado um telefone para a Inscrição")
	}

	_, err := getBaseOrganization(r.db, s.OrganizationID)
	switch {
	case err == sql.ErrNoRows:
		return s, fmt.Errorf("Não foi encontrada Organização com ID: %d", s.OrganizationID)
	case err != nil:
		return s, err
	}

	return s, nil
}

package repo

import (
	"fmt"

	"github.com/Coderockr/vitrine-social/server/security"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

// OrganizationRepository is a implementation for Postgres
type OrganizationRepository struct {
	db      *sqlx.DB
	catRepo *CategoryRepository
}

// NewOrganizationRepository creates a new repository
func NewOrganizationRepository(db *sqlx.DB) *OrganizationRepository {
	return &OrganizationRepository{
		db:      db,
		catRepo: NewCategoryRepository(db),
	}
}

const allFields string = `
	id, name, logo, phone, about, video, email, password, slug,
	street as "address.street", number as "address.number",
	complement as "address.complement", neighborhood as "address.neighborhood",
	city as "address.city", state as "address.state", zipcode as "address.zipcode"
`

// getBaseOrganization returns only the data about a organization, not its relations
func getBaseOrganization(db *sqlx.DB, id int64) (*model.Organization, error) {
	o := &model.Organization{}
	err := db.Get(o, "SELECT "+allFields+" FROM organizations WHERE id = $1", id)
	return o, err
}

// GetBaseOrganization returns only the data about a organization, not its relations
func (r *OrganizationRepository) GetBaseOrganization(id int64) (*model.Organization, error) {
	return getBaseOrganization(r.db, id)
}

// Get one Organization from database
func (r *OrganizationRepository) Get(id int64) (*model.Organization, error) {
	o, err := getBaseOrganization(r.db, id)
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&o.Images, "SELECT * FROM organizations_images WHERE organization_id = $1", id)
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&o.Needs, "SELECT * FROM needs WHERE organization_id = $1", id)
	if err != nil {
		return nil, err
	}

	for i := range o.Needs {
		o.Needs[i].Category, err = r.catRepo.Get(o.Needs[i].CategoryID)
		if err != nil {
			fmt.Println("test?")
			return nil, err
		}

		o.Needs[i].Images, err = getNeedImages(r.db, &o.Needs[i])
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}

// Create receives a Organization and creates it in the database, returning the updated Organization or error if failed
func (r *OrganizationRepository) Create(o model.Organization) (model.Organization, error) {
	row := r.db.QueryRow(
		`INSERT INTO organizations (
			name, logo, phone, about, video, email, slug, password,
			street, number, complement, neighborhood, city, state, zipcode
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
		)
		RETURNING id
		`,
		o.Name,
		o.Logo,
		o.Phone,
		o.About,
		o.Video,
		o.Email,
		o.Slug,
		o.Password,
		o.Address.Street,
		o.Address.Number,
		o.Address.Complement,
		o.Address.Neighborhood,
		o.Address.City,
		o.Address.State,
		o.Address.Zipcode,
	)

	err := row.Scan(&o.ID)

	if err != nil {
		return o, err
	}

	return o, nil
}

// Update - Receive an Organization and update it in the database, returning the updated Organization or error if failed
func (r *OrganizationRepository) Update(o model.Organization) (model.Organization, error) {
	_, err := r.db.Exec(
		`UPDATE organizations SET
			name = $1,
			logo = $2,
			phone = $3,
			about = $4,
			video = $5,
			email = $6,
			street = $7,
			number = $8,
			complement = $9,
			neighborhood = $10,
			city = $11,
			state = $12,
			zipcode = $13
		WHERE id = $14
		`,
		o.Name,
		o.Logo,
		o.Phone,
		o.About,
		o.Video,
		o.Email,
		o.Address.Street,
		o.Address.Number,
		o.Address.Complement,
		o.Address.Neighborhood,
		o.Address.City,
		o.Address.State,
		o.Address.Zipcode,
		o.ID,
	)

	if err != nil {
		return o, err
	}

	return o, nil
}

// DeleteImage - Receive an id and remove the image
func (r *OrganizationRepository) DeleteImage(imageID int64, organizationID int64) error {
	_, err := r.db.Exec(`DELETE FROM organizations_images WHERE id = $1 AND organization_id = $2`, imageID, organizationID)
	return err
}

// GetByEmail returns a organization by its email
func (r *OrganizationRepository) GetByEmail(email string) (*model.Organization, error) {
	o := model.Organization{}
	err := r.db.Get(&o, `SELECT `+allFields+` FROM organizations WHERE email = $1`, email)
	return &o, err
}

// GetUserByEmail returns a organization user by its email
func (r *OrganizationRepository) GetUserByEmail(email string) (model.User, error) {
	o, err := r.GetByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	return o.User, nil
}

// ResetPasswordTo resets the organization password to the value informed
func (r *OrganizationRepository) ResetPasswordTo(o *model.Organization, password string) error {
	hash, err := security.GenerateHash(password)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`UPDATE organizations SET password = $1 WHERE id = $2`, hash, o.ID)
	if err != nil {
		return err
	}
	o.Password = hash
	return nil
}

// CreateImage creates a new organization image based on the struct
func (r *OrganizationRepository) CreateImage(i model.OrganizationImage) (model.OrganizationImage, error) {
	err := r.db.QueryRow(
		`INSERT INTO organizations_images (organization_id, name, url)
			VALUES($1, $2, $3)
			RETURNING id
		`,
		i.OrganizationID,
		i.Image.Name,
		i.Image.URL,
	).Scan(&i.ID)

	if err != nil {
		return i, err
	}

	return i, nil
}

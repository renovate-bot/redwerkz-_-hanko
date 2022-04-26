package persistence

import (
	"database/sql"
	"fmt"
	"github.com/gobuffalo/pop/v6"
	"github.com/teamhanko/hanko/persistence/models"
)

type JwkPersister struct {
	db *pop.Connection
}

func NewJwkPersister(db *pop.Connection) *JwkPersister {
	return &JwkPersister{db: db}
}

func (p *JwkPersister) Get(id int) (*models.Jwk, error) {
	jwk := models.Jwk{}
	err := p.db.Find(&jwk, id)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get jwk: %w", err)
	}
	return &jwk, nil
}

func (p *JwkPersister) GetAll() ([]models.Jwk, error) {
	jwks := []models.Jwk{}
	err := p.db.All(&jwks)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get all jwks: %w", err)
	}
	return jwks, nil
}

func (p *JwkPersister) GetLast() (*models.Jwk, error) {
	jwk := models.Jwk{}
	err := p.db.Order("id asc").Last(&jwk)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get jwk: %w", err)
	}
	return &jwk, nil
}

func (p *JwkPersister) Create(jwk models.Jwk) error {
	vErr, err := p.db.ValidateAndCreate(&jwk)
	if err != nil {
		return fmt.Errorf("failed to store jwk: %w", err)
	}

	if vErr != nil && vErr.HasAny() {
		return fmt.Errorf("jwk object validation failed: %w", vErr)
	}

	return nil
}

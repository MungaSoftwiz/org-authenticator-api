package org

import (
	"github.com/MungaSoftwiz/org-authenticator-api/types"
	"github.com/jmoiron/sqlx"
)

type OrganisationStorage struct {
	db *sqlx.DB
}

func NewOrganisationStorage(db *sqlx.DB) *OrganisationStorage {
	return &OrganisationStorage{db: db}
}

func (s *OrganisationStorage) GetAllOrganisations() ([]types.Organisation, error) {
	var organisations []types.Organisation
	err := s.db.Select(&organisations, "SELECT * FROM organisations")
	return organisations, err
}

func (s *OrganisationStorage) GetOrganisationByID(orgID string) (*types.Organisation, error) {
	var organisation types.Organisation
	err := s.db.Get(&organisation, "SELECT * FROM organisations WHERE orgId = $1", orgID)
	return &organisation, err
}

func (s *OrganisationStorage) CreateOrganisation(organisation types.Organisation) error {
	_, err := s.db.NamedExec(`INSERT INTO organisations (orgId, name, description) VALUES (:org_id, :name, :description)`, organisation)
	return err
}

func (s *OrganisationStorage) AddUserToOrganisation(orgID string, userID string) error {
	_, err := s.db.Exec(`INSERT INTO organisation_users (orgId, userId) VALUES ($1, $2)`, orgID, userID)
	return err
}

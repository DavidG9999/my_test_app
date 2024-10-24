package postgres

import (
	"fmt"
	"strings"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type OrganizationPostgres struct {
	db *sqlx.DB
}

func NewOrganizationPostgres(db *sqlx.DB) *OrganizationPostgres {
	return &OrganizationPostgres{
		db: db,
	}
}

func (r *OrganizationPostgres) CreateOrganization(organization entity.Organization) (entity.Organization, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, address, chief, financial_chief, inn_kpp) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, address, chief, financial_chief, inn_kpp", organizationsTable)
	row := r.db.QueryRow(query, organization.Name, organization.Address, organization.Chief, organization.FinancialChief, organization.InnKpp)
	if err := row.Scan(&organization.Id, &organization.Name, &organization.Address, &organization.Chief, &organization.FinancialChief, &organization.InnKpp); err != nil {
		return entity.Organization{}, err
	}
	return organization, nil
}

func (r *OrganizationPostgres) GetOrganizations() ([]entity.Organization, error) {
	var organizations []entity.Organization
	query := fmt.Sprintf("SELECT * FROM %s", organizationsTable)
	err := r.db.Select(&organizations, query)
	return organizations, err
}

func (r *OrganizationPostgres) GetOrganizationById(organizationId int) (entity.Organization, error) {
	var organization entity.Organization
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", organizationsTable)
	err := r.db.Get(&organization, query, organizationId)
	return organization, err
}

func (r *OrganizationPostgres) GetOrganizationByInnKpp(innKpp string) (entity.Organization, error) {
	var organization entity.Organization
	query := fmt.Sprintf("SELECT * FROM %s WHERE inn_kpp=$1", organizationsTable)
	err := r.db.Get(&organization, query, innKpp)
	return organization, err
}

func (r *OrganizationPostgres) UpdateOrganization(organizationId int, updateData entity.UpdateOrganizationInput) (entity.Organization, error) {
	var updateOrganization entity.Organization

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, updateData.Name)
		argId++
	}
	if updateData.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, updateData.Address)
		argId++
	}
	if updateData.Chief != nil {
		setValues = append(setValues, fmt.Sprintf("chief=$%d", argId))
		args = append(args, updateData.Chief)
		argId++
	}
	if updateData.FinancialChief != nil {
		setValues = append(setValues, fmt.Sprintf("financial_chief=$%d", argId))
		args = append(args, updateData.FinancialChief)
		argId++
	}
	if updateData.InnKpp != nil {
		setValues = append(setValues, fmt.Sprintf("inn_kpp=$%d", argId))
		args = append(args, updateData.InnKpp)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id, name, address, chief, financial_chief, inn_kpp", organizationsTable, setQuery, argId)
	args = append(args, organizationId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	err := r.db.Get(&updateOrganization, query, args...)
	return updateOrganization, err
}

func (r *OrganizationPostgres) DeleteOrganization(organizationId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", organizationsTable)
	_, err := r.db.Exec(query, organizationId)
	return err
}

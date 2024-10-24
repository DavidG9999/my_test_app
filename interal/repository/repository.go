package repository

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUSer(user entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
}

type UserRepository interface {
}

type AutoRepository interface {
	CreateAuto(auto entity.Auto) (entity.Auto, error)
	GetAutos() ([]entity.Auto, error)
	GetAutoByStateNumber(stateNumber string) (entity.Auto, error)
	UpdateAuto(accountId int, updareData entity.UpdateAutoInput) (entity.Auto, error)
	DeleteAuto(autoId int) error
}

type ContragentRepository interface {
}

type DispetcherRepository interface {
}

type DriverRepository interface {
}

type MechanicRepository interface {
}

type OrganizationRepository interface {
	CreateOrganization(organization entity.Organization) (entity.Organization, error)
	GetOrganizations() ([]entity.Organization, error)
	GetOrganizationById(organizationId int) (entity.Organization, error)
	GetOrganizationByInnKpp(innKpp string) (entity.Organization, error)
	UpdateOrganization(organizationId int, updateData entity.UpdateOrganizationInput) (entity.Organization, error)
	DeleteOrganization(organizationId int) error
}

type AccountRepository interface {
	CreateAccount(account entity.Account) (entity.Account, error)
	GetAccounts() ([]entity.Account, error)
	GetAccountByAccountNumber(accountNymber string) (entity.Account, error)
	UpdateAccount(autoId int, updateData entity.UpdateAccountInput) (entity.Account, error)
	DeleteAccount(accountId int) error
}

type PutlistRepository interface {
}

type Repository struct {
	Authorization
	UserRepository
	AccountRepository
	AutoRepository
	ContragentRepository
	DispetcherRepository
	DriverRepository
	MechanicRepository
	OrganizationRepository
	PutlistRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:          postgres.NewAuthPostgres(db),
		AutoRepository:         postgres.NewAutoPostgres(db),
		AccountRepository:      postgres.NewAccountPostgres(db),
		OrganizationRepository: postgres.NewOrganizationPostgres(db),
	}
}

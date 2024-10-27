package repository

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUSer(user entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
	GetUserById(id int) (entity.User, error)
	UpdateName(id int, updateData entity.UpdateNameUserInput) (entity.User, error)
	UpdatePassword(id int, updateData entity.UpdatePasswordUserInput) (entity.User, error)
	DeleteUser(id int) error
}

type AutoRepository interface {
	CreateAuto(auto entity.Auto) (entity.Auto, error)
	GetAutos() ([]entity.Auto, error)
	UpdateAuto(accountId int, updareData entity.UpdateAutoInput) (entity.Auto, error)
	DeleteAuto(autoId int) error
}

type ContragentRepository interface {
	CreateContragent(contragent entity.Contragent) (entity.Contragent, error)
	GetContragents() ([]entity.Contragent, error)
	UpdateContragent(contragentId int, updateData entity.UpdateContragentInput) (entity.Contragent, error)
	DeleteContragent(contragentId int) error
}

type DispetcherRepository interface {
	CreateDispetcher(dispetcher entity.Dispetcher) (entity.Dispetcher, error)
	GetDispetchers() ([]entity.Dispetcher, error)
	UpdateDispetcher(dispetcherId int, updateData entity.UpdateDispetcherInput) (entity.Dispetcher, error)
	DeleteDispetcher(dispetcherId int) error
}

type DriverRepository interface {
	CreateDriver(driver entity.Driver) (entity.Driver, error)
	GetDrivers() ([]entity.Driver, error)
	UpdateDriver(driverId int, updateData entity.UpdateDriverInput) (entity.Driver, error)
	DeleteDriver(driverId int) error
}

type MechanicRepository interface {
	CreateMechanic(mechanic entity.Mechanic) (entity.Mechanic, error)
	GetMechanics() ([]entity.Mechanic, error)
	UpdateMechanic(mechanicId int, updateData entity.UpdateMechanicInput) (entity.Mechanic, error)
	DeleteMechanic(mechanicId int) error
}

type OrganizationRepository interface {
	CreateOrganization(organization entity.Organization) (entity.Organization, error)
	GetOrganizations() ([]entity.Organization, error)
	UpdateOrganization(organizationId int, updateData entity.UpdateOrganizationInput) (entity.Organization, error)
	DeleteOrganization(organizationId int) error
}

type AccountRepository interface {
	CreateAccount(organizationId int, account entity.Account) (entity.Account, error)
	GetAccounts(organizationId int) (entity.Organization, []entity.Account, error)
	UpdateAccount(accountId int, updateData entity.UpdateAccountInput) (entity.Account, error)
	DeleteAccount(accountId int) error
}

type PutlistRepository interface {
	CreatePutlist(userId int, putlist entity.PutlistHeader) (entity.PutlistHeader, error)
	GetPutlists(userId int) ([]entity.GetPutlistResponse, error)
	GetPutlistByNumber(userId int, putlistNumber int) (entity.GetPutlistResponse, []entity.GetPutlistBodyResponse, error)
	UpdatePutlist(userId int, putlistNumber int, updateData entity.UpdatePutlistHeaderInput) (entity.PutlistHeader, error)
	DeletePutlist(userId int, putlistNumber int) error
}

type PutlistBodyRepository interface {
	CreatePutlistBody(putlistId int, putlistBody entity.PutlistBody) (entity.PutlistBody, error)
	GetPutlistBodies(putlistId int) ([]entity.GetPutlistBodyResponse, error)
	UpdatePutlistBody(putlistBodyId int, updateData entity.UpdatePutlistBodyInput) (entity.PutlistBody, error)
	DeletePutlistBody(putlistBodyId int) error
}

type Repository struct {
	Authorization
	AccountRepository
	AutoRepository
	ContragentRepository
	DispetcherRepository
	DriverRepository
	MechanicRepository
	OrganizationRepository
	PutlistRepository
	PutlistBodyRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:          postgres.NewAuthPostgres(db),
		AutoRepository:         postgres.NewAutoPostgres(db),
		OrganizationRepository: postgres.NewOrganizationPostgres(db),
		AccountRepository:      postgres.NewAccountPostgres(db),
		ContragentRepository:   postgres.NewContragentPostgres(db),
		DispetcherRepository:   postgres.NewDispetcherPostgres(db),
		DriverRepository:       postgres.NewDriverPostgres(db),
		MechanicRepository:     postgres.NewMechanicPostgres(db),
		PutlistRepository:      postgres.NewPutlistPostgres(db),
		PutlistBodyRepository:  postgres.NewPutlistBodyPostgres(db),
	}
}

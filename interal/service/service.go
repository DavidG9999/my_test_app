package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
	GetUserById(id int) (entity.User, error)
	UpdateName(id int, updateData entity.UpdateNameUserInput) (entity.User, error)
	UpdatePassword(id int, updateData entity.UpdatePasswordUserInput) (entity.User, error)
	DeleteUser(id int) error
}

type Auto interface {
	CreateAuto(auto entity.Auto) (entity.Auto, error)
	GetAutos() ([]entity.Auto, error)
	UpdateAuto(autoId int, updateData entity.UpdateAutoInput) (entity.Auto, error)
	DeleteAuto(autoId int) error
}

type Contragent interface {
	CreateContragent(contragent entity.Contragent) (entity.Contragent, error)
	GetContragents() ([]entity.Contragent, error)
	UpdateContragent(contragentId int, updateData entity.UpdateContragentInput) (entity.Contragent, error)
	DeleteContragent(contragentId int) error
}

type Dispetcher interface {
	CreateDispetcher(dispetcher entity.Dispetcher) (entity.Dispetcher, error)
	GetDispetchers() ([]entity.Dispetcher, error)
	UpdateDispetcher(dispetcherId int, updateData entity.UpdateDispetcherInput) (entity.Dispetcher, error)
	DeleteDispetcher(dispetcherId int) error
}

type Driver interface {
	CreateDriver(driver entity.Driver) (entity.Driver, error)
	GetDrivers() ([]entity.Driver, error)
	UpdateDriver(driverId int, updateData entity.UpdateDriverInput) (entity.Driver, error)
	DeleteDriver(driverId int) error
}

type Mechanic interface {
	CreateMechanic(mechanic entity.Mechanic) (entity.Mechanic, error)
	GetMechanics() ([]entity.Mechanic, error)
	UpdateMechanic(mechanicId int, updateData entity.UpdateMechanicInput) (entity.Mechanic, error)
	DeleteMechanic(mechanicId int) error
}

type Organization interface {
	CreateOrganization(organization entity.Organization) (entity.Organization, error)
	GetOrganizations() ([]entity.Organization, error)
	UpdateOrganization(organizationId int, updateData entity.UpdateOrganizationInput) (entity.Organization, error)
	DeleteOrganization(organizationId int) error
}

type Account interface {
	CreateAccount(organizationId int, account entity.Account) (entity.Account, error)
	GetAccounts(organizationId int) (entity.Organization, []entity.Account, error)
	UpdateAccount(accountId int, updateData entity.UpdateAccountInput) (entity.Account, error)
	DeleteAccount(accountId int) error
}
type Putlist interface {
	CreatePutlist(userId int, putlist entity.PutlistHeader) (entity.PutlistHeader, error)
	GetPutlistHeaders(userId int) ([]entity.GetPutlistResponse, error)
	GetPutlistByNumber(userId int, putlistNumber int) (entity.GetPutlistResponse, []entity.GetPutlistBodyResponse, error)
	UpdatePutlist(userId int, putlistNumber int, updateData entity.UpdatePutlistHeaderInput) (entity.PutlistHeader, error)
	DeletePutlist(userId int, putlistNumber int) error
	CreatePutlistBody(putlistNumber int, putlistBody entity.PutlistBody) (entity.PutlistBody, error)
	GetPutlistBodies(putlistNumber int) ([]entity.GetPutlistBodyResponse, error)
	UpdatePutlistBody(putlistBodyId int, updateData entity.UpdatePutlistBodyInput) (entity.PutlistBody, error)
	DeletePutlistBody(putlistBodyId int) error
}

type Service struct {
	Authorization
	Auto
	Contragent
	Dispetcher
	Driver
	Mechanic
	Organization
	Account
	Putlist
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Auto:          NewAutoService(repos.AutoRepository),
		Organization:  NewOrganizationService(repos.OrganizationRepository),
		Account:       NewAccountService(repos.AccountRepository),
		Contragent:    NewContragentService(repos.ContragentRepository),
		Dispetcher:    NewDispetcherService(repos.DispetcherRepository),
		Driver:        NewDriverService(repos.DriverRepository),
		Mechanic:      NewMechanicService(repos.MechanicRepository),
		Putlist:       NewPutlistService(repos.PutlistRepository, repos.PutlistBodyRepository),
	}
}

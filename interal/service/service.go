package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type User interface {
}

type Account interface {
	CreateAccount(account entity.Account) (entity.Account, error)
	GetAccounts() ([]entity.Account, error)
	GetAccountByAccountNumber(accountNymber string) (entity.Account, error)
	UpdateAccount(accountId int, updateData entity.UpdateAccountInput) (entity.Account, error)
	DeleteAccount(accountId int) error
}

type Auto interface {
	CreateAuto(auto entity.Auto) (entity.Auto, error)
	GetAutos() ([]entity.Auto, error)
	GetAutoByStateNumber(stateNumber string) (entity.Auto, error)
	UpdateAuto(autoId int, updateData entity.UpdateAutoInput) (entity.Auto, error)
	DeleteAuto(autoId int) error
}

type Contragent interface {
}

type Dispetcher interface {
}

type Driver interface {
}

type Mechanic interface {
}

type Organization interface {
	CreateOrganization(organization entity.Organization) (entity.Organization, error)
	GetOrganizations() ([]entity.Organization, error)
	GetOrganizationById(organizationId int) (entity.Organization, error)
	GetOrganizationByInnKpp(innKpp string) (entity.Organization, error)
	UpdateOrganization(organizationId int, updateData entity.UpdateOrganizationInput) (entity.Organization, error)
	DeleteOrganization(organizationId int) error
}

type Putlist interface {
}

type Service struct {
	Authorization
	User
	Account
	Auto
	Contragent
	Dispetcher
	Driver
	Mechanic
	Organization
	Putlist
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Auto:          NewAutoService(repos.AutoRepository),
		Account:       NewAccountService(repos.AccountRepository),
		Organization: NewOrganizationService(repos.OrganizationRepository),
	}
}

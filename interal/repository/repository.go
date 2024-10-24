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
	UpdateAuto(autoId int, updareData entity.UpdateAutoInput) (entity.Auto, error)
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
}

type PutlistRepository interface {
}

type Repository struct {
	Authorization
	UserRepository
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
		Authorization:  postgres.NewAuthPostgres(db),
		AutoRepository: postgres.NewAutoPostgres(db),
	}
}

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
}

type Putlist interface {
}

type Service struct {
	Authorization
	User
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
	}
}

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

type UserService interface{
	
}

type AutoService interface {
}

type ContragentService interface {
}

type DispetcherService interface {
}

type DriverService interface {
}

type MechanicService interface {
}

type Organizationservice interface {
}

type PutlistService interface {
}

type Service struct {
	Authorization
	UserService
	AutoService
	ContragentService
	DispetcherService
	DriverService
	MechanicService
	Organizationservice
	PutlistService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

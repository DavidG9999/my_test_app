package service

import "github.com/DavidG9999/my_test_app/interal/repository"

type Authorization interface {
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
	AutoService
	ContragentService
	DispetcherService
	DriverService
	MechanicService
	Organizationservice
	PutlistService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}

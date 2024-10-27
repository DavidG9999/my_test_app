package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type DriverService struct {
	repos repository.DriverRepository
}

func NewDriverService(repos repository.DriverRepository) *DriverService {
	return &DriverService{
		repos: repos,
	}
}

func (s *DriverService) CreateDriver(driver entity.Driver) (entity.Driver, error) {
	return s.repos.CreateDriver(driver)
}

func (s *DriverService) GetDrivers() ([]entity.Driver, error) {
	return s.repos.GetDrivers()
}

func (s *DriverService) UpdateDriver(driverId int, updateData entity.UpdateDriverInput) (entity.Driver, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.Driver{}, err
	}
	return s.repos.UpdateDriver(driverId, updateData)
}

func (s *DriverService) DeleteDriver(driverId int) error {
	return s.repos.DeleteDriver(driverId)
}

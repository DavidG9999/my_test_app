package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type DispetcherService struct {
	repos repository.DispetcherRepository
}

func NewDispetcherService(repos repository.DispetcherRepository) *DispetcherService {
	return &DispetcherService{
		repos: repos,
	}
}

func (s *DispetcherService) CreateDispetcher(dispetcher entity.Dispetcher) (entity.Dispetcher, error) {
	return s.repos.CreateDispetcher(dispetcher)
}

func (s *DispetcherService) GetDispetchers() ([]entity.Dispetcher, error) {
	return s.repos.GetDispetchers()
}

func (s *DispetcherService) UpdateDispetcher(dispetcherId int, updateData entity.UpdateDispetcherInput) (entity.Dispetcher, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.Dispetcher{}, err
	}
	return s.repos.UpdateDispetcher(dispetcherId, updateData)
}

func (s *DispetcherService) DeleteDispetcher(dispetcherId int) error {
	return s.repos.DeleteDispetcher(dispetcherId)
}

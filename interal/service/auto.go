package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type AutoService struct {
	repo repository.AutoRepository
}

func NewAutoService(repo repository.AutoRepository) *AutoService {
	return &AutoService{
		repo: repo,
	}
}

func (s *AutoService) CreateAuto(auto entity.Auto) (entity.Auto, error) {
	return s.repo.CreateAuto(auto)
}

func (s *AutoService) GetAutos() ([]entity.Auto, error) {
	return s.repo.GetAutos()
}

func (s *AutoService) UpdateAuto(autoId int, updateData entity.UpdateAutoInput) (entity.Auto, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.Auto{}, err
	}

	return s.repo.UpdateAuto(autoId, updateData)
}

func (s *AutoService) DeleteAuto(autoId int) error {
	return s.repo.DeleteAuto(autoId)
}

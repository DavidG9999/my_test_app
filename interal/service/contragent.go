package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type ContragentService struct {
	repo repository.ContragentRepository
}

func NewContragentService(repo repository.ContragentRepository) *ContragentService {
	return &ContragentService{
		repo: repo,
	}
}

func (s *ContragentService) CreateContragent(contragent entity.Contragent) (entity.Contragent, error) {
	return s.repo.CreateContragent(contragent)
}

func (s *ContragentService) GetContragents() ([]entity.Contragent, error) {
	return s.repo.GetContragents()
}

func (s *ContragentService) UpdateContragent(contragentId int, updateData entity.UpdateContragentInput) (entity.Contragent, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.Contragent{}, err
	}
	return s.repo.UpdateContragent(contragentId, updateData)
}

func (s *ContragentService) DeleteContragent(contragentId int) error {
	return s.repo.DeleteContragent(contragentId)

}

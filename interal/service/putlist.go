package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type PutlistService struct {
	repo     repository.PutlistRepository
	repoBody repository.PutlistBodyRepository
}

func NewPutlistService(repo repository.PutlistRepository, repoBody repository.PutlistBodyRepository) *PutlistService {
	return &PutlistService{
		repo:     repo,
		repoBody: repoBody,
	}
}

func (s *PutlistService) CreatePutlist(userId int, putlist entity.PutlistHeader) (entity.PutlistHeader, error) {
	return s.repo.CreatePutlist(userId, putlist)
}

func (s *PutlistService) GetPutlistHeaders(userId int) ([]entity.GetPutlistResponse, error) {
	return s.repo.GetPutlistHeaders(userId)
}

func (s *PutlistService) GetPutlistByNumber(userId int, putlistNumber int) (entity.GetPutlistResponse, []entity.GetPutlistBodyResponse, error) {
	return s.repo.GetPutlistByNumber(userId, putlistNumber)
}

func (s *PutlistService) UpdatePutlist(userId int, putlistNumber int, updateData entity.UpdatePutlistHeaderInput) (entity.PutlistHeader, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.PutlistHeader{}, err
	}

	return s.repo.UpdatePutlist(userId, putlistNumber, updateData)
}

func (s *PutlistService) DeletePutlist(userId int, putlistNumber int) error {
	return s.repo.DeletePutlist(userId, putlistNumber)
}

func (s *PutlistService) CreatePutlistBody(putlistNumber int, putlistBody entity.PutlistBody) (entity.PutlistBody, error) {
	return s.repoBody.CreatePutlistBody(putlistNumber, putlistBody)
}

func (s *PutlistService) GetPutlistBodies(putlistNumber int) ([]entity.GetPutlistBodyResponse, error) {
	return s.repoBody.GetPutlistBodies(putlistNumber)
}

func (s *PutlistService) UpdatePutlistBody(putlistBodyId int, updateData entity.UpdatePutlistBodyInput) (entity.PutlistBody, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.PutlistBody{}, err
	}

	return s.repoBody.UpdatePutlistBody(putlistBodyId, updateData)

}

func (s *PutlistService) DeletePutlistBody(putlistBodyId int) error {
	return s.repoBody.DeletePutlistBody(putlistBodyId)
}

package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type MechanicService struct {
	repos repository.MechanicRepository
}

func NewMechanicService(repos repository.MechanicRepository) *MechanicService {
	return &MechanicService{
		repos: repos,
	}
}

func (s *MechanicService) CreateMechanic(Mechanic entity.Mechanic) (entity.Mechanic, error) {
	return s.repos.CreateMechanic(Mechanic)
}

func (s *MechanicService) GetMechanics() ([]entity.Mechanic, error) {
	return s.repos.GetMechanics()
}

func (s *MechanicService) UpdateMechanic(MechanicId int, updateData entity.UpdateMechanicInput) (entity.Mechanic, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.Mechanic{}, err
	}
	return s.repos.UpdateMechanic(MechanicId, updateData)
}

func (s *MechanicService) DeleteMechanic(MechanicId int) error {
	return s.repos.DeleteMechanic(MechanicId)
}

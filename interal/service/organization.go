package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type OrganizationService struct {
	repos repository.OrganizationRepository
}

func NewOrganizationService(repos repository.OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		repos: repos,
	}
}

func (s *OrganizationService) CreateOrganization(organization entity.Organization) (entity.Organization, error) {
	return s.repos.CreateOrganization(organization)
}

func (s *OrganizationService) GetOrganizations() ([]entity.Organization, error) {
	return s.repos.GetOrganizations()
}

func (s *OrganizationService) UpdateOrganization(organizationId int, updateData entity.UpdateOrganizationInput) (entity.Organization, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.Organization{}, err
	}
	return s.repos.UpdateOrganization(organizationId, updateData)
}

func (s *OrganizationService) DeleteOrganization(organizationId int) error {
	return s.repos.DeleteOrganization(organizationId)
}

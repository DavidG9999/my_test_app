package service

import (
	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/DavidG9999/my_test_app/interal/repository"
)

type AccountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) CreateAccount(organizationId int, account entity.Account) (entity.Account, error) {
	return s.repo.CreateAccount(organizationId, account)
}

func (s *AccountService) GetAccounts(organizationId int) (entity.Organization, []entity.Account, error) {
	return s.repo.GetAccounts(organizationId)
}

func (s *AccountService) UpdateAccount(accountId int, updateData entity.UpdateAccountInput) (entity.Account, error) {
	err := updateData.Validate()
	if err != nil {
		return entity.Account{}, err
	}
	return s.repo.UpdateAccount(accountId, updateData)
}

func (s *AccountService) DeleteAccount(accountId int) error {
	return s.repo.DeleteAccount(accountId)
}

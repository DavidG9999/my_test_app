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

func (s *AccountService) CreateAccount(account entity.Account) (entity.Account, error) {
	return s.repo.CreateAccount(account)
}

func (s *AccountService) GetAccounts() ([]entity.Account, error) {
	return s.repo.GetAccounts()
}

func (s *AccountService) GetAccountByAccountNumber(accountNumber string) (entity.Account, error) {
	return s.repo.GetAccountByAccountNumber(accountNumber)
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

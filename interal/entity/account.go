package entity

import "errors"

type Account struct {
	Id             int    `json:"id" db:"id"`
	AccountNumber  string `json:"account_number" binding:"required,min=20,max=20" db:"account_number"`
	BankName       string `json:"bank_name" binding:"required" db:"bank_name"`
	BankIdNumber   string `json:"bank_id_number" binding:"required,min=9,max=9" db:"bank_id_number"`
	OrganizationId int    `json:"organization_id" db:"organization_id"`
}

type UpdateAccountInput struct {
	AccountNumber *string `json:"account_number"`
	BankName      *string `json:"bank_name"`
	BankIdNumber  *string `json:"bank_id_number"`
}

func (i UpdateAccountInput) Validate() error {
	if i.AccountNumber == nil && i.BankName == nil && i.BankIdNumber == nil {
		return errors.New("update structure has no values")
	}
	if i.AccountNumber != nil && len(*i.AccountNumber) != 20 {
		return errors.New("invalid field format: account_number")
	}
	if i.BankIdNumber != nil && len(*i.BankIdNumber) != 9 {
		return errors.New("invalid field format:  bank_id_number")
	}
	if i.BankName != nil {
		if *i.BankName == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}

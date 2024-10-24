package entity

import "errors"

type Account struct {
	Id            int    `json:"-" db:"id"`
	AccountNumber string `json:"account_number" binding:"required,eq=20" db:"account_number"`
	BankName      string `json:"bank_name" binding:"required" db:"bank_name"`
	BankIdNumber  string `json:"bank_id_number" binding:"required,eq=9" db:"bank_id_number"`
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
	if i.AccountNumber != nil && len(*i.AccountNumber) == 20 {
		return errors.New("invalid field format: account_number")
	}
	if i.BankIdNumber != nil && len(*i.BankIdNumber) == 9 {
		return errors.New("invalid field format:  bank_id_number")
	}
	return nil
}

package postgres

import (
	"fmt"
	"strings"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{
		db: db,
	}
}

func (r *AccountPostgres) CreateAccount(account entity.Account) (entity.Account, error) {
	query := fmt.Sprintf("INSERT INTO %s (account_number, bank_name, bank_id_number) VALUES ($1, $2, $3) RETURNING id, account_number, bank_name, bank_id_number", accountsTable)
	row := r.db.QueryRow(query, account.AccountNumber, account.BankName, account.BankIdNumber)
	if err := row.Scan(&account.Id, &account.AccountNumber, &account.BankName, &account.BankIdNumber); err != nil {
		return entity.Account{}, err
	}
	return account, nil
}

func (r *AccountPostgres) GetAccounts() ([]entity.Account, error) {
	var accounts []entity.Account
	query := fmt.Sprintf("SELECT * FROM %s", accountsTable)
	err := r.db.Select(&accounts, query)
	return accounts, err
}

func (r *AccountPostgres) GetAccountByAccountNumber(accountNumber string) (entity.Account, error) {
	var account entity.Account
	query := fmt.Sprintf("SELECT * FROM %s WHERE account_number = $1", accountsTable)
	err := r.db.Get(&account, query, accountNumber)
	return account, err
}

func (r *AccountPostgres) UpdateAccount(accountId int, updateData entity.UpdateAccountInput) (entity.Account, error) {
	var updateAccount entity.Account

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if updateData.AccountNumber != nil {
		setValues = append(setValues, fmt.Sprintf("account_number=$%d", argId))
		args = append(args, updateData.AccountNumber)
		argId++
	}
	if updateData.BankName != nil {
		setValues = append(setValues, fmt.Sprintf("bank_name=$%d", argId))
		args = append(args, updateData.BankName)
		argId++
	}
	if updateData.BankIdNumber != nil {
		setValues = append(setValues, fmt.Sprintf("bank_id_number=$%d", argId))
		args = append(args, updateData.BankIdNumber)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id, account_number, bank_name, bank_id_number", accountsTable, setQuery, argId)
	args = append(args, accountId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	err := r.db.Get(&updateAccount, query, args...)

	return updateAccount, err
}

func (r *AccountPostgres) DeleteAccount(accountId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE account_id = $1", accountsTable)
	_, err := r.db.Exec(query, accountId)
	return err
}

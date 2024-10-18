package entity

type Account struct {
	Id            int    `json:"-"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	BankIdNumber  string `json:"bank_identity_number"`
}

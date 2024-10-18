package entity

type Organization struct {
	Id              int    `json:"-"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	AccountId      int    `json:"account_id"`
	Chief           string `json:"chief"`
	FinancialChief string `json:"financial_chief"`
	InnKpp         string `json:"inn_kpp"`
}

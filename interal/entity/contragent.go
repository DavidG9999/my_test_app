package entity

type Contragent struct {
	Id      int    `json:"-"`
	Name    string `json:"name"`
	Address string `json:"address"`
	InnKpp  string `json:"inn_kpp"`
}

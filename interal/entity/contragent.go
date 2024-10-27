package entity

import "errors"

type Contragent struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name" binding:"required"`
	Address string `json:"address" db:"address" binding:"required"`
	InnKpp  string `json:"inn_kpp" db:"inn_kpp" binding:"required,min=20,max=20"`
}

type UpdateContragentInput struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
	InnKpp  *string `json:"inn_kpp"`
}

func (i UpdateContragentInput) Validate() error {
	if i.Name == nil && i.Address == nil && i.InnKpp == nil {
		return errors.New("update structure has no values")
	}
	if i.InnKpp != nil {
		if len(*i.InnKpp) != 20 {
			return errors.New("invalid field format: inn_kpp")
		}
	}
	if i.Name != nil {
		if *i.Name == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.Address != nil {
		if *i.Address == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}

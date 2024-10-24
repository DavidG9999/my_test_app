package entity

import "errors"

type Auto struct {
	Id          int    `json:"id" db:"id"`
	Brand       string `json:"brand" binding:"required" db:"brand"`
	Model       string `json:"model" binding:"required" db:"model"`
	StateNumber string `json:"state_number" binding:"required,min=8,max=9" db:"state_number"`
}

type UpdateAutoInput struct {
	Brand       *string `json:"brand"`
	Model       *string `json:"model"`
	StateNumber *string `json:"state_number"`
}

func (i UpdateAutoInput) Validate() error {
	if i.Brand == nil && i.Model == nil && i.StateNumber == nil {
		return errors.New("update structure has no values")
	}
	if i.StateNumber != nil {
		if len(*i.StateNumber) > 12 || len(*i.StateNumber) < 8 {
			return errors.New("invalid state_number param")
		}
	}
	return nil
}

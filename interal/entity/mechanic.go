package entity

import "errors"

type Mechanic struct {
	Id       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name" binding:"required"`
}

type UpdateMechanicInput struct {
	FullName *string `json:"full_name"`
}

func (i UpdateMechanicInput) Validate() error {
	if i.FullName == nil {
		return errors.New("update structure has no values")
	}
	if i.FullName != nil {
		if *i.FullName == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}

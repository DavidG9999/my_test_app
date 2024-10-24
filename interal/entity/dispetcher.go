package entity

import "errors"

type Dispetcher struct {
	Id       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name" binding:"required"`
}

type UpdateDispetcherInput struct {
	FullName *string `json:"full_name"`
}

func (i UpdateDispetcherInput) Validate() error {
	if i.FullName == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

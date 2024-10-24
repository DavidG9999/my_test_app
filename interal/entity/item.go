package entity

import "errors"

type Item struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

type UpdateItemInput struct {
	Name *string `json:"name"`
}

func (i UpdateItemInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

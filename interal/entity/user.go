package entity

import "errors"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Email    string `json:"email" binding:"required,email" db:"email"`
	Password string `json:"password" binding:"required" db:"password"`
}

type UpdateNameUserInput struct {
	Name *string `json:"name"`
}

type UpdatePasswordUserInput struct {
	Password *string `json:"password"`
}

func (i UpdateNameUserInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdatePasswordUserInput) Validate() error {
	if i.Password == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

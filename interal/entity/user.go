package entity

import "errors"

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Email    string `json:"email" binding:"required,email" db:"email"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type UpdateNameUserInput struct {
	Name *string `json:"name" binding:"required"`
}

type UpdatePasswordUserInput struct {
	Password *string `json:"password" binding:"required"`
}

func (i UpdateNameUserInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}
	if i.Name != nil {
		if *i.Name == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}

func (i UpdatePasswordUserInput) Validate() error {
	if i.Password == nil {
		return errors.New("update structure has no values")
	}
	if i.Password != nil {
		if *i.Password == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}

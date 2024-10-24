package entity

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Email    string `json:"email" binding:"required,email" db:"email"`
	Password string `json:"password" binding:"required" db:"password"`
}

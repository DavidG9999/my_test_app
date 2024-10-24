package entity

type Driver struct {
	Id       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name" binding:"required"`
	License  string `json:"license" db:"license" binding:"required,min=10,max=10"`
	Class    string `json:"class" db:"class" binding:"required"`
}

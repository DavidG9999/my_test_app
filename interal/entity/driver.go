package entity

type Driver struct {
	Id       int    `json:"-"`
	FullName string `json:"full_name"`
	License  string `json:"license"`
	Class    string `json:"class"`
}

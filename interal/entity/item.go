package entity

type Item struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
}

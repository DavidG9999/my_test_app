package entity

type Dispetcher struct {
	Id       int    `json:"-"`
	FullName string `json:"full_name"`
}

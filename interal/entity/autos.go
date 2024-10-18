package entity

type Auto struct {
	Id           int    `json:"-"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	StateN8umber string `json:"state_number"`
}

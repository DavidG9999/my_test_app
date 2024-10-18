package entity

type PutlistHeader struct {
	Id             int    `json:"-"`
	Number         int    `json:"number"`
	OrganizationId int    `json:"organiation_id"`
	DateWith       string `json:"date_with"`
	DateFor        string `json:"date_for"`
	AutoId         int    `json:"auto_id"`
	DriverId       string `json:"driver_id"`
	DispetcherId   int    `json:"dispetcher_id"`
	MehanicId      string `json:"mehanic_id"`
}

type PutlistBody struct {
	Id           int    `json:"-"`
	PutlistId    int    `json:"putlist_id"`
	ContragentId int    `json:"contragent_id"`
	ItemId       int    `json:"item_id"`
	TimeId       string `json:"time_id"`
	TimeFor      string `json:"time_for"`
}

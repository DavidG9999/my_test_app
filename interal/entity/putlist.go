package entity

import "errors"

type PutlistHeader struct {
	Id             int    `json:"id" db:"id"`
	Number         int    `json:"number" db:"number" binding:"required"`
	OrganizationId int    `json:"organiation_id" db:"organization_id" binding:"required"`
	DateWith       string `json:"date_with" db:"date_with" binding:"required"`
	DateFor        string `json:"date_for" db:"date_for" binding:"required"`
	AutoId         int    `json:"auto_id" db:"auto_id" binding:"required"`
	DriverId       int    `json:"driver_id" db:"driver_id" binding:"required"`
	DispetcherId   int    `json:"dispetcher_id" db:"dispetcher_id" binding:"required"`
	MechanicId     int    `json:"mechanic_id" db:"mechanic_id" binding:"required"`
}

type PutlistBody struct {
	Id           int    `json:"id" db:"id"`
	PutlistId    int    `json:"putlist_id" db:"putlist_id" binding:"required"`
	ContragentId int    `json:"contragent_id" db:"contragent_id" binding:"required"`
	ItemId       int    `json:"item_id" db:"item_id" binding:"required"`
	TimeWith     string `json:"time_with" db:"time_with" binding:"required"`
	TimeFor      string `json:"time_for" db:"time_for" binding:"required"`
}

type UpdatePutlistHeaderInput struct {
	Number         *int    `json:"number"`
	OrganizationId *int    `json:"organiation_id"`
	DateWith       *string `json:"date_with"`
	DateFor        *string `json:"date_for"`
	AutoId         *int    `json:"auto_id"`
	DriverId       *int    `json:"driver_id"`
	DispetcherId   *int    `json:"dispetcher_id"`
	MechanicId     *int    `json:"mechanic_id"`
}

type UpdatePutlistBodyInput struct {
	PutlistId    *int    `json:"putlist_id"`
	ContragentId *int    `json:"contragent_id"`
	ItemId       *int    `json:"item_id" db:"item_id"`
	TimeWith     *string `json:"time_with" db:"time_with"`
	TimeFor      *string `json:"time_for" db:"time_for"`
}

func (i UpdatePutlistHeaderInput) Validate() error {
	if i.Number == nil && i.OrganizationId == nil && i.DateWith == nil && i.DateFor == nil && i.AutoId == nil && i.DriverId == nil && i.DispetcherId == nil && i.MechanicId == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdatePutlistBodyInput) Validate() error {
	if i.PutlistId == nil && i.ContragentId == nil && i.ItemId == nil && i.TimeWith == nil && i.TimeFor == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

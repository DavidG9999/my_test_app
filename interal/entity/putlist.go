package entity

import "errors"

type PutlistHeader struct {
	Id           int    `json:"id" db:"id"`
	UserId       int    `json:"user_id" db:"user_id"`
	Number       int    `json:"number" db:"number" binding:"required"`
	AccountId    int    `json:"account_id" db:"account_id" binding:"required"`
	DateWith     string `json:"date_with" db:"date_with" binding:"required"`
	DateFor      string `json:"date_for" db:"date_for" binding:"required"`
	AutoId       int    `json:"auto_id" db:"auto_id" binding:"required"`
	DriverId     int    `json:"driver_id" db:"driver_id" binding:"required"`
	DispetcherId int    `json:"dispetcher_id" db:"dispetcher_id" binding:"required"`
	MechanicId   int    `json:"mechanic_id" db:"mechanic_id" binding:"required"`
}

type PutlistBody struct {
	Id            int    `json:"id" db:"id"`
	PutlistNumber int    `json:"putlist_number" db:"putlist_header_number"`
	Number        int    `json:"number" db:"number" binding:"required"`
	ContragentId  int    `json:"contragent_id" db:"contragent_id" binding:"required"`
	Item          string `json:"item" db:"item" binding:"required"`
	TimeWith      string `json:"time_with" db:"time_with" binding:"required"`
	TimeFor       string `json:"time_for" db:"time_for" binding:"required"`
}

type GetPutlistResponse struct {
	Number              int    `json:"number" db:"number" binding:"required"`
	Organization        string `json:"organization" db:"organization" binding:"required"`
	OrganizationAddress string `json:"organization_address" db:"organization_address" binding:"required"`
	Chief               string `json:"chief" db:"chief" binding:"required"`
	Financial_chief     string `json:"financial_chief" db:"financial_chief" binding:"required"`
	OrganizationInnKpp  string `json:"organization_inn_kpp" db:"organization_inn_kpp" binding:"required"`
	AccountNumber       string `json:"account_number" db:"account_number" binding:"required"`
	BankName            string `json:"bank_name" db:"bank_name" binding:"required"`
	BankIdNumber        string `json:"bank_id_number" db:"bank_id_number" binding:"required"`
	Date_with           string `json:"date_with" db:"date_with" binding:"required"`
	Date_for            string `json:"date_for" db:"date_for" binding:"required"`
	Brand               string `json:"brand" db:"brand" binding:"required"`
	Model               string `json:"model" db:"model" binding:"required"`
	StateNumber         string `json:"state_number" db:"state_number" binding:"required"`
	DriverFIO           string `json:"driver_fio" db:"driver_fio" binding:"required"`
	License             string `json:"license" db:"license" binding:"required"`
	Class               string `json:"class" db:"class" binding:"required"`
	DispetcherFIO       string `json:"dispetcher_fio" db:"dispetcher_fio" binding:"required"`
	MehanicFIO          string `json:"mehanic_fio" db:"mehanic_fio" binding:"required"`
}

type GetPutlistBodyResponse struct {
	Id                int    `json:"id,omitempty" db:"id"`
	PutlistNumber     int    `json:"putlist_number,omitempty" db:"putlist_number"`
	Number            int    `json:"putlist_body_number" db:"putlist_body_number"`
	Contragent        string `json:"contragent" db:"contragent" binding:"required"`
	ContragentAddress string `json:"contragent_address" db:"contragent_address" binding:"required"`
	ContragentInnKpp  string `json:"contragent_inn_kpp" db:"contragent_inn_kpp" binding:"required"`
	Item              string `json:"item" db:"item" binding:"required"`
	Time_with         string `json:"time_with" db:"time_with" binding:"required"`
	Time_for          string `json:"time_for" db:"time_for" binding:"required"`
}

type UpdatePutlistHeaderInput struct {
	AccountId    *int    `json:"account_id"`
	DateWith     *string `json:"date_with"`
	DateFor      *string `json:"date_for"`
	AutoId       *int    `json:"auto_id"`
	DriverId     *int    `json:"driver_id"`
	DispetcherId *int    `json:"dispetcher_id"`
	MechanicId   *int    `json:"mechanic_id"`
}

type UpdatePutlistBodyInput struct {
	Number       *int    `json:"number"`
	ContragentId *int    `json:"contragent_id"`
	Item         *string `json:"item"`
	TimeWith     *string `json:"time_with"`
	TimeFor      *string `json:"time_for"`
}

func (i UpdatePutlistHeaderInput) Validate() error {
	if i.AccountId == nil && i.DateWith == nil && i.DateFor == nil && i.AutoId == nil && i.DriverId == nil && i.DispetcherId == nil && i.MechanicId == nil {
		return errors.New("update structure has no values")
	}
	if i.AccountId != nil {
		if *i.AccountId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.DateWith != nil {
		if *i.DateWith == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.DateFor != nil {
		if *i.DateFor == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.AutoId != nil {
		if *i.AutoId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.DriverId != nil {
		if *i.DriverId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.DispetcherId != nil {
		if *i.DispetcherId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.MechanicId != nil {
		if *i.MechanicId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}

func (i UpdatePutlistBodyInput) Validate() error {
	if i.Number == nil && i.ContragentId == nil && i.Item == nil && i.TimeWith == nil && i.TimeFor == nil {
		return errors.New("update structure has no values")
	}
	if i.Number != nil {
		if *i.Number == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.ContragentId != nil {
		if *i.ContragentId == 0 {
			return errors.New("update structure has empty values")
		}
	}
	if i.Item != nil {
		if *i.Item == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.TimeWith != nil {
		if *i.TimeWith == "" {
			return errors.New("update structure has empty values")
		}
	}
	if i.TimeFor != nil {
		if *i.TimeFor == "" {
			return errors.New("update structure has empty values")
		}
	}
	return nil
}

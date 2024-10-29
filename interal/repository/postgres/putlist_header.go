package postgres

import (
	"fmt"
	"strings"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PutlistPostgres struct {
	db *sqlx.DB
}

func NewPutlistPostgres(db *sqlx.DB) *PutlistPostgres {
	return &PutlistPostgres{
		db: db,
	}
}

func (r *PutlistPostgres) CreatePutlist(userId int, putlist entity.PutlistHeader) (entity.PutlistHeader, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, number, account_id, date_with, date_for, auto_id, driver_id, dispetcher_id, mechanic_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, user_id, account_id, number, date_with, date_for, auto_id, driver_id, dispetcher_id, mechanic_id", putlistHeadersTable)
	row := r.db.QueryRow(query, userId, putlist.Number, putlist.AccountId, putlist.DateWith, putlist.DateFor, putlist.AutoId, putlist.DriverId, putlist.DispetcherId, putlist.MechanicId)
	if err := row.Scan(&putlist.Id, &putlist.UserId, &putlist.Number, &putlist.AccountId, &putlist.DateWith, &putlist.DateFor, &putlist.AutoId, &putlist.DriverId, &putlist.DispetcherId, &putlist.MechanicId); err != nil {
		return entity.PutlistHeader{}, err
	}
	return putlist, nil
}

func (r *PutlistPostgres) GetPutlistHeaders(userId int) ([]entity.GetPutlistResponse, error) {
	var putlists []entity.GetPutlistResponse

	query := fmt.Sprintf("SELECT ph.number, org.name AS organization, org.address AS organization_address, org.chief, org.financial_chief, org.inn_kpp AS organization_inn_kpp, ac.account_number, ac.bank_name, ac.bank_id_number, ph.date_with, ph.date_for, au.brand, au.model, au.state_number, dr.full_name AS driver_fio, dr.license, dr.class, di.full_name AS dispetcher_fio, mh.full_name AS mehanic_fio FROM %s AS ph JOIN %s AS ac ON ph.account_id = ac.id JOIN %s AS org ON ac.organization_id = org.id JOIN %s AS au ON ph.auto_id = au.id JOIN %s AS dr ON ph.driver_id = dr.id JOIN %s AS di ON ph.dispetcher_id = di.id JOIN %s AS mh ON ph.mechanic_id = mh.id WHERE user_id = $1", putlistHeadersTable, accountsTable, organizationsTable, autosTable, driversTable, dispetchersTable, mechanicsTable)
	err := r.db.Select(&putlists, query, userId)
	if err != nil {
		return []entity.GetPutlistResponse{}, err
	}

	return putlists, nil
}

func (r *PutlistPostgres) GetPutlistByNumber(userId int, putlistNumber int) (entity.GetPutlistResponse, []entity.GetPutlistBodyResponse, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return entity.GetPutlistResponse{}, []entity.GetPutlistBodyResponse{}, err
	}

	var putlist entity.GetPutlistResponse
	getPutlistQuery := fmt.Sprintf("SELECT ph.number, org.name AS organization, org.address AS organization_address, org.chief, org.financial_chief, org.inn_kpp AS organization_inn_kpp, ac.account_number, ac.bank_name, ac.bank_id_number, ph.date_with, ph.date_for, au.brand, au.model, au.state_number, dr.full_name AS driver_fio, dr.license, dr.class, di.full_name AS dispetcher_fio, mh.full_name AS mehanic_fio FROM %s AS ph JOIN %s AS ac ON ph.account_id = ac.id JOIN %s AS org ON ac.organization_id = org.id JOIN %s AS au ON ph.auto_id = au.id JOIN %s AS dr ON ph.driver_id = dr.id JOIN %s AS di ON ph.dispetcher_id = di.id JOIN %s AS mh ON ph.mechanic_id = mh.id WHERE user_id =$1 AND number = $2", putlistHeadersTable, accountsTable, organizationsTable, autosTable, driversTable, dispetchersTable, mechanicsTable)
	err = r.db.Get(&putlist, getPutlistQuery, userId, putlistNumber)
	if err != nil {
		tx.Rollback()
		return entity.GetPutlistResponse{}, []entity.GetPutlistBodyResponse{}, err
	}

	var putlistBodies []entity.GetPutlistBodyResponse
	getPullistBodyQuery := fmt.Sprintf("SELECT pb.number AS putlist_body_number, cnt.name AS contragent, cnt.address AS contragent_address, cnt.inn_kpp AS contragent_inn_kpp, pb.item, pb.time_with, pb.time_for FROM %s AS pb JOIN %s AS cnt ON pb.contragent_id = cnt.id WHERE putlist_header_number = $1", putlistBodiesTable, contragentsTable)
	err = r.db.Select(&putlistBodies, getPullistBodyQuery, putlistNumber)
	if err != nil {
		tx.Rollback()
		return entity.GetPutlistResponse{}, []entity.GetPutlistBodyResponse{}, err
	}

	return putlist, putlistBodies, tx.Commit()
}

func (r *PutlistPostgres) UpdatePutlist(userId int, putlistNumber int, updateData entity.UpdatePutlistHeaderInput) (entity.PutlistHeader, error) {
	var updatePutlist entity.PutlistHeader

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.AccountId != nil {
		setValues = append(setValues, fmt.Sprintf("account_id=$%d", argId))
		args = append(args, updateData.AccountId)
		argId++
	}
	if updateData.DateWith != nil {
		setValues = append(setValues, fmt.Sprintf("date_with=$%d", argId))
		args = append(args, updateData.DateWith)
		argId++
	}
	if updateData.DateFor != nil {
		setValues = append(setValues, fmt.Sprintf("date_for=$%d", argId))
		args = append(args, updateData.DateFor)
		argId++
	}
	if updateData.AutoId != nil {
		setValues = append(setValues, fmt.Sprintf("auto_id=$%d", argId))
		args = append(args, updateData.AutoId)
		argId++
	}
	if updateData.DriverId != nil {
		setValues = append(setValues, fmt.Sprintf("driver_id=$%d", argId))
		args = append(args, updateData.DriverId)
		argId++
	}
	if updateData.DispetcherId != nil {
		setValues = append(setValues, fmt.Sprintf("dispetcher_id=$%d", argId))
		args = append(args, updateData.DispetcherId)
		argId++
	}
	if updateData.MechanicId != nil {
		setValues = append(setValues, fmt.Sprintf("mechanic_id=$%d", argId))
		args = append(args, updateData.MechanicId)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id=$%d AND number=$%d  RETURNING id, number, account_id, date_with, date_for, auto_id, driver_id, dispetcher_id, mechanic_id, user_id", putlistHeadersTable, setQuery, argId, argId+1)
	args = append(args, userId, putlistNumber)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	err := r.db.Get(&updatePutlist, query, args...)

	return updatePutlist, err
}

func (r *PutlistPostgres) DeletePutlist(userId int, putlistNumber int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id =$1 AND number = $2", putlistHeadersTable)
	_, err := r.db.Exec(query, userId, putlistNumber)
	return err
}

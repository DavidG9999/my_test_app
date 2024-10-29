package postgres

import (
	"fmt"
	"strings"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PutlistBodyPostgres struct {
	db *sqlx.DB
}

func NewPutlistBodyPostgres(db *sqlx.DB) *PutlistBodyPostgres {
	return &PutlistBodyPostgres{
		db: db,
	}
}

func (r *PutlistBodyPostgres) CreatePutlistBody(putlistNumber int, putlistBody entity.PutlistBody) (entity.PutlistBody, error) {
	query := fmt.Sprintf("INSERT INTO %s (putlist_header_number, number, contragent_id, item, time_with, time_for) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, putlist_header_number, number, contragent_id, item, time_with, time_for", putlistBodiesTable)
	row := r.db.QueryRow(query, putlistNumber, putlistBody.Number, putlistBody.ContragentId, putlistBody.Item, putlistBody.TimeWith, putlistBody.TimeFor)
	if err := row.Scan(&putlistBody.Id, &putlistBody.PutlistNumber, &putlistBody.Number, &putlistBody.ContragentId, &putlistBody.Item, &putlistBody.TimeWith, &putlistBody.TimeFor); err != nil {
		return entity.PutlistBody{}, err
	}
	return putlistBody, nil
}

func (r *PutlistBodyPostgres) GetPutlistBodies(putlistNumber int) ([]entity.GetPutlistBodyResponse, error) {
	var putlistBodies []entity.GetPutlistBodyResponse

	query := fmt.Sprintf("SELECT pb.id, pb.putlist_header_number AS putlist_number, pb.number AS putlist_body_number, cnt.name AS contragent, cnt.address AS contragent_address, cnt.inn_kpp AS contragent_inn_kpp, pb.item, pb.time_with, pb.time_for FROM %s AS pb JOIN %s AS cnt ON pb.contragent_id = cnt.id WHERE putlist_header_number = $1", putlistBodiesTable, contragentsTable)
	err := r.db.Select(&putlistBodies, query, putlistNumber)
	if err != nil {
		return []entity.GetPutlistBodyResponse{}, err
	}
	return putlistBodies, nil
}

func (r *PutlistBodyPostgres) UpdatePutlistBody(putlistBodyId int, updateData entity.UpdatePutlistBodyInput) (entity.PutlistBody, error) {
	var updatePutlistBody entity.PutlistBody

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if updateData.Number != nil {
		setValues = append(setValues, fmt.Sprintf("number=$%d", argId))
		args = append(args, updateData.Number)
		argId++
	}
	if updateData.ContragentId != nil {
		setValues = append(setValues, fmt.Sprintf("contragent_id=$%d", argId))
		args = append(args, updateData.ContragentId)
		argId++
	}
	if updateData.Item != nil {
		setValues = append(setValues, fmt.Sprintf("item=$%d", argId))
		args = append(args, updateData.Item)
		argId++
	}
	if updateData.TimeWith != nil {
		setValues = append(setValues, fmt.Sprintf("time_with=$%d", argId))
		args = append(args, updateData.TimeWith)
		argId++
	}
	if updateData.TimeFor != nil {
		setValues = append(setValues, fmt.Sprintf("time_for=$%d", argId))
		args = append(args, updateData.TimeFor)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id, putlist_header_number, number, contragent_id, item, time_with, time_for", putlistBodiesTable, setQuery, argId)
	args = append(args, putlistBodyId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	err := r.db.Get(&updatePutlistBody, query, args...)

	return updatePutlistBody, err
}

func (r *PutlistBodyPostgres) DeletePutlistBody(putlistBodyId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", putlistBodiesTable)
	_, err := r.db.Exec(query, putlistBodyId)
	return err
}

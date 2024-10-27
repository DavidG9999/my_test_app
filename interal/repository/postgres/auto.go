package postgres

import (
	"fmt"
	"strings"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AutoPostgres struct {
	db *sqlx.DB
}

func NewAutoPostgres(db *sqlx.DB) *AutoPostgres {
	return &AutoPostgres{
		db: db,
	}
}

func (r *AutoPostgres) CreateAuto(auto entity.Auto) (entity.Auto, error) {
	query := fmt.Sprintf("INSERT INTO %s (brand, model, state_number) VALUES ($1, $2, $3) RETURNING id, brand, model, state_number", autosTable)
	row := r.db.QueryRow(query, auto.Brand, auto.Model, auto.StateNumber)
	if err := row.Scan(&auto.Id, &auto.Brand, &auto.Model, &auto.StateNumber); err != nil {
		return entity.Auto{}, err
	}
	return auto, nil
}

func (r *AutoPostgres) GetAutos() ([]entity.Auto, error) {
	var autos []entity.Auto
	query := fmt.Sprintf("SELECT * FROM %s", autosTable)
	err := r.db.Select(&autos, query)
	return autos, err
}

func (r *AutoPostgres) UpdateAuto(autoId int, updateData entity.UpdateAutoInput) (entity.Auto, error) {
	var updateAuto entity.Auto

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.Brand != nil {
		setValues = append(setValues, fmt.Sprintf("brand=$%d", argId))
		args = append(args, updateData.Brand)
		argId++
	}

	if updateData.Model != nil {
		setValues = append(setValues, fmt.Sprintf("model=$%d", argId))
		args = append(args, updateData.Model)
		argId++
	}

	if updateData.StateNumber != nil {
		setValues = append(setValues, fmt.Sprintf("state_number=$%d", argId))
		args = append(args, updateData.StateNumber)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id, brand, model, state_number", autosTable, setQuery, argId)
	args = append(args, autoId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	err := r.db.Get(&updateAuto, query, args...)
	return updateAuto, err

}

func (r *AutoPostgres) DeleteAuto(autoId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", autosTable)
	_, err := r.db.Exec(query, autoId)
	return err

}

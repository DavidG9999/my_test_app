package postgres

import (
	"fmt"
	"strings"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ContragentPostgres struct {
	db *sqlx.DB
}

func NewContragentPostgres(db *sqlx.DB) *ContragentPostgres {
	return &ContragentPostgres{
		db: db,
	}
}

func (r *ContragentPostgres) CreateContragent(contragent entity.Contragent) (entity.Contragent, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, address, inn_kpp) VALUES ($1, $2, $3) RETURNING id, name, address, inn_kpp", contragentsTable)
	row := r.db.QueryRow(query, contragent.Name, contragent.Address, contragent.InnKpp)
	if err := row.Scan(&contragent.Id, &contragent.Name, &contragent.Address, &contragent.InnKpp); err != nil {
		return entity.Contragent{}, err
	}
	return contragent, nil
}

func (r *ContragentPostgres) GetContragents() ([]entity.Contragent, error) {
	var contragents []entity.Contragent
	query := fmt.Sprintf("SELECT * FROM %s", contragentsTable)
	err := r.db.Select(&contragents, query)
	return contragents, err
}

func (r *ContragentPostgres) UpdateContragent(contragentId int, updateData entity.UpdateContragentInput) (entity.Contragent, error) {
	var updateContragent entity.Contragent

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, updateData.Name)
		argId++
	}

	if updateData.Address != nil {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, updateData.Address)
		argId++
	}

	if updateData.InnKpp != nil {
		setValues = append(setValues, fmt.Sprintf("inn_kpp=$%d", argId))
		args = append(args, updateData.InnKpp)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id, name, address, inn_kpp", contragentsTable, setQuery, argId)
	args = append(args, contragentId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	err := r.db.Get(&updateContragent, query, args...)
	return updateContragent, err

}

func (r *ContragentPostgres) DeleteContragent(contragentId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", contragentsTable)
	_, err := r.db.Exec(query, contragentId)
	return err

}

package postgres

import (
	"fmt"
	"strings"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DriverPostgres struct {
	db *sqlx.DB
}

func NewDriverPostgres(db *sqlx.DB) *DriverPostgres {
	return &DriverPostgres{
		db: db,
	}
}

func (r *DriverPostgres) CreateDriver(driver entity.Driver) (entity.Driver, error) {
	query := fmt.Sprintf("INSERT INTO %s (full_name,license,class) VALUES ($1,$2,$3) RETURNING id, full_name,license, class", driversTable)
	row := r.db.QueryRow(query, driver.FullName, driver.License, driver.Class)
	if err := row.Scan(&driver.Id, &driver.FullName, &driver.License, &driver.Class); err != nil {
		return entity.Driver{}, err
	}
	return driver, nil
}

func (r *DriverPostgres) GetDrivers() ([]entity.Driver, error) {
	var drivers []entity.Driver
	query := fmt.Sprintf("SELECT * FROM %s", driversTable)
	err := r.db.Select(&drivers, query)
	return drivers, err
}

func (r *DriverPostgres) UpdateDriver(driverId int, updateData entity.UpdateDriverInput) (entity.Driver, error) {
	var updateDriver entity.Driver

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updateData.FullName != nil {
		setValues = append(setValues, fmt.Sprintf("full_name=$%d", argId))
		args = append(args, updateData.FullName)
		argId++
	}

	if updateData.License != nil {
		setValues = append(setValues, fmt.Sprintf("license=$%d", argId))
		args = append(args, updateData.License)
		argId++
	}

	if updateData.Class != nil {
		setValues = append(setValues, fmt.Sprintf("class=$%d", argId))
		args = append(args, updateData.Class)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d RETURNING id, full_name, license, class", driversTable, setQuery, argId)
	args = append(args, driverId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	err := r.db.Get(&updateDriver, query, args...)
	return updateDriver, err

}

func (r *DriverPostgres) DeleteDriver(driverId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", driversTable)
	_, err := r.db.Exec(query, driverId)
	return err

}

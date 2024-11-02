package postgres

import (
	"fmt"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
)

type DispetcherPostgres struct {
	db *sqlx.DB
}

func NewDispetcherPostgres(db *sqlx.DB) *DispetcherPostgres {
	return &DispetcherPostgres{
		db: db,
	}
}

func (r *DispetcherPostgres) CreateDispetcher(dispetcher entity.Dispetcher) (entity.Dispetcher, error) {
	query := fmt.Sprintf("INSERT INTO %s (full_name) VALUES ($1) RETURNING id, full_name", dispetchersTable)
	row := r.db.QueryRow(query, dispetcher.FullName)
	if err := row.Scan(&dispetcher.Id, &dispetcher.FullName); err != nil {
		return entity.Dispetcher{}, err
	}
	return dispetcher, nil
}

func (r *DispetcherPostgres) GetDispetchers() ([]entity.Dispetcher, error) {
	var dispetchers []entity.Dispetcher
	query := fmt.Sprintf("SELECT * FROM %s", dispetchersTable)
	err := r.db.Select(&dispetchers, query)
	return dispetchers, err
}

func (r *DispetcherPostgres) UpdateDispetcher(dispetcherId int, updateData entity.UpdateDispetcherInput) (entity.Dispetcher, error) {
	var updateDispetcher entity.Dispetcher

	query := fmt.Sprintf("UPDATE %s SET full_name=$1 WHERE id=$2 RETURNING id, full_name", dispetchersTable)
	err := r.db.Get(&updateDispetcher, query, updateData.FullName, dispetcherId)
	return updateDispetcher, err

}

func (r *DispetcherPostgres) DeleteDispetcher(dispetcherId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", dispetchersTable)
	_, err := r.db.Exec(query, dispetcherId)
	return err

}

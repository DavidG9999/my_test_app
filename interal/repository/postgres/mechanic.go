package postgres

import (
	"fmt"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
)

type MechanicPostgres struct {
	db *sqlx.DB
}

func NewMechanicPostgres(db *sqlx.DB) *MechanicPostgres {
	return &MechanicPostgres{
		db: db,
	}
}

func (r *MechanicPostgres) CreateMechanic(mechanic entity.Mechanic) (entity.Mechanic, error) {
	query := fmt.Sprintf("INSERT INTO %s (full_name) VALUES ($1) RETURNING id, full_name", mechanicsTable)
	row := r.db.QueryRow(query, mechanic.FullName)
	if err := row.Scan(&mechanic.Id, &mechanic.FullName); err != nil {
		return entity.Mechanic{}, err
	}
	return mechanic, nil
}

func (r *MechanicPostgres) GetMechanics() ([]entity.Mechanic, error) {
	var meсhanics []entity.Mechanic
	query := fmt.Sprintf("SELECT * FROM %s", mechanicsTable)
	err := r.db.Select(&meсhanics, query)
	return meсhanics, err
}

func (r *MechanicPostgres) UpdateMechanic(meсhanicId int, updateData entity.UpdateMechanicInput) (entity.Mechanic, error) {
	var updateMeсhanic entity.Mechanic

	query := fmt.Sprintf("UPDATE %s SET full_name=$1 WHERE id=$2 RETURNING id, full_name", mechanicsTable)
	err := r.db.Get(&updateMeсhanic, query, updateData.FullName, meсhanicId)
	return updateMeсhanic, err

}

func (r *MechanicPostgres) DeleteMechanic(mechanicId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", mechanicsTable)
	_, err := r.db.Exec(query, mechanicId)
	return err

}

package postgres

import (
	"fmt"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUSer(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err

}

func (r *AuthPostgres) GetUserById(id int) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, id)

	return user, err

}

func (r *AuthPostgres) UpdateName(id int, updateData entity.UpdateNameUserInput) (entity.User, error) {
	var updateUser entity.User

	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2 RETURNING name, email, password_hash", usersTable)
	err := r.db.Get(&updateUser, query, updateData.Name, id)
	return updateUser, err
}

func (r *AuthPostgres) UpdatePassword(id int, updateData entity.UpdatePasswordUserInput) (entity.User, error) {
	var updateUser entity.User

	query := fmt.Sprintf("UPDATE %s SET password_hash=$1 WHERE id=$2 RETURNING name, email, password_hash", usersTable)
	err := r.db.Get(&updateUser, query, updateData.Password, id)
	return updateUser, err
}

func (r *AuthPostgres) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err := r.db.Exec(query, id)
	return err
}

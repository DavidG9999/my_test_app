package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	accountsTable = "accounts"
	autosTable = "autos"
	contragentsTable = "contragents"
	dispetchersTable = "dispetchers"
	driversTable = "drivers"
	itemsTable = "items"
	mechanicsTable = "mechanics"
	organizationsTable = "organizations"
	putlistHeadersTable = "putlist_headers"
	putlistBodiesTable = "putlist_bodies"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

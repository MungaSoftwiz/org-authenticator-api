package db

import (
	"database/sql"
	"fmt"

	"github.com/MungaSoftwiz/org-authenticator-api/config"
	_ "github.com/lib/pq"
)

func NewPostgreSQLStorage(cfg config.PostgreSQLConfig) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

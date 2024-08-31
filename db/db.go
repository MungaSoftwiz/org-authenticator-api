package db

import (
	"fmt"

	"github.com/MungaSoftwiz/org-authenticator-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgreSQLStorage(cfg config.PostgreSQLConfig) (*sqlx.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	sqlxDB := sqlx.NewDb(db.DB, "postgres")
	return sqlxDB, nil
}

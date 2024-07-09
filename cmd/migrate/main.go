package main

import (
	"log"
	"os"

	"github.com/MungaSoftwiz/org-authenticator-api/config"
	"github.com/MungaSoftwiz/org-authenticator-api/db"
	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.PostgreSQLConfig{
		Host:     config.Env.Host,
		User:     config.Env.User,
		Password: config.Env.Password,
		Port:     config.Env.Port,
		DBName:   config.Env.DBName,
	}

	db, err := db.NewPostgreSQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgresMigrate.WithInstance(db.DB, &postgresMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}

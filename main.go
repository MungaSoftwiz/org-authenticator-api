package main

import (
	"database/sql"
	"log"

	"github.com/MungaSoftwiz/org-authenticator-api/cmd/api"
	"github.com/MungaSoftwiz/org-authenticator-api/config"
	"github.com/MungaSoftwiz/org-authenticator-api/db"
)

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully!")
}

func main() {

	db, err := db.NewPostgreSQLStorage(config.PostgreSQLConfig{
		Host:     config.Env.Host,
		User:     config.Env.User,
		Password: config.Env.Password,
		Port:     config.Env.Port,
		DBName:   config.Env.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

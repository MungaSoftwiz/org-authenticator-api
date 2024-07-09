package api

import (
	"log"
	"net/http"

	"github.com/MungaSoftwiz/org-authenticator-api/service/org"
	"github.com/MungaSoftwiz/org-authenticator-api/service/user"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type APIServer struct {
	addr string
	db   *sqlx.DB
}

func NewAPIServer(addr string, db *sqlx.DB) *APIServer {
	return &APIServer{addr: addr, db: sqlx.NewDb(db.DB, "postgres")}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	userStorage := user.NewStorage(sqlx.NewDb(s.db.DB, "postgres"))
	userHandler := user.NewHandler(userStorage)
	userHandler.RegisterRoutes(subrouter)

	organisationStorage := org.NewOrganisationStorage(s.db)
	organisationHandler := org.NewOrganisationHandler(organisationStorage)
	organisationHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

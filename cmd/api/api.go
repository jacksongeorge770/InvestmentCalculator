package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacksongeorge770/InvestmentCalculator/service/calc"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	calcStore := calc.NewStore(s.db)
	calcHandler := calc.NewHandler(calcStore)
	calcHandler.RegisterRoutes(subrouter)

	fs := http.FileServer(http.Dir("./frontend"))
	router.PathPrefix("/").Handler(fs)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)

}

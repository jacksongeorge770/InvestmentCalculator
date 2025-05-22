package calc

import (
	"database/sql"

	"github.com/jacksongeorge770/InvestmentCalculator/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// save it to database
func (s *Store) SaveCalculation(calc types.Calculation) error {
	_, err := s.db.Exec(
		"INSERT INTO calculations (user_id, principal, rate, compounds_per_year, years, result, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		calc.UserID, calc.Principal, calc.Rate, calc.CompoundsPerYear, calc.Years, calc.Result, calc.CreatedAt,
	)
	return err
}

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

func (s *Store) SaveCalculation(calc types.Calculation) error {
	_, err := s.db.Exec(
		"INSERT INTO calculations (user_id, principal, rate, compounds_per_year, years, result, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		calc.UserID, calc.Principal, calc.Rate, calc.CompoundsPerYear, calc.Years, calc.Result, calc.CreatedAt,
	)
	return err
}

// func (s *Store) GetCalculationsByUserID(userID int) ([]types.Calculation, error) {
// 	rows, err := s.db.Query("SELECT * FROM calculations WHERE user_id = ?", userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var calcs []types.Calculation
// 	for rows.Next() {
// 		var c types.Calculation
// 		if err := rows.Scan(&c.ID, &c.UserID, &c.Principal, &c.Rate, &c.CompoundsPerYear, &c.Years, &c.Result, &c.CreatedAt); err != nil {
// 			return nil, err
// 		}
// 		calcs = append(calcs, c)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return calcs, nil
// }

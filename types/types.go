package types

import "time"

type CalculationPayload struct {
	Principal        float64 `json:"principal" validate:"required,gt=0"`
	Rate             float64 `json:"rate" validate:"required,gt=0"`
	CompoundsPerYear int     `json:"compoundsPerYear" validate:"required,gt=0"`
	Years            float64 `json:"years" validate:"required,gt=0"`
}

type Calculation struct {
	ID               int       `json:"id"`
	UserID           int       `json:"userId"`
	Principal        float64   `json:"principal"`
	Rate             float64   `json:"rate"`
	CompoundsPerYear int       `json:"compoundsPerYear"`
	Years            float64   `json:"years"`
	Result           float64   `json:"result"`
	CreatedAt        time.Time `json:"createdAt"`
}
type CalculationStore interface {
	SaveCalculation(Calculation) error
	// GetCalculationsByUserID(userID int) ([]Calculation, error)
}

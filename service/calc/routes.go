package calc

import (
	"math"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jacksongeorge770/InvestmentCalculator/types"
	"github.com/jacksongeorge770/InvestmentCalculator/utils"
)

type Handler struct {
	store types.CalculationStore
}

func NewHandler(store types.CalculationStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/calculate", h.handleCalculate).Methods(http.MethodPost)
	// router.HandleFunc("/history", h.handleGetHistory).Methods(http.MethodGet)
}

func (h *Handler) handleCalculate(w http.ResponseWriter, r *http.Request) {
	var payload types.CalculationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//calculation
	result := payload.Principal * math.Pow(
		1+payload.Rate/float64(payload.CompoundsPerYear),
		float64(payload.CompoundsPerYear)*payload.Years,
	)

	calc := types.Calculation{
		UserID:           1,
		Principal:        payload.Principal,
		Rate:             payload.Rate,
		CompoundsPerYear: payload.CompoundsPerYear,
		Years:            payload.Years,
		Result:           result,
		CreatedAt:        time.Now(),
	}
	if err := h.store.SaveCalculation(calc); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]float64{"result": result})
}

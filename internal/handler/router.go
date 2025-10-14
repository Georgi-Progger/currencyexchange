package handler

import (
	"database/sql"
	"net/http"
)

func (h *Handler) SetupRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Currency routes
	mux.HandleFunc("POST /api/currency", h.CreateCurrency)
	mux.HandleFunc("GET /api/currencies", h.GetCurrencies)
	mux.HandleFunc("GET /api/currency/{code}", h.GetCurrency)

	// Exchangerate cuurency routes
	mux.HandleFunc("POST /api/exchangeRates", h.CreateExchangerate)
	mux.HandleFunc("GET /api/exchangeRates", h.GetExchangeRates)
	mux.HandleFunc("GET /api/exchangeRates/{codes}", h.GetExchangeRate)

	return mux
}

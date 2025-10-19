package handler

import (
	"currencyexchange/internal/middleware"
	"database/sql"
	"net/http"
)

func (h *Handler) SetupRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	// Currency routes
	mux.HandleFunc("POST /api/currency", h.CreateCurrency)
	mux.HandleFunc("GET /api/currencies", h.GetCurrencies)
	mux.HandleFunc("GET /api/currency/{code}", h.GetCurrency)

	// Exchangerate cuurency routes
	mux.HandleFunc("POST /api/exchangeRates", h.CreateExchangerate)
	mux.HandleFunc("GET /api/exchangeRates", h.GetExchangeRates)
	mux.HandleFunc("GET /api/exchangeRates/{codes}", h.GetExchangeRate)
	mux.HandleFunc("PATCH /api/exchangeRates", h.UpdateExchangeRate)
	mux.HandleFunc("GET /api/exchangeRates/exchange", h.GetCalculateExchangerate)

	return middleware.Logging(mux)
}

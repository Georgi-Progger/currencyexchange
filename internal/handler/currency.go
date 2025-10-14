package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (h *Handler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		newErrorResponse(err.Error())
	}

	code := r.FormValue("code")
	name := r.FormValue("name")
	symbol := r.FormValue("symbol")

	if len(code) == 0 || len(name) == 0 || len(symbol) == 0 {
		newErrorResponse("value is empty")
	}

	currency, err := h.usecase.CreateCurrency(r.Context(), code, name, symbol)
	if err != nil {
		newErrorResponse(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currency)
}

func (h *Handler) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.usecase.GetCurrencies(r.Context())
	if err != nil {
		newErrorResponse(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currencies)
}

func (h *Handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/api/currency/")
	code = strings.TrimSpace(code)

	if len(code) == 0 {
		newErrorResponse("currency code is required")
	}

	currencies, err := h.usecase.GetCurrency(r.Context(), code)
	if err != nil {
		newErrorResponse(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currencies)
}

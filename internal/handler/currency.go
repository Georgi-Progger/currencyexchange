package handler

import (
	exerror "currencyexchange/internal/exerrors"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (h *Handler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		JSONError(w, "failed parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	name := r.FormValue("name")
	symbol := r.FormValue("symbol")

	if len(code) == 0 || len(name) == 0 || len(symbol) == 0 {
		JSONError(w, "fields is empty", http.StatusBadRequest)
		return
	}

	currency, err := h.usecase.CreateCurrency(r.Context(), code, name, symbol)
	if err != nil {
		if errors.Is(err, exerror.ErrCurrencyExists) {
			JSONError(w, exerror.ErrCurrencyExists.Error(), http.StatusConflict)
			return
		}
		JSONError(w, "failed create currency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currency)
}

func (h *Handler) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.usecase.GetCurrencies(r.Context())
	if err != nil {
		JSONError(w, "failed get currencies: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currencies)
}

func (h *Handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/api/currency/")
	code = strings.TrimSpace(code)

	if len(code) == 0 {
		JSONError(w, "currency code is empty", http.StatusBadRequest)
		return
	}

	currencies, err := h.usecase.GetCurrency(r.Context(), code)
	if err != nil {
		JSONError(w, "failed get currency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currencies)
}

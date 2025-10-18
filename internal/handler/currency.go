// handler/currency.go
package handler

import (
	"currencyexchange/internal/apperror"
	"currencyexchange/internal/models"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (h *Handler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	var req models.Currency
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JSONError(w, "invalid JSON format", http.StatusBadRequest)
		return
	}

	if len(req.Code) == 0 || len(req.FullName) == 0 || len(req.Sign) == 0 {
		JSONError(w, "fields code, full_name and sign are required", http.StatusBadRequest)
		return
	}

	currency, err := h.usecase.CreateCurrency(r.Context(), req)
	if err != nil {
		if errors.Is(err, apperror.ErrCurrencyExists) {
			JSONError(w, apperror.ErrCurrencyExists.Error(), http.StatusConflict)
			return
		}
		JSONError(w, "failed to create currency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, currency, http.StatusCreated)
}

func (h *Handler) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.usecase.GetCurrencies(r.Context())
	if err != nil {
		JSONError(w, "failed to get currencies: "+err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, currencies, http.StatusOK)
}

func (h *Handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/api/currency/")
	code = strings.TrimSpace(code)

	if len(code) == 0 {
		JSONError(w, "currency code is required", http.StatusBadRequest)
		return
	}

	currency, err := h.usecase.GetCurrency(r.Context(), code)
	if err != nil {
		if errors.Is(err, apperror.ErrCurrencyNotExists) {
			JSONError(w, apperror.ErrCurrencyNotExists.Error(), http.StatusNotFound)
			return
		}
		JSONError(w, "failed to get currency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, currency, http.StatusOK)
}

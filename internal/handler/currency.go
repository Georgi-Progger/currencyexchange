package handler

import (
	"currencyexchange/internal/apperror"
	"currencyexchange/internal/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	var req models.Currency
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JSONError(w, apperror.ErrInvalidJSON)
		return
	}

	if err := req.Validate(); err != nil {
		JSONError(w, err)
		return
	}

	currency, err := h.usecase.CreateCurrency(r.Context(), req)
	if err != nil {
		JSONError(w, err)
		return
	}

	JSONResponse(w, currency, http.StatusCreated)
}

func (h *Handler) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.usecase.GetCurrencies(r.Context())
	if err != nil {
		JSONError(w, err)
		return
	}

	JSONResponse(w, currencies, http.StatusOK)
}

func (h *Handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if len(code) == 0 {
		JSONError(w, apperror.ErrValidation)
		return
	}

	currency, err := h.usecase.GetCurrency(r.Context(), code)
	if err != nil {
		JSONError(w, err)
		return
	}

	JSONResponse(w, currency, http.StatusOK)
}

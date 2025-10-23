package handler

import (
	"currencyexchange/internal/apperror"
	"currencyexchange/internal/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateExchangerate(w http.ResponseWriter, r *http.Request) {
	var req models.ExchangeRateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JSONError(w, apperror.ErrInvalidJSON)
		return
	}

	if err := req.Validate(); err != nil {
		JSONError(w, err)
		return
	}

	exchange, err := h.usecase.CreateExchangeRate(r.Context(), req.Rate.String(), req.BaseCurrency, req.TargetCurrency)
	if err != nil {
		JSONError(w, err)
		return
	}

	JSONResponse(w, exchange, http.StatusCreated)
}

func (h *Handler) GetExchangeRates(w http.ResponseWriter, r *http.Request) {
	exchangerates, err := h.usecase.GetExchangeRates(r.Context())
	if err != nil {
		JSONError(w, err)
		return
	}
	JSONResponse(w, exchangerates, http.StatusOK)
}

func (h *Handler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	codes := r.URL.Query().Get("codes")
	if len(codes) == 0 {
		JSONError(w, apperror.ErrValidation)
		return
	}

	exchangerate, err := h.usecase.GetExchangeRate(r.Context(), codes)
	if err != nil {
		JSONError(w, err)
		return
	}

	JSONResponse(w, exchangerate, http.StatusOK)
}

func (h *Handler) UpdateExchangeRate(w http.ResponseWriter, r *http.Request) {
	codes := r.URL.Query().Get("codes")

	if len(codes) == 0 {
		JSONError(w, apperror.ErrValidation)
		return
	}

	var req models.UpdateExchangeRateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JSONError(w, err)
		return
	}

	exchangerate, err := h.usecase.UpdateExchangeRate(r.Context(), codes, req.Rate.String())
	if err != nil {
		JSONError(w, err)
		return
	}

	JSONResponse(w, exchangerate, http.StatusOK)
}

func (h *Handler) GetCalculateExchangerate(w http.ResponseWriter, r *http.Request) {
	baseCurrency := r.URL.Query().Get("from")
	targetCurrency := r.URL.Query().Get("to")
	amount := r.URL.Query().Get("amount")

	if len(baseCurrency) == 0 || len(targetCurrency) == 0 || len(amount) == 0 {
		JSONError(w, apperror.ErrValidation)
		return
	}

	calculateExchangerate, err := h.usecase.CalculateExchangeRate(r.Context(), baseCurrency, targetCurrency, amount)
	if err != nil {
		JSONError(w, err)
		return
	}

	JSONResponse(w, calculateExchangerate, http.StatusOK)
}

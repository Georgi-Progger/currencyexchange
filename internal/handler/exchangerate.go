package handler

import (
	"currencyexchange/internal/apperror"
	"currencyexchange/internal/models"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (h *Handler) CreateExchangerate(w http.ResponseWriter, r *http.Request) {
	var req models.ExchangeRateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JSONError(w, "invalid JSON format", http.StatusBadRequest)
		return
	}

	if len(req.BaseCurrency) == 0 || len(req.TargetCurrency) == 0 {
		JSONError(w, "fields is empty", http.StatusBadRequest)
		return
	}

	exchange, err := h.usecase.CreateExchangeRate(r.Context(), req.Rate.String(), req.BaseCurrency, req.TargetCurrency)
	if err != nil {
		JSONError(w, "exchangerate create failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	JSONResponse(w, exchange, http.StatusCreated)
}

func (h *Handler) GetExchangeRates(w http.ResponseWriter, r *http.Request) {
	exchangerates, err := h.usecase.GetExchangeRates(r.Context())
	if err != nil {
		JSONError(w, "get exchangerates failed: "+err.Error(), http.StatusBadRequest)
		return
	}
	JSONResponse(w, exchangerates, http.StatusOK)
}

func (h *Handler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	codes := strings.TrimPrefix(r.URL.Path, "/api/exchangerate/")
	codes = strings.TrimSpace(codes)

	if len(codes) == 0 {
		JSONError(w, "codes are empty", http.StatusBadRequest)
		return
	}

	exchangerate, err := h.usecase.GetExchangeRate(r.Context(), codes)
	if err != nil {
		JSONError(w, "get exchangerate reate failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	JSONResponse(w, exchangerate, http.StatusOK)
}

func (h *Handler) UpdateExchangeRate(w http.ResponseWriter, r *http.Request) {
	codes := strings.TrimPrefix(r.URL.Path, "/api/exchangerate/")
	codes = strings.TrimSpace(codes)

	if len(codes) == 0 {
		JSONError(w, "codes are empty", http.StatusBadRequest)
		return
	}

	var req models.UpdateExchangeRateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JSONError(w, "invalid JSON format", http.StatusBadRequest)
		return
	}

	exchangerate, err := h.usecase.UpdateExchangeRate(r.Context(), codes, req.Rate.String())
	if err != nil {
		JSONError(w, "update exchangerate failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	JSONResponse(w, exchangerate, http.StatusOK)
}

func (h *Handler) GetCalculateExchangerate(w http.ResponseWriter, r *http.Request) {
	baseCurrency := r.URL.Query().Get("from")
	targetCurrency := r.URL.Query().Get("to")
	amount := r.URL.Query().Get("amount")

	if len(baseCurrency) == 0 || len(targetCurrency) == 0 || len(amount) == 0 {
		JSONError(w, "fields are empty", http.StatusBadRequest)
		return
	}

	calculateExchangerate, err := h.usecase.CalculateExchangeRate(r.Context(), baseCurrency, targetCurrency, amount)
	if err != nil {
		if errors.Is(err, apperror.ErrCurrencyNotExists) {
			JSONError(w, apperror.ErrCurrencyNotExists.Error(), http.StatusBadRequest)
			return
		}
		JSONError(w, "get exchangerate reate failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	JSONResponse(w, calculateExchangerate, http.StatusOK)
}

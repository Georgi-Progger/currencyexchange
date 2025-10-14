package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (h *Handler) CreateExchangerate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		newErrorResponse(err.Error())
	}

	rate := r.FormValue("rate")
	baseCurrencyCode := r.FormValue("baseCurrencyCode")
	targetCurrencyCode := r.FormValue("targetCurrencyCode")

	if len(baseCurrencyCode) == 0 || len(targetCurrencyCode) == 0 {
		newErrorResponse("value is empty")
	}

	currency, err := h.usecase.CreateExchangeRate(r.Context(), rate, baseCurrencyCode, targetCurrencyCode)
	if err != nil {
		newErrorResponse(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currency)
}

func (h *Handler) GetExchangeRates(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.usecase.GetExchangeRates(r.Context())
	if err != nil {
		newErrorResponse(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currencies)
}

func (h *Handler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	codes := strings.TrimPrefix(r.URL.Path, "/api/exchangerate/")
	codes = strings.TrimSpace(codes)

	if len(codes) == 0 {
		newErrorResponse("currency code is required")
	}

	currencies, err := h.usecase.GetExchangeRate(r.Context(), codes)
	if err != nil {
		newErrorResponse(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currencies)
}

func (h *Handler) UpdateExchangeRate(w http.ResponseWriter, r *http.Request) {
	codes := strings.TrimPrefix(r.URL.Path, "/api/exchangerate/")
	codes = strings.TrimSpace(codes)

	if len(codes) == 0 {
		newErrorResponse("currency code is required")
	}

	err := r.ParseForm()
	if err != nil {
		newErrorResponse(err.Error())
	}

	rate := r.FormValue("rate")

	currencies, err := h.usecase.UpdateExchangeRate(r.Context(), codes, rate)
	if err != nil {
		newErrorResponse(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currencies)
}

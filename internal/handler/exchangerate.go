package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (h *Handler) CreateExchangerate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		JSONError(w, "failed parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	rate := r.FormValue("rate")
	baseCurrencyCode := r.FormValue("baseCurrencyCode")
	targetCurrencyCode := r.FormValue("targetCurrencyCode")

	if len(baseCurrencyCode) == 0 || len(targetCurrencyCode) == 0 {
		JSONError(w, "fields is empty", http.StatusBadRequest)
		return
	}

	currency, err := h.usecase.CreateExchangeRate(r.Context(), rate, baseCurrencyCode, targetCurrencyCode)
	if err != nil {
		JSONError(w, "exchangerate create failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currency)
}

func (h *Handler) GetExchangeRates(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.usecase.GetExchangeRates(r.Context())
	if err != nil {
		JSONError(w, "get exchangerates failed: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currencies)
}

func (h *Handler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	codes := strings.TrimPrefix(r.URL.Path, "/api/exchangerate/")
	codes = strings.TrimSpace(codes)

	if len(codes) == 0 {
		JSONError(w, "codes are empty", http.StatusBadRequest)
		return
	}

	currencies, err := h.usecase.GetExchangeRate(r.Context(), codes)
	if err != nil {
		JSONError(w, "get exchangerate reate failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currencies)
}

func (h *Handler) UpdateExchangeRate(w http.ResponseWriter, r *http.Request) {
	codes := strings.TrimPrefix(r.URL.Path, "/api/exchangerate/")
	codes = strings.TrimSpace(codes)

	if len(codes) == 0 {
		JSONError(w, "codes are empty", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		JSONError(w, "parse form failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	rate := r.FormValue("rate")

	currencies, err := h.usecase.UpdateExchangeRate(r.Context(), codes, rate)
	if err != nil {
		JSONError(w, "update exchangerate failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currencies)
}

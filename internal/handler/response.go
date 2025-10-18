package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func JSONError(w http.ResponseWriter, message string, statusCode int) {
	JSONResponse(w, ErrorResponse{Message: message}, statusCode)
}

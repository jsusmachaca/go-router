package response

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, body any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func JsonErrorFromString(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	respone := map[string]string{
		"error": errorMessage,
	}
	if err := json.NewEncoder(w).Encode(respone); err != nil {
		http.Error(w, errorMessage, statusCode)
	}
}

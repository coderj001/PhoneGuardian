package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
	return
}

func HealthCheck(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  "Ok",
		"message": "API is healthy",
	}

	respondJSON(w, http.StatusOK, response)
}

package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func isValidStatus(status string) bool {
	validStatuses := []string{"pending", "in-progress", "completed"}
	for _, s := range validStatuses {
		if status == s {
			return true
		}
	}
	return false
}

func isValidPriority(priority string) bool {
	validPriorities := []string{"low", "medium", "high"}
	for _, p := range validPriorities {
		if priority == p {
			return true
		}
	}
	return false
}

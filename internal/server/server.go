package server

import (
	"encoding/json"
	"net/http"
)

func Error(data map[string]interface{}, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write(jsonResponse)
	return
}

func Ok(data map[string]interface{}, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	return
}

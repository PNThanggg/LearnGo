package Project_01

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, code, errorResponse{Error: message})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json payload: %v\n", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err1 := w.Write(data)
	if err1 != nil {
		log.Printf("Failed to write response: %v\n", err1)
		return
	}
}

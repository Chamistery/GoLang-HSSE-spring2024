package server

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type DecodeRequest struct {
	InputString string `json:"inputString"`
}

type DecodeResponse struct {
	OutputString string `json:"outputString"`
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("v1.0.0"))
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	var req DecodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(req.InputString)
	if err != nil {
		http.Error(w, "Failed to decode base64 string", http.StatusBadRequest)
		return
	}

	resp := DecodeResponse{OutputString: string(decoded)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func HardOpHandler(w http.ResponseWriter, r *http.Request) {
	delay := rand.Intn(11) + 10
	time.Sleep(time.Duration(delay) * time.Second)
	if rand.Intn(2) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrSeatAlreadyBooked = errors.New("seat is already booked")
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

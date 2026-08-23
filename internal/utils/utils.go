package utils

import (
	"context"
	"encoding/json"
	"errors"
	goredis "github.com/redis/go-redis/v9"
	"log"
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

func NewClient(addr string) *goredis.Client {
	rdb := goredis.NewClient(&goredis.Options{Addr: addr})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis ping: %v", err)
	}
	log.Printf("connected to redis at %s", addr)

	return rdb
}

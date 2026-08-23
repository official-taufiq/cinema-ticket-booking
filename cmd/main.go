package main

import (
	"log"
	"net/http"

	"github.com/official-taufiq/cinema-ticket-booking/internal/booking"
	"github.com/redis/go-redis/v9"
)

func main() {

	mux := http.NewServeMux()

	store := booking.NewRedisStore(redis.NewClient(&redis.Options{Network: "tcp", Addr: "localhost:6379"}))
	svc := booking.NewService(store)

	bookingHandler := booking.NewHandler(svc)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))
	mux.HandleFunc("GET /movies", booking.ListMovies)
	mux.HandleFunc("GET /movies/{movieID}/seats", bookingHandler.ListSeats)
	mux.HandleFunc("POST /movies/{movieID}/seats/{seatID}/hold", bookingHandler.HoldSeat)

	mux.HandleFunc("PUT /sessions/{sessionID}/confirm", bookingHandler.ConfirmSession)
	mux.HandleFunc("DELETE /sessions/{sessionID}", bookingHandler.ReleaseSession)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/official-taufiq/cinema-ticket-booking/internal/booking"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir("static")))
	mux.HandleFunc("GET /movies", booking.ListMovies)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

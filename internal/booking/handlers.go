package booking

import (
	"encoding/json"
	"net/http"
)

func ListMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

var movies = []struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}{
	{ID: "inception", Title: "Inception", Rows: 5, SeatsPerRow: 6},
	{ID: "dune", Title: "Dune", Rows: 4, SeatsPerRow: 6},
}

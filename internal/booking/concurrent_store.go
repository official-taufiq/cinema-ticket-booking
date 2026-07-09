package booking

import "github.com/official-taufiq/cinema-ticket-booking/internal/booking/utils"

type ConcurrentStore struct {
	bookings map[string]Booking
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurrentStore) Book(b Booking) error {
	if _, ok := s.bookings[b.SeatID]; !ok {
		s.bookings[b.SeatID] = b
		return nil
	}
	return utils.ErrSeatAlreadyBooked
}

func (s *ConcurrentStore) ListBookings(movieID string) []Booking {
	var bookings []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			bookings = append(bookings, b)
		}
	}
	return bookings
}

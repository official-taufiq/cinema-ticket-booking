package booking

import (
	"sync"

	"github.com/official-taufiq/cinema-ticket-booking/internal/booking/utils"
)

type MemoryStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.bookings[b.SeatID]; !ok {
		s.bookings[b.SeatID] = b
		return nil
	}
	return utils.ErrSeatAlreadyBooked
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {
	s.RLock()
	defer s.RUnlock()

	var bookings []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			bookings = append(bookings, b)
		}
	}
	return bookings
}

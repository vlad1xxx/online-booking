package service

import "context"

type CreateBookingRequest struct {
	ChatID int64
	Date   string
	Time   string
	Phone  string
}

type BookingClient interface {
	CreateBooking(ctx context.Context, req *CreateBookingRequest) error
	GetAvailableDates(ctx context.Context) ([]string, error)
	GetAvailableTimes(ctx context.Context, selectedDate string) ([]string, error)
	// CreateReservation(ctx context.Context)
}

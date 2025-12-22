package service

import "context"

type BookingClient interface {
	GetAvailableDates(ctx context.Context) ([]string, error)
	CreateBooking(ctx context.Context, req *CreateBookingRequest) error
}

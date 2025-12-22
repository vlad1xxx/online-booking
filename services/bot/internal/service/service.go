package service

type Service struct {
	booking BookingClient
	state   StateStore
}

func New(booking BookingClient, state StateStore) *Service {
	return &Service{
		booking: booking,
		state:   state,
	}
}

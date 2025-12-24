package client

import (
	"context"
	"time"

	"github.com/vlad1xxx/online-booking/services/bot/internal/service"
)

type FakeBookingClient struct {
	num int
}

func NewFakeBookingClient() *FakeBookingClient {
	return &FakeBookingClient{
		num: 1,
	}
}

func (f *FakeBookingClient) CreateBooking(ctx context.Context, req *service.CreateBookingRequest) error {
	return nil
}

func (f *FakeBookingClient) GetAvailableDates(ctx context.Context) ([]string, error) {
	dates := make([]string, 10)
	today := time.Now()
	for i := 0; i < 10; i++ {
		date := today.AddDate(0, 0, 1)
		dates[i] = date.Format("2006-01-02")
	}

	return dates, nil
}

func (f *FakeBookingClient) GetAvailableTimes(ctx context.Context, selectedDate string) ([]string, error) {
	return []string{"10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00"}, nil
}



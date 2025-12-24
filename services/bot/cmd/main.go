package main

import (
	"log"
	"time"

	"github.com/vlad1xxx/online-booking/services/bot/internal/client"
	"github.com/vlad1xxx/online-booking/services/bot/internal/handler"
	"github.com/vlad1xxx/online-booking/services/bot/internal/service"
	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "8074091096:AAE0ENx1A6CIKHLUbzomepRY1TF8Ua83PJk",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	state := service.NewMemoryState()
	bookingClient := client.NewFakeBookingClient()

	svc := service.New(bookingClient, state)
	h := handler.NewHandler(b, svc)
	h.RegisterHandlers()

	b.Start()
}

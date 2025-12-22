package main

import (
	"fmt"
	"log"
	"time"

	"github.com/vlad1xxx/online-booking/services/bot/internal/handler"
	"github.com/vlad1xxx/online-booking/services/bot/internal/service"
	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "8074091096:AAGf1qF1a7zr4QL4VjV6QkW-8Jd3IvJOSvo",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	state := service.NewMemoryState()
	bookingClient := service.

	service := service.New(state)
	h := handler.NewHandler(b)
}

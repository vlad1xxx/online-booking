package handler

import (
	"log"
	"strings"

	"github.com/vlad1xxx/online-booking/services/bot/internal/service"
	tele "gopkg.in/telebot.v4"
)

type Handler struct {
	bot     *tele.Bot
	service *service.Service
}

func NewHandler(bot *tele.Bot, service *service.Service) *Handler {
	return &Handler{
		bot:     bot,
		service: service,
	}
}

func (h *Handler) RegisterHandlers() {
	h.registerStart()
	h.registerCallbacks()
	h.registerStaticButtons()
}

func (h *Handler) registerStart() {
	menu := h.service.CreateBookKeyboard()
	h.bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Добро пожаловать", menu)
	})
}

func (h *Handler) registerCallbacks() {
	h.bot.Handle(tele.OnCallback, func(c tele.Context) error {
		data := c.Callback().Data
		log.Printf("Received callback: %s", data)

		data = strings.TrimSpace(data)

		if data == "cancel_booking" {
			c.Respond()
			return h.service.HandleCancel(c)
		}

		if strings.HasPrefix(data, "date_") {
			date := strings.TrimPrefix(data, "date_")
			c.Respond()
			if err := h.service.HandleDateSelection(c, date); err != nil {
				return err
			}
			return nil
		}

		if strings.HasPrefix(data, "time_") {
			time := strings.TrimPrefix(data, "time_")
			c.Respond()
			if err := h.service.HandleTimeSelection(c, time); err != nil {
				return err
			}
			return nil
		}

		c.Respond()
		return nil
	})
}

func (h *Handler) registerStaticButtons() {
	h.bot.Handle(&service.BtnBook, func(c tele.Context) error {
		c.Respond()
		return h.service.StartBooking(c)
	})
}

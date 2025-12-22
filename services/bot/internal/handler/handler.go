package handler

import (
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
	menu := h.service.CreateBookMenu()

	h.bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Добро пожаловать", menu)
	})

	h.bot.Handle(&service.BtnBook, func(c tele.Context) error {
		return h.service.StartBooking(c)
	})
}

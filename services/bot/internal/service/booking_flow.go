package service

import tele "gopkg.in/telebot.v4"

func (s *Service) StartBooking(c tele.Context) error {
	return c.Send("book")
}

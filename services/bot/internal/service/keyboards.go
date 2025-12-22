package service

import tele "gopkg.in/telebot.v4"

func (s *Service) CreateBookMenu() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	menu.Inline(menu.Row(BtnBook))
	return menu
}

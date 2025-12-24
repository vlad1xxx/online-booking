package service

import tele "gopkg.in/telebot.v4"

func (s *Service) CreateBookKeyboard() *tele.ReplyMarkup {
	menu := &tele.ReplyMarkup{}
	menu.Inline(menu.Row(BtnBook))
	return menu
}

func (s *Service) CreateDateKeyboard(dates []string) *tele.ReplyMarkup {
	keyboard := &tele.ReplyMarkup{}
	var rows []tele.Row

	for _, date := range dates {
		btn := keyboard.Data(date, "date_"+date)
		rows = append(rows, tele.Row{btn})
	}

	rows = append(rows, keyboard.Row(BtnCancel))

	keyboard.Inline(rows...)
	return keyboard
}

func (s *Service) CreateTimeKeyboard(times []string) *tele.ReplyMarkup {
	keyboard := &tele.ReplyMarkup{}
	var rows []tele.Row

	for _, time := range times {

		btn := keyboard.Data(time, "time_"+time)
		rows = append(rows, tele.Row{btn})
	}

	rows = append(rows, keyboard.Row(BtnCancel))

	keyboard.Inline(rows...)
	return keyboard
}

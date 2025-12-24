package service

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func (s *Service) StartBooking(c tele.Context) error {
	s.state.Set(c.Chat().ID, StepChooseDate)
	dates, err := s.booking.GetAvailableDates(context.Background())
	fmt.Println(dates)
	if err != nil {
		return c.Send("Ошибка при загрузке дат")
	}

	kb := s.CreateDateKeyboard(dates)
	return c.Send("Выберите дату на которую вы хотите записаться", kb)
}

func (s *Service) HandleDateSelection(c tele.Context, selectedDate string) error {
	s.state.Update(c.Chat().ID, selectedDate)
	s.state.Set(c.Chat().ID, StepChooseTime)
	times, err := s.booking.GetAvailableTimes(context.Background(), selectedDate)
	if err != nil {
		return c.Edit("Ошибка при загрузке времени")
	}

	kb := s.CreateTimeKeyboard(times)
	return c.Edit("Выберите время на которое хотите записаться", kb)
}

func (s *Service) HandleTimeSelection(c tele.Context, selectedTime string) error {
	s.state.Update(c.Chat().ID, selectedTime)
	s.state.Set(c.Chat().ID, StepInputNumber)
	return c.Edit("Введите свой номер телефона")
}

func (s *Service) HandleInputPhoneNumber(c tele.Context, phone string) error {
	s.state.Update(c.Chat().ID, phone)
	s.state.Set(c.Chat().ID, StepConfirm)
	return c.Send(fmt.Sprintf("Запись на стрижку\nДата: %s\nВремя: %s\nНомер телефона: %s", s.state.Get(c.Chat().ID).Date, s.state.Get(c.Chat().ID).Time, s.state.Get(c.Chat().ID).Number))
}

func (s *Service) HandleCancel(c tele.Context) error {
	s.state.Set(c.Chat().ID, StepNone)
	s.state.Clear(c.Chat().ID)
	return c.Send("Отмена")
}

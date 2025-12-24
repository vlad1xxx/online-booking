package service

import tele "gopkg.in/telebot.v4"

var (
	BtnBook    = tele.Btn{Text: "✂️ Записаться", Unique: "book_start"}
	BtnCancel  = tele.Btn{Text: "❌ Отмена", Unique: "cancel_booking"}
	BtnBack    = tele.Btn{Text: "◀️ Назад", Unique: "back_step"}
	BtnConfirm = tele.Btn{Text: "✅ Подтвердить", Unique: "confirm_booking"}
)

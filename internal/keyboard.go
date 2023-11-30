package internal

import "gopkg.in/telebot.v3"

var (
	startMenu       = &telebot.ReplyMarkup{ResizeKeyboard: true}
	aboutTheFund    = startMenu.Text("О фонде")
	preferences     = startMenu.Text("Преференции")
	investmentTerms = startMenu.Text("Условия инвестиций")
	rate            = startMenu.Text("Тариф")
	findOutMore     = startMenu.Text("Узнать ещё")

	linkMenu = &telebot.ReplyMarkup{}
	link     = linkMenu.URL("Посмотреть видео ", "https://drive.google.com/file/d/1fdIHa1E4CzWrFd4-FUGgll5gT6jec_HL/view?usp=sharing")
)

func (t *Telegram) getStartMenu() *telebot.ReplyMarkup {
	startMenu.Reply(
		startMenu.Row(aboutTheFund, preferences),
		startMenu.Row(investmentTerms, rate),
		startMenu.Row(findOutMore),
	)

	return startMenu
}

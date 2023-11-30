package internal

import "info-bot/config"

type App struct {
	Config *config.Config
}

func New(config *config.Config) *App {
	return &App{
		Config: config,
	}
}

func (a *App) Run() error {
	tg, err := NewBot(a.Config.Telegram)
	if err != nil {
		return err
	}

	tg.Start()

	return nil
}

package internal

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"info-bot/config"
	"time"
)

type Telegram struct {
	bot *telebot.Bot
}

func NewBot(config config.Telegram) (*Telegram, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token: config.Token,
		Poller: &telebot.LongPoller{
			Timeout: 50 * time.Millisecond,
		},
	})
	if err != nil {
		return nil, err
	}

	return &Telegram{
		bot: bot,
	}, nil
}

func (t *Telegram) Start() {
	t.bot.Handle("/start", t.handleStart)
	t.bot.Handle(&aboutTheFund, t.handleAboutTheFundBtn)
	t.bot.Handle(&preferences, t.handlePreferencesBtn)
	t.bot.Handle(&investmentTerms, t.handleInvestmentTermsBtn)
	t.bot.Handle(&rate, t.handleRateBtn)
	t.bot.Handle(&findOutMore, t.handleFindOutMoreBtn)
	t.bot.SetCommands([]telebot.Command{
		{
			Text:        "start",
			Description: "Запуск функционала",
		},
	})
	t.bot.Start()

}

func (t *Telegram) handleStart(ctx telebot.Context) error {
	return ctx.Send(&telebot.Video{
		File:     telebot.FromDisk("video/start.MP4"),
		Width:    0,
		Height:   0,
		Duration: 0,
		Caption: fmt.Sprintf("Здравствуйте, %s \nВас приветствует интеллектуальный помощник клиентской "+
			"поддержки корпорации EMIVN.\nВ этом боте вы сможете узнать краткую информацию о токене и проекте EMIVN.\n", ctx.Sender().Username),
		Thumbnail: nil,
		Streaming: false,
		MIME:      "",
		FileName:  "",
	}, t.getStartMenu())

	//linkMenu.Inline(linkMenu.Row(link))
	//
	//return ctx.Send("Интервью EMIVN", linkMenu)
}

func (t *Telegram) handleAboutTheFundBtn(ctx telebot.Context) error {
	ids := []int64{1046179822, 1093312909}

	for _, id := range ids {
		chat, err := t.bot.ChatByID(id)
		if err != nil {
			continue
		}

		if _, err := t.bot.Send(chat, "@"+ctx.Sender().Username+" Интересуется информацией о фонде"); err != nil {
			continue
		}
	}

	return ctx.Send(&telebot.Video{
		File:     telebot.FromDisk("video/answer_1.MP4"),
		Width:    0,
		Height:   0,
		Duration: 0,
		Caption: "Об инвестиционном фонде EMIVN\nИнвестиционный фонд EMIVN основан в 2019 году. С первых дней " +
			"работы, основываясь на принципе открытости, фонд стал проводить публичные мероприятия. Деятельность фонда стала " +
			"привлекательна для инвесторов за счет грамотной диверсификации инвестиций. Аудиторский отдел проводил ежедневную " +
			"работу по поиску наиболее привлекательных проектов, основным направлением для инвестирования являлись объекты " +
			"недвижимости. В процессе работы ИФ EMIVN заслужил высокую оценку инвесторов и, как следует, объёмы инвестиций " +
			"возросли более чем в 10 раз. ИФ EMIVN разработал крепкий, работающий и масштабируемый инвестиционный механизм. " +
			"Наша устойчивая и прочная операционная платформа, дисциплинированная мульти-факторная система по контролю над " +
			"рисками позволяет качественно увеличить доходность для инвесторов фонда.",
		Thumbnail: nil,
		Streaming: false,
		MIME:      "",
		FileName:  "",
	})
}

func (t *Telegram) handlePreferencesBtn(ctx telebot.Context) error {
	ids := []int64{1046179822, 1093312909}

	for _, id := range ids {
		chat, err := t.bot.ChatByID(id)
		if err != nil {
			continue
		}

		if _, err := t.bot.Send(chat, "@"+ctx.Sender().Username+" Интересуется информацией о преференциях"); err != nil {
			return err
		}
	}

	return ctx.Send(&telebot.Video{
		File:     telebot.FromDisk("video/answer_2.MP4"),
		Width:    0,
		Height:   0,
		Duration: 0,
		Caption: "- Получение пассивного дохода держателями инвестиционного токена EMIVN не единственное " +
			"преимущество. Фондом разработаны уникальные предложения для инвесторов, к которым относятся:\n\n- " +
			"Возможность одними из первых пользоваться VR-игрой EMIVN RACE, аналогов которой нет в мире;\n\n- Возможность " +
			"пользоваться продуктами EMIVN на специальных условиях;\n\n- Членство в бизнес-клубе EMIVN, участие в " +
			"мероприятиях клуба;\n\n- Возможность приобретения брендированных ювелирных изделий с логотипом EMIVN;\n\n- " +
			"Возможность покупки люксовых спорткаров у партнеров фонда EMIVN за токены EMIVN ниже рыночной стоимости до 20%;\n\n- " +
			"Возможность приобретения объектов недвижимого имущества у партнеров фонда EMIVN за токены EMIVN ниже рыночной " +
			"стоимости до 15%;\n\n- Возможность возведения объектов недвижимости разного класса за токены EMIVN на выгодных условиях.",
		Thumbnail: nil,
		Streaming: false,
		MIME:      "",
		FileName:  "",
	})
}

func (t *Telegram) handleInvestmentTermsBtn(ctx telebot.Context) error {
	ids := []int64{1046179822, 1093312909}

	for _, id := range ids {
		chat, err := t.bot.ChatByID(id)
		if err != nil {
			continue
		}
		if _, err := t.bot.Send(chat, "@"+ctx.Sender().Username+" Интересуется информацией об условиях инвестиций"); err != nil {
			return err
		}
	}

	return ctx.Send("Инвестиционный фонд EMIVN предоставляет понятные и доступные условия для инвестирования денежных средств. Инвесторы покупают токены фонда EMIVN (EMIVN). Приобретенные токены фонда EMIVN (EMIVN) направляются на стейкинг, в зависимости от выбранного пакета. Инвесторы получают пассивную доходность, которая, стоит отметить, в разы выше стандартных банковских предложений. Средства, полученные от продажи токенов EMIVN, направляются на стратегическое управление, то есть инвестируются в различные проекты для защиты активов и нивелирования рисков.")
}

func (t *Telegram) handleRateBtn(ctx telebot.Context) error {
	ids := []int64{1046179822, 1093312909}

	for _, id := range ids {
		chat, err := t.bot.ChatByID(id)
		if err != nil {
			continue
		}
		if _, err := t.bot.Send(chat, "@"+ctx.Sender().Username+" Интересуется информацией о тарифе"); err != nil {
			return err
		}
	}

	return ctx.Send("BRONZE  3 месяца 6%\nSILVER  6 месяцев 15%\nGOLD  12 месяцев  38%\nPLATINUM  24 месяца 100%")
}

func (t *Telegram) handleFindOutMoreBtn(ctx telebot.Context) error {
	ids := []int64{1046179822, 1093312909}

	for _, id := range ids {
		chat, err := t.bot.ChatByID(id)
		if err != nil {
			continue
		}

		if _, err := t.bot.Send(chat, "@"+ctx.Sender().Username+" потенциальный лид"); err != nil {
			return err
		}
	}

	return ctx.Send("Благодарим за повышенный интерес к нашему проекту, в скором времени к вам обратится менеджер, для продолжения обсуждения\nВремя работы клиентской поддержки с 10:00 до 22:00 по мск")
}

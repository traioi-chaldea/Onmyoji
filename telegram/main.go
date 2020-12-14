package telegram

import (
	"time"
	"fmt"
	vp "github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	filename = "telegram.yaml"
)

func Handler() {
	// Load yaml config
	vp.SetConfigName(filename)
	vp.SetConfigType("yaml")
	vp.AddConfigPath("config/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Get variables
	tURL     := vp.GetString("T_URL")
        //tChannel := vp.GetString("T_CHANNEL")
        //botName  := vp.GetString("BOT_NAME")
        botTOKEN := vp.GetString("BOT_TOKEN")

	// Telegram bot handler
	b, err := tb.NewBot(tb.Settings {
		URL: tURL,
		Token: botTOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(fmt.Errorf("Bot Telegram error: %s \n", err))
	}

	HandleBountyLocation(b)
	b.Handle(tb.OnText, func(m *tb.Message) {
		switch m.Text {
		case "/bountys":
			b.Send(m.Chat, "Test")
		}
	})


	b.Start()
}

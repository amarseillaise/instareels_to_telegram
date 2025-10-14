package bot

import (
	tele "gopkg.in/telebot.v4"
	"time"
)

func InitBot(token *string) (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  *token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(pref)
	return bot, err
}

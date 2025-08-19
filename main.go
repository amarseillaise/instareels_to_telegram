package main

import (
	"log"
	"regexp"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "7727195679:AAECHnq-R2tK6-Srejur0GeNpyZT15qQgCo",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(tele.OnText, func(c tele.Context) error {
		pattern := "\\.*instagram.com/reel\\.*/"
		_url := c.Text()
		is_valid_url, _ := regexp.MatchString(pattern, _url)
		if is_valid_url {
			DownloadReel(_url)
		}
		return nil
	})

	bot.Start()
}

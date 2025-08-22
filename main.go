package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	services "github.com/amarseillaise/instareels_to_telegram/services"

	env "github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func main() {
	initEnv()
	token := os.Getenv("TELETOKEN")
	pref := tele.Settings{
		Token:  token,
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
			res := services.DownloadReel(_url)
			fmt.Printf("res: %v\n", res)
		}
		return nil
	})

	bot.Start()
}

func initEnv() {
	if err := env.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

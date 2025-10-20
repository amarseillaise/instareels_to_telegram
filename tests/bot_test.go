package tests

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	env "github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/telebot.v4"
)

func TestSendVideo(t *testing.T) {

	// setup
	env.Load("../.env")
	token, ok := os.LookupEnv("TELETOKEN_TEST")
	if ok == false {
		log.Fatal("Test environment is not defined")
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	bot, _ := telebot.NewBot(pref)

	descriptionBytes, _ := os.ReadFile("./data/test.txt")
	descriptionContent := string(descriptionBytes)

	video := telebot.FromDisk("./data/test.mp4")
	teleVideo := &telebot.Video{File: video, Caption: descriptionContent}

	tele_id_env := os.Getenv("TELEGRAM_ID_TEST")
	tele_id, _ := strconv.ParseInt(tele_id_env, 10, 64)
	recepient := &telebot.Chat{ID: tele_id}

	t.Run("successful video send", func(t *testing.T) {
		_, err := bot.Send(recepient, teleVideo)
		assert.NoError(t, err)
	})
}

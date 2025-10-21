package bot

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

type ReelInfo struct {
	Video   *tele.Video
	Caption string
}

func InitBot(token *string) (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  *token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := tele.NewBot(pref)
	return bot, err
}

func MakeVideo(videoPath string) *tele.Video {
	teleVideo := &tele.Video{File: tele.FromDisk(videoPath)}
	return teleVideo
}

func MakeCaption(captionPath string) string {
	var captionContent string
	captionBytes, err := os.ReadFile(captionPath)
	if err == nil {
		captionContent = string(captionBytes)
		// trim to 1023 because of telegram limits
		runes := []rune(captionContent)
		if len(runes) >= 1024 {
			captionContent = string(runes[:1023])
		}
	}
	return captionContent
}

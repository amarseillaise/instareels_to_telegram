package bot

import (
	re "regexp"

	"gopkg.in/telebot.v4"

	"github.com/amarseillaise/instareels_to_telegram/services"
)

func OnTextHandler(c telebot.Context) error {
	pattern := "\\.*instagram.com/reel\\.*/"
	_url := c.Text()
	is_valid_url, _ := re.MatchString(pattern, _url)
	if is_valid_url {
		c.Notify(telebot.UploadingVideo)
		shortcode := services.ParseShortcode(_url)
		videoPath, captionPath, err := services.GetReelPath(shortcode)
		if err == nil {
			teleVideo := MakeVideo(videoPath)
			videoCaption := MakeCaption(captionPath)
			teleVideo.Caption = videoCaption
			return c.Reply(teleVideo)
		} else {
			return c.Reply("Error downloading reel")
		}
	}
	return nil
}

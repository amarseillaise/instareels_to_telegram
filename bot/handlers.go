package bot

import (
	re "regexp"

	tele "gopkg.in/telebot.v4"

	"github.com/amarseillaise/instareels_to_telegram/services"
)

func OnTextHandler(c tele.Context) error {
	pattern := "\\.*instagram.com/reel\\.*/"
	_url := c.Text()
	is_valid_url, _ := re.MatchString(pattern, _url)
	if is_valid_url {
		shortcode := services.ParseShortcode(_url)
		res, err := services.GetReel(shortcode)
		if err != nil {
			return c.Reply("Error downloading reel")
		}
		return c.Reply(res.Video)
	}
	return nil
}

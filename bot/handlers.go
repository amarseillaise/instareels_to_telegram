package bot

import (
	re "regexp"

	tele "gopkg.in/telebot.v4"

	"github.com/amarseillaise/instareels_to_telegram/adapters"
)

func OnTextHandler(c tele.Context) error {
	pattern := "\\.*instagram.com/reel\\.*/"
	_url := c.Text()
	is_valid_url, _ := re.MatchString(pattern, _url)
	if is_valid_url {
		shortcode := adapters.ParseShortcode(_url)
		res := adapters.DownloadReel(shortcode)
		if res != nil {
			return c.Reply("Error downloading reel")
		}
		return c.Reply("Here is your reel")
	}
	return nil
}

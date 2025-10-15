package services

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/telebot.v4"
)

const (
	tempDir = "temp"
)

type ReelInfo struct {
	Description string
	Video       telebot.Video
}

func GetReel(shortcode string) (ReelInfo, error) {
	reel := ReelInfo{}
	err := DownloadReel(shortcode)
	// TODO: make caching and try to read from cache first
	if err != nil {
		return reel, err
	}
	descriptionPath := findFile(shortcode, []string{".txt"})
	descriptionBytes, err := os.ReadFile(descriptionPath)
	if err != nil {
		return reel, err
	}
	descriptionContent := string(descriptionBytes)

	videoPath := findFile(shortcode, []string{".mp4", ".avi", ".mkv", ".mov"})
	video := telebot.FromDisk(videoPath)
	teleVideo := telebot.Video{File: video}

	reel.Description = descriptionContent
	reel.Video = teleVideo
	return reel, nil

}

func findFile(shortcode string, extensions []string) string {
	res := ""
	for _, ext := range extensions {
		mathces, err := filepath.Glob(fmt.Sprintf("%s/%s/*%s", tempDir, shortcode, ext))
		if err == nil && len(mathces) > 0 {
			res = mathces[0]
		}
	}
	return res
}

func ParseShortcode(_url string) string {
	pattern := "reel/.+/"
	re := regexp.MustCompile(pattern)
	match := re.FindString(_url)
	resultsSlice := strings.Split(match, "/")
	shortcode := resultsSlice[1]
	return fmt.Sprintf("-%s", shortcode)
}

package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var paths = struct {
	interpitator string
	script       string
	tempFiles    string
}{
	interpitator: "python",
	script:       "./instaloader/instaloader.py",
	tempFiles:    "temp/{shortcode}",
}

func DownloadReel(shortcode string) error {
	err := executeCMD(shortcode)
	return err
}

func ParseShortcode(_url string) string {
	pattern := "reel/.+/"
	re := regexp.MustCompile(pattern)
	match := re.FindString(_url)
	resultsSlice := strings.Split(match, "/")
	shortcode := resultsSlice[1]
	return fmt.Sprintf("-%s", shortcode)
}

func executeCMD(shortcode string) error {
	allArgs := getScriptArgs()
	scriptArgs := []string{paths.script}
	scriptArgs = append(scriptArgs, allArgs...)
	scriptArgs = append(scriptArgs, "--", shortcode)

	cmd := exec.Command(paths.interpitator, scriptArgs...)

	additionalEnv := "python=./.venv/Scripts/python"
	newEnv := append(os.Environ(), additionalEnv)
	cmd.Env = newEnv

	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing Python script: %v\nOutput: %s", err, res)
	} else {
		log.Print(string(res))
	}
	return err
}

func getScriptArgs() []string {
	return []string{
		// "--login", os.Getenv("INSTLOGIN"),
		// "--password", os.Getenv("INSTPASSWD"),
		"--dirname-pattern", paths.tempFiles,
		"--no-pictures",
		"--no-video-thumbnails",
		"--no-metadata-json",
		"--no-iphone",
		"--quiet",
	}
}

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
	interpitator: "./.venv/bin/python",
	script:       "./instaloader/instaloader.py",
	tempFiles:    "temp/{shortcode}",
}

func DownloadReel(_url string) error {
	shortcode := parseShortcode(_url)
	err := executeCMD(shortcode)
	return err
}

func parseShortcode(_url string) string {
	pattern := "reel/.+/"
	re := regexp.MustCompile(pattern)
	match := re.FindString(_url)
	resultsSlice := strings.Split(match, "/")
	shortcode := resultsSlice[1]
	return fmt.Sprintf("-- -%s", shortcode)
}

func executeCMD(shortcode string) error {
	scriptArgs := getScriptArgs()
	args := buildArgs(scriptArgs)
	cmd := exec.Command(paths.interpitator, paths.script, args, shortcode)
	res, err := cmd.CombinedOutput()
	fmt.Printf("err: %v\n", err)
	if err != nil {
		log.Fatalf("Error executing Python script: %v\nOutput: %s", err, res)
	} else {
		log.Print(string(res))
	}
	return err
}

func getScriptArgs() map[string]bool {
	var scriptArgs = map[string]bool{
		fmt.Sprintf("--login %s", os.Getenv("INSTLOGIN")):     true,
		fmt.Sprintf("--password %s", os.Getenv("INSTPASSWD")): true,
		fmt.Sprintf("--dirname-pattern %s", paths.tempFiles):  true,
		"--no-pictures":         true,
		"--no-video-thumbnails": true,
		"--no-metadata-json":    true,
		"--no-iphone":           true,
		"--quiet":               true,
	}
	return scriptArgs
}

func buildArgs(args map[string]bool) string {
	var resArgs []string
	for arg, needed := range args {
		if needed {
			resArgs = append(resArgs, arg)
		}
	}
	res := strings.Join(resArgs, " ")
	return res
}

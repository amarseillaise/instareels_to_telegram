package services

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

type Paths struct {
	interpitator string
	script       string
	tempFiles    string
}

func initPaths() Paths {
	paths := Paths{
		interpitator: "",
		script:       "./instaloader/instaloader.py",
		tempFiles:    fmt.Sprintf("%s/{shortcode}", tempDir),
	}

	switch os := runtime.GOOS; os {
	case "windows":
		paths.interpitator = "./.venv/Scripts/python"
	default:
		paths.interpitator = "./.venv/bin/python"
	}
	return paths
}

var paths = initPaths()

func DownloadReel(shortcode string) error {
	err := executeCMD(shortcode)
	return err
}

func executeCMD(shortcode string) error {
	allArgs := getScriptArgs()
	scriptArgs := []string{paths.script}
	scriptArgs = append(scriptArgs, allArgs...)
	scriptArgs = append(scriptArgs, "--", fmt.Sprintf("-%s", shortcode))

	cmd := exec.Command(paths.interpitator, scriptArgs...)

	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing Python script: %v\nOutput: %s", err, res)
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

package main

import (
	"log"
	"os/exec"
)

func getSriptPath() string {
	return "./instaloader/instaloader.py"
}

func executeCMD(scriptPath string) error {
	cmd := exec.Command("python", scriptPath)
	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing Python script: %v\nOutput: %s", err, res)
	} else {
		log.Print(res)
	}
	return err
}

func DownloadReel(_url string) error {
	p := getSriptPath()
	err := executeCMD(p)
	return err
}

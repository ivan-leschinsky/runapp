package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"errors"
	"strings"
	"runtime"
)

func exit(err error) {
	log.Fatalln(err)
	os.Exit(1)
}

func runApp(appPath string, ShowNotifications bool) {
	exec.Command("open", appPath).Output()
	if ShowNotifications {
		showNotification(fmt.Sprintf("App %s started", appPath), "Go App runner")
	}
}
func showNotification(text string, title string) {
	if runtime.GOOS == "darwin" {
		cmd := fmt.Sprintf("osascript -e 'display notification \"%s\" with title \"%s\"'", text, title)
		exec.Command("bash","-c",cmd).Output()
	}
}
func checkApp(appPath string, ShowNotifications bool) {
	cmd := fmt.Sprintf("ps aux | grep \"%s\" | grep -v grep", appPath)
	out, err := exec.Command("bash","-c",cmd).Output()
	if err != nil {
		if err.Error() == "exit status 1" {
			runApp(appPath, ShowNotifications)
			return;
		}
	}
	if !strings.Contains(string(out), appPath) {
		runApp(appPath, ShowNotifications)
	}
}

func main() {
	log.SetFlags(0)

	cfg := loadConfig(fmt.Sprintf("%s/.runapp.yml", os.Getenv("HOME")))
	// cfg := loadConfig("./runapp.yml")
	if cfg == nil {
		exit(errors.New("No valid config found"))
	}

	for _, appPath := range cfg.Apps {
		checkApp(appPath, cfg.ShowNotifications)
	}
}

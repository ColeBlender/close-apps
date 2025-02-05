package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	scriptList := `
		tell application "System Events"
			get displayed name of every process whose background only is false
		end tell
	`
	output, err := exec.Command("osascript", "-e", scriptList).Output()
	if err != nil {
		fmt.Println("Error fetching running applications:", err)
		return
	}

	apps := strings.Split(strings.TrimSpace(string(output)), ", ")

	for _, app := range apps {
		if app == "iTerm2" {
			continue
		}
		quitApp(app)
	}

	emptyTrash()

	time.Sleep(500 * time.Millisecond)

	quitApp("iTerm2")
}

func quitApp(app string) {
	scriptQuit := fmt.Sprintf(`tell application "%s" to quit`, app)
	_, quitErr := exec.Command("osascript", "-e", scriptQuit).Output()
	if quitErr != nil {
		fmt.Printf("Error quitting %s: %v\n", app, quitErr)
	}
}

func emptyTrash() {
	scriptEmptyTrash := `tell application "Finder" to empty the trash`
	_, err := exec.Command("osascript", "-e", scriptEmptyTrash).Output()
	if err != nil {
		fmt.Println("Error emptying the trash:", err)
	}
}

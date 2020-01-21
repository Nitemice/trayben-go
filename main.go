package main

import (
	"os"
	"strings"

	"github.com/getlantern/systray"
	"./icon"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)

	// Retrieve the message, and set it to the tooltip and menu item text
	msg := strings.Join(os.Args[1:], " ")
	systray.SetTooltip(msg)
	systray.AddMenuItem(msg, "")

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Close Trayben")
	mQuit.SetIcon(icon.Data)

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

}

func onExit() {
}

package main

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func onTrayReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("ytb2mpv")
	systray.SetTooltip("ytb2mpv is running")
	openSettingUI := systray.AddMenuItem("Open Setting", "Open Setting to customize path and startup options")
	quitDaemon := systray.AddMenuItem("Quit", "Quit ytb2mpv")

	go func() {
		for {
			select {
			case <-openSettingUI.ClickedCh:
				ShowSettingGUI()
			case <-quitDaemon.ClickedCh:
				systray.Quit()
				os.Exit(0)

			}
		}
	}()
}

func RunTray() {
	systray.Run(onTrayReady, func() {
		os.Exit(0)
	})
}

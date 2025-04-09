package main

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/pkg/browser"
)

func onTrayReady() {
	systray.SetIcon(MainIconData)
	systray.SetTitle("ytb2mpv")
	systray.SetTooltip("ytb2mpv is running")
	ytb2mpvInfo := systray.AddMenuItem("ytb2mpv daemon v"+VERSION, "ytb2mpv daemon v"+VERSION)
	ytb2mpvInfo.SetIcon(MainIconData)
	systray.AddSeparator()
	checkForUpdate := systray.AddMenuItem("Check for update", "Check for update")
	startOnLoginCheckbox := systray.AddMenuItemCheckbox("Start on login", "Start ytb2mpv on login", true)
	openSettingUI := systray.AddMenuItem("Open Setting", "Open Setting to customize path and startup options")
	quitDaemon := systray.AddMenuItem("Quit", "Quit ytb2mpv")

	go func() {
		for {
			select {
			case <-startOnLoginCheckbox.ClickedCh:
				if startOnLoginCheckbox.Checked() {
					startOnLoginCheckbox.Uncheck()
				} else {
					startOnLoginCheckbox.Check()
				}
			case <-checkForUpdate.ClickedCh:
				SendNotify("ytb2mpv", "Checking for update...", false)
			case <-ytb2mpvInfo.ClickedCh:
				browser.OpenURL("https://github.com/michioxd/ytb2mpv")
			case <-openSettingUI.ClickedCh:
				ShowSettingGUI()
			case <-quitDaemon.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func RunTray() {
	systray.Run(onTrayReady, func() {
		os.Exit(0)
	})
}

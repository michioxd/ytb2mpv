package main

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/pkg/browser"
	"github.com/spf13/viper"
)

func onTrayReady() {
	isDebug := ""

	if RELEASE_MODE != "1" {
		isDebug = " (Debug)"
	}

	systray.SetIcon(MainIconData)
	systray.SetTitle("ytb2mpv")
	systray.SetTooltip("ytb2mpv is running")
	if RELEASE_MODE != "1" {
		systray.AddMenuItem("Running in Debug mode", "").Disable()
		systray.AddMenuItem("Start on login has been disabled!", "").Disable()
		systray.AddSeparator()
	}
	ytb2mpvInfo := systray.AddMenuItem("ytb2mpv daemon v"+VERSION+isDebug, "ytb2mpv daemon v"+VERSION+isDebug)
	ytb2mpvInfo.SetIcon(MainIconData)
	systray.AddSeparator()
	checkForUpdate := systray.AddMenuItem("Check for update", "Check for update")
	startOnLoginCheckbox := systray.AddMenuItemCheckbox("Start on login", "Start ytb2mpv on login", viper.GetBool("start_w_system"))
	openSettingUI := systray.AddMenuItem("Open Setting", "Open Setting to customize path and startup options")
	quitDaemon := systray.AddMenuItem("Quit", "Quit ytb2mpv")

	go func() {
		for {
			select {
			case <-startOnLoginCheckbox.ClickedCh:
				if startOnLoginCheckbox.Checked() {
					startOnLoginCheckbox.Uncheck()
					viper.Set("start_w_system", false)
					RegisterStartup(false)
				} else {
					startOnLoginCheckbox.Check()
					viper.Set("start_w_system", true)
					RegisterStartup(true)
				}
				if err := viper.WriteConfig(); err != nil {
					SendNotify("Error", "Failed to save config file: "+err.Error(), true)
				}
			case <-checkForUpdate.ClickedCh:
				SendNotify("ytb2mpv", "Checking for update...", false)
			case <-ytb2mpvInfo.ClickedCh:
				browser.OpenURL("https://github.com/michioxd/ytb2mpv")
			case <-openSettingUI.ClickedCh:
				ShowSettingGUI(func(b bool) {
					if b {
						startOnLoginCheckbox.Check()
					} else {
						startOnLoginCheckbox.Uncheck()
					}
				})
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

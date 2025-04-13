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
		addDebugModeIndicators()
	}

	ytb2mpvInfo := createAppInfoMenuItem(isDebug)
	mpvInfo, ytDlpInfo := createDependencyInfoMenuItems()

	systray.AddSeparator()
	openYtDlpDownloader := systray.AddMenuItem("Open yt-dlp downloader UI", "Open yt-dlp downloader UI")
	checkForUpdate := systray.AddMenuItem("Check for update", "Check for update")
	startOnLoginCheckbox := systray.AddMenuItemCheckbox(
		"Start on login",
		"Start ytb2mpv on login",
		viper.GetBool("start_w_system"),
	)
	openSettingUI := systray.AddMenuItem("Open Setting", "Open Setting to customize path and startup options")
	quitDaemon := systray.AddMenuItem("Quit", "Quit ytb2mpv")

	go handleMenuItemClicks(
		openYtDlpDownloader,
		startOnLoginCheckbox,
		checkForUpdate,
		ytb2mpvInfo,
		openSettingUI,
		quitDaemon,
		mpvInfo,
		ytDlpInfo,
	)
}

func handleMenuItemClicks(
	openYtDlpDownloader *systray.MenuItem,
	startOnLoginCheckbox *systray.MenuItem,
	checkForUpdate *systray.MenuItem,
	ytb2mpvInfo *systray.MenuItem,
	openSettingUI *systray.MenuItem,
	quitDaemon *systray.MenuItem,
	mpvInfo *systray.MenuItem,
	ytDlpInfo *systray.MenuItem,
) {
	for {
		select {
		case <-openYtDlpDownloader.ClickedCh:
			DownloaderGUI("")

		case <-startOnLoginCheckbox.ClickedCh:
			handleStartOnLogin(startOnLoginCheckbox)

		case <-checkForUpdate.ClickedCh:
			SendNotify("ytb2mpv", "Checking for update...", false)

		case <-ytb2mpvInfo.ClickedCh:
			browser.OpenURL("https://github.com/michioxd/ytb2mpv")

		case <-openSettingUI.ClickedCh:
			ShowSettingGUI(func(b bool) {
				updateStartOnLoginStatus(b, startOnLoginCheckbox)
				updateDependencyInfo(mpvInfo, ytDlpInfo)
			})

		case <-quitDaemon.ClickedCh:
			systray.Quit()
		}
	}
}

func addDebugModeIndicators() {
	systray.AddMenuItem("Running in Debug mode", "").Disable()
	systray.AddMenuItem("Start on login has been disabled!", "").Disable()
	systray.AddSeparator()
}

func createAppInfoMenuItem(isDebug string) *systray.MenuItem {
	ytb2mpvInfo := systray.AddMenuItem(
		"ytb2mpv daemon v"+VERSION+isDebug,
		"ytb2mpv daemon v"+VERSION+isDebug,
	)
	ytb2mpvInfo.SetIcon(MainIconData)
	return ytb2mpvInfo
}

func createDependencyInfoMenuItems() (*systray.MenuItem, *systray.MenuItem) {
	mpvInfo := systray.AddMenuItem("mpv info", "mpv info")
	ytDlpInfo := systray.AddMenuItem("yt-dlp info", "yt-dlp info")

	mpvInfo.Disable()
	ytDlpInfo.Disable()

	updateDependencyInfo(mpvInfo, ytDlpInfo)

	return mpvInfo, ytDlpInfo
}

func handleStartOnLogin(checkbox *systray.MenuItem) {
	enabled := !checkbox.Checked()

	if enabled {
		checkbox.Check()
	} else {
		checkbox.Uncheck()
	}

	viper.Set("start_w_system", enabled)
	RegisterStartup(enabled)

	if err := viper.WriteConfig(); err != nil {
		SendNotify("Error", "Failed to save config file: "+err.Error(), true)
	}
}

func updateStartOnLoginStatus(enabled bool, checkbox *systray.MenuItem) {
	if enabled {
		checkbox.Check()
	} else {
		checkbox.Uncheck()
	}
}

func updateDependencyInfo(mpvInfo, ytDlpInfo *systray.MenuItem) {
	if MPV_STATUS == 0 {
		mpvInfo.SetTitle("mpv: " + MPV_VERSION)
	} else {
		mpvInfo.SetTitle("mpv: Not Found")
	}

	if YTDLP_STATUS == 0 {
		ytDlpInfo.SetTitle("yt-dlp: " + YTDLP_VERSION)
	} else {
		ytDlpInfo.SetTitle("yt-dlp: Not Found")
	}
}

func RunTray() {
	systray.Run(onTrayReady, func() {
		os.Exit(0)
	})
}

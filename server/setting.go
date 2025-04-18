package main

import (
	"strings"

	"github.com/gonutz/wui/v2"
	"github.com/spf13/viper"
)

func ShowSettingGUI(onUpdateSetting func(bool)) {
	icon, _ := wui.NewIconFromExeResource(2)
	settingWindowFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Segoe UI",
		Height: -11,
	})

	settingWindow := wui.NewWindow()
	settingWindow.SetFont(settingWindowFont)
	settingWindow.SetInnerSize(410, 225)
	settingWindow.SetPosition(100, 100)
	settingWindow.SetTitle("ytb2mpv setting")
	settingWindow.SetHasMinButton(false)
	settingWindow.SetHasMaxButton(false)
	settingWindow.SetResizable(false)
	settingWindow.SetIcon(icon)

	label1 := wui.NewLabel()
	label1.SetBounds(15, 15, 150, 13)
	label1.SetText("Path to mpv.exe")
	settingWindow.Add(label1)

	pathMpvEdit := wui.NewEditLine()
	pathMpvEdit.SetBounds(15, 34, 295, 20)
	pathMpvEdit.SetText(viper.GetString("path_mpv"))
	settingWindow.Add(pathMpvEdit)

	pathMpvEditBtn := wui.NewButton()
	pathMpvEditBtn.SetBounds(315, 33, 80, 22)
	pathMpvEditBtn.SetText("Browse")
	pathMpvEditBtn.SetOnClick(func() {
		open := wui.NewFileOpenDialog()
		open.SetInitialPath("C:\\")
		open.SetTitle("Select mpv.exe executable")
		open.AddFilter("Executable file", ".exe")
		if accept, path := open.ExecuteSingleSelection(settingWindow); accept {
			pathMpvEdit.SetText(path)
		}
	})
	settingWindow.Add(pathMpvEditBtn)

	label2 := wui.NewLabel()
	label2.SetBounds(15, 65, 150, 13)
	label2.SetText("Path to yt-dlp.exe")
	settingWindow.Add(label2)

	pathYtDlpEdit := wui.NewEditLine()
	pathYtDlpEdit.SetBounds(15, 84, 295, 20)
	pathYtDlpEdit.SetText(viper.GetString("path_ytdlp"))
	settingWindow.Add(pathYtDlpEdit)

	pathYtDlpEditBtn := wui.NewButton()
	pathYtDlpEditBtn.SetBounds(315, 83, 80, 22)
	pathYtDlpEditBtn.SetText("Browse")
	pathYtDlpEditBtn.SetOnClick(func() {
		open := wui.NewFileOpenDialog()
		open.SetInitialPath("C:\\")
		open.SetTitle("Select yt-dlp.exe executable")
		open.AddFilter("Executable file", ".exe")
		if accept, path := open.ExecuteSingleSelection(settingWindow); accept {
			pathYtDlpEdit.SetText(path)
		}
	})
	settingWindow.Add(pathYtDlpEditBtn)

	exitBtn := wui.NewButton()
	exitBtn.SetBounds(335, 185, 60, 25)
	exitBtn.SetText("Exit")
	exitBtn.SetOnClick(func() {
		settingWindow.Close()
	})
	settingWindow.Add(exitBtn)

	label3 := wui.NewLabel()
	label3.SetEnabled(false)
	label3.SetBounds(15, 190, 200, 13)
	label3.SetText("https://github.com/michioxd/ytb2mpv")
	settingWindow.Add(label3)

	startOnLogin := wui.NewCheckBox()
	startOnLogin.SetBounds(15, 167, 121, 17)
	startOnLogin.SetText("Start on Sign in")
	startOnLogin.SetChecked(viper.GetBool("start_w_system"))
	settingWindow.Add(startOnLogin)

	autoCheckUpdate := wui.NewCheckBox()
	autoCheckUpdate.SetBounds(130, 167, 180, 17)
	autoCheckUpdate.SetText("Auto check for update")
	autoCheckUpdate.SetChecked(viper.GetBool("auto_check_update"))
	settingWindow.Add(autoCheckUpdate)

	customMpvArgEdit := wui.NewEditLine()
	customMpvArgEdit.SetBounds(15, 144, 380, 20)
	customMpvArgEdit.SetText(viper.GetString("mpv_args"))
	settingWindow.Add(customMpvArgEdit)

	saveBtn := wui.NewButton()
	saveBtn.SetBounds(270, 185, 60, 25)
	saveBtn.SetText("Save")
	saveBtn.SetOnClick(func() {
		if strings.Contains(customMpvArgEdit.Text(), "ytdl_hook-ytdl_path=") {
			wui.MessageBoxError("Error", "Please remove ytdl_hook-ytdl_path= from custom mpv arguments")
			return
		}

		saveBtn.SetEnabled(false)
		pathMpvEdit.SetEnabled(false)
		pathYtDlpEdit.SetEnabled(false)
		pathMpvEditBtn.SetEnabled(false)
		pathYtDlpEditBtn.SetEnabled(false)
		startOnLogin.SetEnabled(false)
		autoCheckUpdate.SetEnabled(false)
		exitBtn.SetEnabled(false)
		customMpvArgEdit.SetEnabled(false)
		saveBtn.SetText("Saving...")

		go func() {
			viper.Set("path_mpv", pathMpvEdit.Text())
			viper.Set("path_ytdlp", pathYtDlpEdit.Text())
			viper.Set("mpv_args", customMpvArgEdit.Text())
			viper.Set("start_w_system", startOnLogin.Checked())
			viper.Set("auto_check_update", autoCheckUpdate.Checked())
			if err := viper.WriteConfig(); err != nil {
				wui.MessageBoxError("Error", "Failed to save config file: "+err.Error())
			}
			CheckEnv()
			RegisterStartup(startOnLogin.Checked())
			onUpdateSetting(startOnLogin.Checked())
			settingWindow.Close()
		}()
	})
	settingWindow.Add(saveBtn)

	label4 := wui.NewLabel()
	label4.SetEnabled(false)
	label4.SetBounds(15, 108, 379, 13)
	label4.SetText("If unspecified, ytb2mpv will use the system PATH.")
	settingWindow.Add(label4)

	label5 := wui.NewLabel()
	label5.SetBounds(15, 128, 150, 13)
	label5.SetText("Custom mpv arguments")
	settingWindow.Add(label5)

	settingWindow.Show()
}

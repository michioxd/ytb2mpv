package main

import "github.com/gonutz/wui/v2"

func ShowSettingGUI() {
	settingWindowFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Segoe UI",
		Height: -11,
	})

	settingWindow := wui.NewWindow()
	settingWindow.SetFont(settingWindowFont)
	settingWindow.SetInnerSize(410, 187)
	settingWindow.SetTitle("ytb2mpv setting")
	settingWindow.SetHasMinButton(false)
	settingWindow.SetHasMaxButton(false)
	settingWindow.SetResizable(false)

	label1 := wui.NewLabel()
	label1.SetBounds(15, 15, 150, 13)
	label1.SetText("Path to mpv.exe")
	settingWindow.Add(label1)

	pathMpvEdit := wui.NewEditLine()
	pathMpvEdit.SetBounds(15, 34, 295, 20)
	settingWindow.Add(pathMpvEdit)

	pathMpvEditBtn := wui.NewButton()
	pathMpvEditBtn.SetBounds(315, 33, 80, 22)
	pathMpvEditBtn.SetText("Browse")
	settingWindow.Add(pathMpvEditBtn)

	label2 := wui.NewLabel()
	label2.SetBounds(15, 65, 150, 13)
	label2.SetText("Path to yt-dlp.exe")
	settingWindow.Add(label2)

	pathYtDlpEdit := wui.NewEditLine()
	pathYtDlpEdit.SetBounds(15, 84, 295, 20)
	settingWindow.Add(pathYtDlpEdit)

	pathYtDlpEditBtn := wui.NewButton()
	pathYtDlpEditBtn.SetBounds(315, 83, 80, 22)
	pathYtDlpEditBtn.SetText("Browse")
	settingWindow.Add(pathYtDlpEditBtn)

	exitBtn := wui.NewButton()
	exitBtn.SetBounds(335, 145, 60, 25)
	exitBtn.SetText("Exit")
	exitBtn.SetOnClick(func() {
		settingWindow.Close()
	})
	settingWindow.Add(exitBtn)

	saveBtn := wui.NewButton()
	saveBtn.SetBounds(270, 145, 60, 25)
	saveBtn.SetText("Save")
	settingWindow.Add(saveBtn)

	label3 := wui.NewLabel()
	label3.SetEnabled(false)
	label3.SetBounds(15, 150, 200, 13)
	label3.SetText("https://github.com/michioxd/ytb2mpv")
	settingWindow.Add(label3)

	startOnLogin := wui.NewCheckBox()
	startOnLogin.SetBounds(15, 127, 121, 17)
	startOnLogin.SetText("Start on Sign in")
	startOnLogin.SetChecked(true)
	settingWindow.Add(startOnLogin)

	label4Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Segoe UI",
		Height: -11,
	})

	label4 := wui.NewLabel()
	label4.SetFont(label4Font)
	label4.SetEnabled(false)
	label4.SetBounds(15, 108, 379, 13)
	label4.SetText("If unspecified, ytb2mpv will use the system PATH.")
	settingWindow.Add(label4)

	settingWindow.Show()
}

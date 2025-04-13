package main

import "github.com/gonutz/wui/v2"

func DownloaderGUI(initialUrl string) {
	if YTDLP_STATUS > 0 {
		wui.MessageBoxError("yt-dlp not found", "yt-dlp not found, please check your setting")
		return
	}
	icon, _ := wui.NewIconFromExeResource(2)
	downloaderFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Segoe UI",
		Height: -11,
	})

	downloader := wui.NewWindow()
	downloader.SetFont(downloaderFont)
	downloader.SetInnerSize(440, 255)
	downloader.SetTitle("ytb2mpv yt-dlp downloader UI")
	downloader.SetHasMaxButton(false)
	downloader.SetResizable(false)
	downloader.SetIcon(icon)

	urlEdit := wui.NewEditLine()
	urlEdit.SetBounds(15, 34, 360, 20)
	downloader.Add(urlEdit)
	if initialUrl != "" {
		urlEdit.SetText(initialUrl)
	}

	label1 := wui.NewLabel()
	label1.SetBounds(15, 15, 150, 13)
	label1.SetText("YouTube Video URL:")
	downloader.Add(label1)

	getData := wui.NewButton()
	getData.SetBounds(380, 33, 46, 22)
	getData.SetText("Get")
	downloader.Add(getData)

	label2 := wui.NewLabel()
	label2.SetBounds(15, 65, 150, 13)
	label2.SetText("Title:")
	downloader.Add(label2)

	videoTitle := wui.NewEditLine()
	videoTitle.SetBounds(45, 62, 380, 20)
	downloader.Add(videoTitle)

	label3 := wui.NewLabel()
	label3.SetBounds(15, 92, 150, 13)
	label3.SetText("Video formats:")
	downloader.Add(label3)

	videoFormats := wui.NewComboBox()
	videoFormats.SetBounds(15, 110, 200, 21)
	videoFormats.SetItems([]string{})
	videoFormats.SetSelectedIndex(0)
	downloader.Add(videoFormats)

	label4 := wui.NewLabel()
	label4.SetBounds(224, 92, 150, 13)
	label4.SetText("Audio formats:")
	downloader.Add(label4)

	audioFormats := wui.NewComboBox()
	audioFormats.SetBounds(225, 110, 200, 21)
	audioFormats.SetItems([]string{})
	audioFormats.SetSelectedIndex(0)
	downloader.Add(audioFormats)

	bestVideoBox := wui.NewCheckBox()
	bestVideoBox.SetBounds(15, 135, 70, 17)
	bestVideoBox.SetText("Best video")
	bestVideoBox.SetChecked(true)
	downloader.Add(bestVideoBox)

	bestAudioBox := wui.NewCheckBox()
	bestAudioBox.SetBounds(224, 135, 73, 17)
	bestAudioBox.SetText("Best audio")
	bestAudioBox.SetChecked(true)
	downloader.Add(bestAudioBox)

	noVideoBox := wui.NewCheckBox()
	noVideoBox.SetBounds(95, 135, 70, 17)
	noVideoBox.SetText("No Video")
	downloader.Add(noVideoBox)

	noAudioBox := wui.NewCheckBox()
	noAudioBox.SetBounds(305, 135, 70, 17)
	noAudioBox.SetText("No Audio")
	downloader.Add(noAudioBox)

	progressBar := wui.NewProgressBar()
	progressBar.SetBounds(15, 160, 410, 30)
	progressBar.SetValue(0.5)
	downloader.Add(progressBar)

	quitBtn := wui.NewButton()
	quitBtn.SetBounds(341, 215, 85, 25)
	quitBtn.SetText("Quit")
	quitBtn.SetOnClick(func() {
		downloader.Close()
	})
	downloader.Add(quitBtn)

	downloadBtn := wui.NewButton()
	downloadBtn.SetBounds(252, 215, 85, 25)
	downloadBtn.SetText("Download")
	downloader.Add(downloadBtn)

	label6 := wui.NewLabel()
	label6.SetEnabled(false)
	label6.SetBounds(15, 221, 235, 13)
	label6.SetText("https://github.com/michioxd/ytb2mpv")
	downloader.Add(label6)

	ytdlpVersion := wui.NewLabel()
	ytdlpVersion.SetEnabled(false)
	ytdlpVersion.SetBounds(225, 15, 200, 13)
	ytdlpVersion.SetText("yt-dlp version: " + YTDLP_VERSION)
	ytdlpVersion.SetAlignment(wui.AlignRight)
	downloader.Add(ytdlpVersion)

	notifyAfterOk := wui.NewCheckBox()
	notifyAfterOk.SetBounds(15, 195, 190, 17)
	notifyAfterOk.SetText("Notify after download complete")
	notifyAfterOk.SetChecked(true)
	downloader.Add(notifyAfterOk)

	downloader.Show()
}

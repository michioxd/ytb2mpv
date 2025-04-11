package main

import (
	"fmt"
	"os"

	"github.com/gonutz/wui/v2"
	"github.com/spf13/viper"
)

func CheckEnv() {
	pathMpv := viper.GetString("path_mpv")
	pathYtdlp := viper.GetString("path_ytdlp")

	MPV_STATUS = CheckMPV(pathMpv)
	YTDLP_STATUS = CheckYTDLP(pathYtdlp)

	switch MPV_STATUS {
	case 1:
		SendNotify("(ytb2mpv) mpv Not Found", "mpv not found, please install mpv then add it into PATH or manually set the path in the setting", true)
	case 2:
		SendNotify("(ytb2mpv) mpv Error", "Cannot execute mpv, make sure mpv is vaild and not corrupted", true)
	case 3:
		SendNotify("(ytb2mpv) mpv Version Error", "mpv version is too old, please update mpv to the latest version", true)
	default:
		if viper.GetString("path_mpv") != "" {
			MPV_PATH = viper.GetString("path_mpv")
		}
	}

	switch YTDLP_STATUS {
	case 1:
		SendNotify("(ytb2mpv) yt-dlp Not Found", "yt-dlp not found, please install yt-dlp then add it into PATH or manually set the path in the setting", true)
	case 2:
		SendNotify("(ytb2mpv) yt-dlp Error", "Cannot execute yt-dlp, make sure yt-dlp is vaild and not corrupted", true)
	case 3:
		SendNotify("(ytb2mpv) yt-dlp Version Error", "yt-dlp version is too old, please update yt-dlp to the latest version", true)

	default:
		if viper.GetString("path_ytdlp") != "" {
			YTDLP_PATH = viper.GetString("path_ytdlp")
		}
	}
}

func InitCfg() {
	exePath, _ := os.Executable()
	APP_EXECUTABLE_PATH = exePath
	configDir, _ := os.UserConfigDir()
	viper.SetDefault("path_mpv", "")
	viper.SetDefault("path_ytdlp", "")
	viper.SetDefault("auto_check_update", true)
	viper.SetDefault("start_w_system", true)
	viper.SetDefault("mpv_args", "--ytdl-format=bestvideo+bestaudio --keep-open")

	viper.SetConfigName("ytb2mpv")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir + "\\ytb2mpv")

	if err := os.MkdirAll(configDir+"\\ytb2mpv", os.ModePerm); err != nil {
		wui.MessageBoxError("Error", "Failed to create config directory: "+err.Error())
	}

	if cfgErr := viper.ReadInConfig(); cfgErr != nil {
		if _, ok := cfgErr.(viper.ConfigFileNotFoundError); ok {
			viper.WriteConfigAs(configDir + "\\ytb2mpv\\ytb2mpv.yaml")
			fmt.Println("Config file not found, created default config file")
		} else {
			wui.MessageBoxError("Error", "Failed to read config file: "+cfgErr.Error())
		}
	}

	CheckEnv()
	RegisterStartup(viper.GetBool("start_w_system"))
}

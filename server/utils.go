package main

import (
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/gen2brain/beeep"
)

func SendNotify(title, message string, isAlert bool) {
	if isAlert {
		err := beeep.Alert(title, message, "assets/information.png")
		if err != nil {
			panic(err)
		}
		return
	}

	err := beeep.Notify(title, message, "assets/information.png")
	if err != nil {
		panic(err)
	}
}

func compareVersions(found, required []int) bool {
	for i := 0; i < len(required); i++ {
		if i >= len(found) {
			return false
		}
		if found[i] > required[i] {
			return true
		}
		if found[i] < required[i] {
			return false
		}
	}
	return true
}

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func parseVersion(version string) []int {
	parts := regexp.MustCompile(`\d+`).FindAllString(version, -1)
	nums := make([]int, len(parts))
	for i, p := range parts {
		nums[i], _ = strconv.Atoi(p)
	}
	return nums
}

func ExecCmd(cmd string, args ...string) *exec.Cmd {
	out := exec.Command(cmd, args...)
	out.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return out
}

func CheckMPV(path ...string) int {
	mpvPath := "mpv"
	if len(path) > 0 && path[0] != "" {
		mpvPath = filepath.Clean(path[0])
	}

	if !CommandExists(mpvPath) {
		return 1
	}

	out, err := ExecCmd(mpvPath, "--version").Output()
	if err != nil {
		return 2
	}

	firstLine := strings.Split(string(out), "\n")[0]
	re := regexp.MustCompile(`mpv\s+(\d+\.\d+\.\d+)`)
	match := re.FindStringSubmatch(firstLine)
	if len(match) < 2 {
		return 3
	}

	version := parseVersion(match[1])
	required := parseVersion("0.39.0")

	if compareVersions(version, required) {
		return 0
	}

	return 4
}

func CheckYTDLP(path ...string) int {
	ytDlpPath := "yt-dlp"
	if len(path) > 0 && path[0] != "" {
		ytDlpPath = filepath.Clean(path[0])
	}

	if !CommandExists(ytDlpPath) {
		return 1
	}

	out, err := ExecCmd(ytDlpPath, "--version").Output()
	if err != nil {
		return 2
	}

	versionStr := strings.TrimSpace(string(out))
	version := parseVersion(versionStr)
	required := parseVersion("2025.02.19")

	if compareVersions(version, required) {
		return 0
	}
	return 3
}

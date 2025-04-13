package ytdlp

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"syscall"

	"github.com/goccy/go-json"
)

func execCmd(cmd string, args ...string) *exec.Cmd {
	out := exec.Command(cmd, args...)
	out.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return out
}

func GetInfo(ytdlpCmd string, url string) (*YtdlpOutput, error) {
	cmd := execCmd(ytdlpCmd, "--no-warnings", "-q", "-O", "%(id)s,%(title)s,%(formats)s#j", url)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	if stderr.Len() > 0 {
		return nil, errors.New(stderr.String())
	}

	outStr := strings.TrimSpace(stdout.String())
	start := strings.Index(outStr, "{")
	if start == -1 {
		return nil, errors.New("video not existed or no JSON found in output")
	}

	jsonPart := outStr[start:]

	var data YtdlpOutput
	if err := json.Unmarshal([]byte(jsonPart), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

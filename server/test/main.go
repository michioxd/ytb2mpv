package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func execCmd(cmd string, args ...string) *exec.Cmd {
	out := exec.Command(cmd, args...)
	out.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return out
}

func main() {
	cmd := execCmd("yt-dlp", "-O", "%(.{id,title,formats})#j", "https://www.youtube.com/watch?v=jWvuUeUyyKU")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("yt-dlp failed:", err)
	}

	outStr := strings.TrimSpace(stdout.String())
	start := strings.Index(outStr, "{")
	if start == -1 {
		fmt.Println("No JSON found in output.")
		return
	}

	jsonPart := outStr[start:]

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonPart), &data); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		fmt.Println("Raw JSON:", jsonPart)
		return
	}

	// Print useful fields
	fmt.Printf("Video ID: %s\n", data["id"])
	fmt.Printf("Title: %s\n", data["title"])
}

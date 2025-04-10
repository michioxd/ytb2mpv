package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting ytb2mpv server...")
	InitCfg()
	go func() {
		RunTray()
	}()
	RunWSServer()
}

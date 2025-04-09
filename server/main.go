package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting ytb2mpv server...")
	go func() {
		RunTray()
	}()
	RunWSServer()
}

package main

func main() {
	InitCfg()
	go func() {
		RunTray()
	}()
	RunWSServer()
}

package main

import (
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

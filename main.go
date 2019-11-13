package main

import (
	"log"

	"golang.org/x/mobile/app"
	//"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case mouse.Event:
				log.Print(e)
				a.Publish()
			}
		}
	})
}

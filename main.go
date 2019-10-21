package main

import (
	ui "github.com/gizak/termui/v3"
	"log"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	initGame()

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			if e.ID == "<Left>" || e.ID == "a" {
				move(Left)
			} else if e.ID == "<Right>" || e.ID == "d" {
				move(Right)
			} else if e.ID == "<Up>" || e.ID == "w" {
				move(Up)
			} else if e.ID == "<Down>" || e.ID == "s" {
				move(Down)
			} else if e.ID == "<Backspace>" {
				//TODO: undo
			} else if e.ID == "<Enter>" && gameOverWaiting {
				resetGame()
			} else if e.ID == "<Escape>" {
				saveGame()
				return
			}
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

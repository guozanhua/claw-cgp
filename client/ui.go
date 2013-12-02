// Author: sheppard(ysf1026@gmail.com) 2013-12-01

package main

import (
	"github.com/nsf/termbox-go"
)

const title = "claw-cgp"

func printTitle() {
	printWide(0, 0, title, termbox.ColorDefault, termbox.ColorRed)
}

func printWide(x, y int, s string, fg, bg termbox.Attribute) {
	for _, r := range s {
		termbox.SetCell(x, y, r, fg, bg)
		x += 1
	}
}

var (
	rooms []string
)

func printResult() {
	y := 1
	for _, room := range rooms {
		printWide(0, y, room, termbox.ColorDefault, termbox.ColorDefault)
		y += 1
	}
}

func drawAll() {
	printTitle()
	printResult()
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	drawAll()

	ctrlxpressed := false
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ && ctrlxpressed {
				break loop
			}
			if ev.Key == termbox.KeyCtrlX {
				ctrlxpressed = true
			} else {
				ctrlxpressed = false
				rooms = append(rooms, "room" + string(ev.Key))
				drawAll()
			}
		case termbox.EventResize:
			drawAll()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

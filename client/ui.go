// Author: sheppard(ysf1026@gmail.com) 2013-12-01

package main

import (
	"github.com/nsf/termbox-go"
)

const test = "claw-cgp"

func printWide(x, y int, s string) {
	for _, r := range s {
		c := termbox.ColorRed
		termbox.SetCell(x, y, r, termbox.ColorDefault, c)
		x += 1
	}
}

func drawAll() {
	printWide(0, 0, test)
	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	drawAll()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			}
		case termbox.EventResize:
			drawAll()
		}
	}
}

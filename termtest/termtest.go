package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	//"os"
	"unicode/utf8"
)



func printAt(x, y int, s string, fg, bg termbox.Attribute) {
	for len(s) > 0 {
		r, rlen := utf8.DecodeRuneInString(s)
		termbox.SetCell(x, y, r, fg, bg)
		s = s[rlen:]
		x++
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	x, y := termbox.Size()
	var curx, cury int

	printAt(0, 0, fmt.Sprintf("[%d, %d]", x, y), termbox.ColorDefault, termbox.ColorDefault)

	printAt(10, 10, "Hello World ££££", termbox.ColorDefault, termbox.ColorDefault)

	termbox.SetCursor(curx, cury)
	termbox.SetCell(10, 10, 'x', termbox.ColorDefault, termbox.ColorDefault)

	printAt(10, 10, "Hello World ££££", termbox.ColorDefault, termbox.ColorDefault)
	printAt(10, 11, "┬┤├└", termbox.ColorDefault, termbox.ColorDefault)

	termbox.Flush()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			printAt(0, 1, fmt.Sprint(ev), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()

			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
				cury--
			case termbox.KeyArrowDown:
				cury++
			case termbox.KeyArrowLeft:
				curx--
			case termbox.KeyArrowRight:
				curx++
			}
			termbox.SetCursor(curx, cury)
			termbox.Flush()

		case termbox.EventResize:
			x, y := ev.Width, ev.Height
			printAt(0, 0, fmt.Sprintf("[%d, %d]      ", x, y), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()
		}
	}
}

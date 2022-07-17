package game

import (
	"os"
	"time"

	"github.com/gdamore/tcell"
)

var (
	InMenu = true
)

func ProcessInput(e tcell.Event, s tcell.Screen, snake snake) {
	switch e := e.(type) {
	case *tcell.EventKey:
		switch e.Key() {
		case tcell.KeyEscape:
			s.Fini()
			os.Exit(0)
		case tcell.KeyUp:
			if !InMenu {
				snake.move(s, 1)
				time.Sleep(2 * time.Second)
			}
		}
	case *tcell.EventMouse:
		{
			if InMenu {
				if e.Buttons() == tcell.Button1 {
					x, y := e.Position()
					if y == 8 && (x > 54 && x < 60) {
						InMenu = false
						s.Clear()
					}
					if y == 11 && (x > 54 && x < 59) {
						s.Fini()
						os.Exit(0)
					}
				}
			}
		}
	}
}

package game

import (
	"os"

	"github.com/gdamore/tcell"
)

var (
	InMenu = true
)

func ProcessInput(e tcell.Event, s tcell.Screen) {
	switch e := e.(type) {
	case *tcell.EventKey:
		switch e.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		}
	case *tcell.EventMouse:
		{
			if InMenu {
				if e.Buttons() == tcell.Button1 {
					x, y := e.Position()
					if y == 8 && (x > 54 && x < 60) { //51-59 5
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

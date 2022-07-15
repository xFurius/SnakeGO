package game

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

var (
	inMenu = true
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
			if inMenu {
				if e.Buttons() == tcell.Button1 {
					x, y := e.Position()
					if y == 5 && (x > 50 && x < 60) { //51-59 5
						fmt.Println("AFSFSA")
						inMenu = false
					}
				}
			}
		}
	}
}

package game

import (
	"log"

	"github.com/gdamore/tcell"
)

func ScreenInit() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalln(err)
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatalln(err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	screen.SetStyle(defStyle)

	return screen
}

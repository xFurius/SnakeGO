package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func Init() {
	logFile, err := CLogFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(logFile)
}

func main() {
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

	screen.Clear()

	w, h := screen.Size()

	screen.SetContent(w-1, h-1, 'H', nil, defStyle)

	for {
		screen.Show()

		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				os.Exit(0)
			}
		}
	}
}

func CLogFile() (*os.File, error) {
	return os.Create("./log.txt")
}

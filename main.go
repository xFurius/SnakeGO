package main

import (
	"fmt"
	"log"
	"os"
	"snake/main/game"
)

func Init() {
	logFile, err := os.Create("./log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(logFile)
}

func main() {
	screen := game.ScreenInit()
	for {
		screen.Show()
	}
	// defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	// screen.SetStyle(defStyle)

	// screen.Clear()

	// w, h := screen.Size()

	// screen.SetContent(w-1, h-1, 'H', nil, defStyle)

	// for {
	// 	screen.Show()

	// 	ev := screen.PollEvent()

	// 	switch ev := ev.(type) {
	// 	case *tcell.EventResize:
	// 		screen.Sync()
	// 	case *tcell.EventKey:
	// 		if ev.Key() == tcell.KeyEscape {
	// 			os.Exit(0)
	// 		}
	// 	}
	// }
}

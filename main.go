package main

import (
	"fmt"
	"log"
	"os"
	"snake/main/game"
	"snake/main/menu"

	"github.com/gdamore/tcell"
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
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalln(err)
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatalln(err)
	}

	screen.Clear()
	arena := game.InitArena(screen, []int{1, 1})

	for {
		screen.Show()

		ev := screen.PollEvent()
		game.ProcessInput(ev, screen)
		menu.RenderMenu(screen)
		if !game.InMenu {
			arena.RenderArena(screen)
		}
	}
}

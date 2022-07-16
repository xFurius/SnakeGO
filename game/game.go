package game

import (
	"log"

	"github.com/gdamore/tcell"
)

func Ihasfi() { //need  to change this name obviously
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalln(err)
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatalln(err)
	}

	screen.Clear()
	arena := InitArena(screen, []int{1, 1})
	snake := InitSnake(1, []point{
		{35, 10},
		{35, 11},
		{35, 12},
	})

	for {
		screen.Show()

		ev := screen.PollEvent()
		ProcessInput(ev, screen, snake)
		RenderMenu(screen)
		if !InMenu {
			arena.RenderArena(screen)
			RenderArenaUI(screen)
			snake.renderSnake(screen)
		}
	}
}

package game

import "github.com/gdamore/tcell"

type Arena struct {
	Width    int
	Height   int
	FruitPos []int
}

func InitArena(screen tcell.Screen, fruitPos []int) Arena {
	w, h := screen.Size()
	return Arena{
		w / 2,
		h / 2,
		fruitPos,
	}
}

func (arena *Arena) RenderArena() {

}

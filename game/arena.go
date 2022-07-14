package game

import "github.com/gdamore/tcell"

type Arena struct {
	Width    int
	Height   int
	FruitPos []int
}

func InitArena(s tcell.Screen, fruitPos []int) Arena {
	w, h := s.Size()
	return Arena{
		w / 2,
		h / 2,
		fruitPos,
	}
}

func (arena *Arena) RenderArena(s tcell.Screen) {

	style := tcell.Style.Background(tcell.StyleDefault, tcell.ColorDefault).Foreground(tcell.ColorRebeccaPurple)
	s.SetStyle(style)
	s.SetContent(1, 1, 'H', nil, style)
}

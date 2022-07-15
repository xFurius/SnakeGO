package game

import "github.com/gdamore/tcell"

type arena struct {
	width    int
	height   int
	fruitPos []int
}

func InitArena(s tcell.Screen, fruitPos []int) arena {
	w, h := s.Size()
	return arena{
		w - 30,
		h - 5,
		fruitPos,
	}
}

func (arena *arena) RenderArena(s tcell.Screen) {
	s.DisableMouse()
	s.Clear()
	style := tcell.Style.Background(tcell.StyleDefault, tcell.ColorDefault).Foreground(tcell.ColorDefault)
	s.SetStyle(style)
	for i := 30; i < arena.width; i++ {
		s.SetContent(i, arena.height, '─', nil, style)
		s.SetContent(i, arena.height-20, '─', nil, style)
	}

	for i := 6; i < arena.height; i++ {
		s.SetContent(arena.width, i, '│', nil, style)
		s.SetContent(arena.width-61, i, '│', nil, style)
	}

	s.SetContent(arena.width-61, arena.height-20, '┌', nil, style)
	s.SetContent(arena.width, arena.height-20, '┐', nil, style)
	s.SetContent(arena.width, arena.height, '┘', nil, style)
	s.SetContent(arena.width-61, arena.height, '└', nil, style)

	// '┌'
	// '└'
	// '┐'
	// '┘'
}

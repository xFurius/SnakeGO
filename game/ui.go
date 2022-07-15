package game

import (
	"github.com/gdamore/tcell"
)

func RenderMenu(s tcell.Screen) {
	s.EnableMouse()
	w, h := s.Size()
	renderText(w/2+2, h-25, "MAIN MENU", s)
	renderText(w/2, h-22, "START", s)
	renderText(w/2-1, h-19, "EXIT", s)
}

func RenderArenaUI(s tcell.Screen) {
	w, h := s.Size()
	renderText(w-70, h-26, "SNAKE THE GAME", s)
	renderText(w-75, h-5, " POINTS: ", s)
}

func renderText(x, y int, text string, s tcell.Screen) {
	style := tcell.Style.Background(tcell.StyleDefault, tcell.ColorDefault).Foreground(tcell.ColorDefault)
	s.SetStyle(style)
	x -= len(text)
	for _, v := range text {
		s.SetContent(x, y, v, nil, style)
		x++
	}
}

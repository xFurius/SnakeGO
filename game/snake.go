package game

import "github.com/gdamore/tcell"

type snake struct {
	lenght    int
	body      []point
	direction int //0-3
}

func InitSnake(lenght, direction int, body []point) snake {
	return snake{
		lenght,
		body,
		direction,
	}
}

func (snake *snake) renderSnake(s tcell.Screen) {
	style := tcell.Style.Background(tcell.StyleDefault, tcell.ColorGreen).Foreground(tcell.ColorDefault)
	s.SetContent(1, 1, ' ', nil, style)
}

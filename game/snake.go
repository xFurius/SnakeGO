package game

import (
	"github.com/gdamore/tcell"
)

type snake struct {
	body      []point
	direction int //0-3
}

func InitSnake(direction int, body []point) snake {
	return snake{
		body,
		direction,
	}
}

func (snake *snake) renderSnake(s tcell.Screen) {
	style := tcell.Style.Background(tcell.StyleDefault, tcell.ColorGreen).Foreground(tcell.ColorDefault)
	for _, v := range snake.body {
		s.SetContent(v.x, v.y, ' ', nil, style)
	}
}

func (snake *snake) move(s tcell.Screen, direction int) {
	style := tcell.Style.Background(tcell.StyleDefault, tcell.ColorGreen).Foreground(tcell.ColorDefault)

	snake.body[0].x++
	s.SetContent(snake.body[0].x, snake.body[0].y, ' ', nil, style)
}

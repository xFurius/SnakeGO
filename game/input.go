package game

import (
	"os"

	"github.com/gdamore/tcell"
)

func ProcessInput(e tcell.Event) {
	switch ev := e.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			os.Exit(0)
		}
	}
}

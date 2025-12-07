package game

import (
	"github.com/gdamore/tcell/v2"
)

// HandleInput processes keyboard events and translates them to game actions.
// Arrow keys move/rotate pieces, spacebar triggers hard drop, +/- adjust level, 'p' pauses, 'q' and Esc quit.
func HandleInput(gs *GameState, ev tcell.Event) {
	switch e := ev.(type) {
	case *tcell.EventKey:
		switch e.Key() {
		case tcell.KeyLeft:
			gs.MovePiece(-1, 0)
		case tcell.KeyRight:
			gs.MovePiece(1, 0)
		case tcell.KeyDown:
			gs.MovePiece(0, 1)
		case tcell.KeyUp:
			gs.RotatePiece()
		case tcell.KeyEsc:
			gs.GameOver = true
		case tcell.KeyRune:
			switch e.Rune() {
			case 'q':
				gs.GameOver = true
			case ' ':
				gs.HardDrop()
			case 'p':
				// Pause functionality can be implemented here
				screen.PollEvent() // simple pause until next key press
			case '-':
				if gs.DecreaseLevel() {
					gs.SetTickerInterval(gs.Level.Number)
				}
			case '+':
				if gs.IncreaseLevel() {
					gs.SetTickerInterval(gs.Level.Number)
				}
			}
		}
	}
}

package game

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

// Loop runs the main game loop until game over.
// Handles timed piece drops, processes input events, and updates the display.
// Uses a ticker for consistent drop speed and a goroutine for non-blocking event polling.
func Loop(gs *GameState) {
	gs.Ticker = time.NewTicker(time.Millisecond * 500) // швидкість падіння
	defer gs.Ticker.Stop()
	evCh := make(chan tcell.Event, 16) // буфер корисний при сплесках подій
	quit := make(chan struct{})

	// goroutine для PollEvent
	go func() {
		defer close(evCh)
		for {
			select {
			case <-quit:
				return
			default:
				ev := screen.PollEvent() // блокує тут, але не в головній горутині
				if ev == nil {
					continue
				}
				evCh <- ev
			}
		}
	}()

	for !gs.GameOver {
		select {
		case <-gs.Ticker.C:
			gs.MovePiece(0, 1)
			lines := gs.ClearLines()
			if lines > 0 {
				gs.TetrisRate.AddTetraLines(lines)
				gs.UpdateScore(lines)
				if gs.UpdateLevel() {
					// adjust tick speed based on level
					gs.SetTickerInterval(gs.Level.Number)
				}
				// track total cleared lines
				gs.Lines += lines
			}
			if gs.IsGameOver() {
				gs.GameOver = true
			}
			gs.EventName = "tick"
			gs.DrawBoard()

		case ev := <-evCh:
			if ev == nil {
				continue
			}
			HandleInput(gs, ev)
			gs.EventName = "input"
			gs.DrawBoard()
			gs.SelectCount++
		}

	}
	close(quit)

	if gs.GameOver {
		if gs.R != nil {
			gs.R.PollEvent()
		} else {
			time.Sleep(time.Second)
		}
	}
}

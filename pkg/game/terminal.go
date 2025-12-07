package game

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

// screen is the global tcell terminal screen used for rendering.
var screen tcell.Screen

// InitTerminal initializes the terminal for Tetris gameplay.
// Sets up tcell screen, hides cursor, and clears the display.
// Must be called before any rendering operations.
func InitTerminal() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("Cannot create screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Cannot init screen: %v", err)
	}
	screen.HideCursor()
	screen.Clear()
}

// RestoreTerminal cleans up and restores the terminal to its original state.
// Should be called in a defer statement after InitTerminal to ensure cleanup.
func RestoreTerminal() {
	screen.Fini()
}

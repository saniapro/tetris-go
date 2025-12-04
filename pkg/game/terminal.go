package game

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen

// -------------------------
// ІНІЦІАЛІЗАЦІЯ
// -------------------------

func InitTerminal() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("Cannot create screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Cannot init screen: %v", err)
	}

	screen.Clear()
}

func RestoreTerminal() {
	screen.Fini()
}

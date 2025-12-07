package main

import "github.com/saniapro/tetris/pkg/game"

// main initializes the terminal, creates a new game, and runs the game loop.
// Ensures terminal is properly restored on exit.
func main() {
	game.InitTerminal()
	defer game.RestoreTerminal()

	gs := game.Init()
	game.Loop(gs)
}

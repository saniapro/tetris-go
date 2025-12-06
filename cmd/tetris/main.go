package main

import "github.com/saniapro/tetris/pkg/game"

func main() {
	game.InitTerminal()
	defer game.RestoreTerminal()

	gs := game.Init()
	game.Loop(gs)
}

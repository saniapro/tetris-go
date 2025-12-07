package game

import "github.com/saniapro/tetris/pkg/tetris"

// Init initializes a new GameState with a fresh board, spawns initial pieces,
// and sets up the renderer and tetromino bag generator.
// Returns a ready-to-play GameState.
func Init() *GameState {
	gs := &GameState{
		Board:      tetris.NewBoard(),
		R:          &ScreenRenderer{},
		Level:      Level{Number: 1, ManualNumber: false},
		TetrisStat: &tetris.TetrisStat{},
		Generator:  tetris.NewBagGenerator(0),
	}

	gs.Current = tetris.SpawnPiece(gs.TetrisStat, gs.Generator.Next())
	gs.Next = tetris.SpawnPiece(gs.TetrisStat, gs.Generator.Next())

	return gs
}

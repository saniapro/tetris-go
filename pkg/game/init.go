package game

import (
	"time"

	"github.com/saniapro/tetris/pkg/tetris"
)

// Init initializes a new GameState with a fresh board, spawns initial pieces,
// and sets up the renderer and tetromino bag generator.
// Returns a ready-to-play GameState.
func Init() *GameState {
	gs := &GameState{
		Board:      tetris.NewBoard(),
		R:          &ScreenRenderer{},
		Level:      tetris.Level{Number: 1},
		TetrisRate: &tetris.TetrisRate{},
		Generator:  tetris.NewBagGenerator(time.Now().UnixNano()),
	}

	gs.Current = tetris.SpawnPiece(gs.Generator.Next())
	gs.Next = tetris.SpawnPiece(gs.Generator.Next())

	return gs
}

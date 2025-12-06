package game

import "github.com/saniapro/tetris/pkg/tetris"

func Init() *GameState {
	gs := &GameState{
		Board:      tetris.NewBoard(),
		R:          &ScreenRenderer{},
		Level:      Level{Number: 1, ManualNumber: false},
		TetrisStat: &tetris.TetrisStat{},
	}

	gs.Current = tetris.SpawnPiece(gs.TetrisStat)
	gs.Next = tetris.SpawnPiece(gs.TetrisStat)

	return gs
}

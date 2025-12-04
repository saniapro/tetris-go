package game

import "github.com/saniapro/tetris/pkg/tetris"

func Init() *GameState {
	gs := &GameState{
		Board: make([][]int, 20),
		Level: 1,
	}

	for i := range gs.Board {
		gs.Board[i] = make([]int, 10)
	}

	gs.Current = tetris.SpawnPiece()
	gs.Next = tetris.SpawnPiece()

	return gs
}

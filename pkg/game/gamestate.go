package game

import "github.com/saniapro/tetris/pkg/tetris"

type GameState struct {
	Board    [][]int
	Current  tetris.Piece
	Next     tetris.Piece
	Score    int
	Level    int
	GameOver bool
}

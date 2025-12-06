package tetris

import "fmt"

const TetrisPiece = 0

type TetrisStat struct {
	Total  int
	Tetris int
}

func (ts *TetrisStat) AddTetraLines(lines int) {
	if lines == 4 {
		ts.Tetris++
	}
}

func (ts *TetrisStat) AddTotal(pieceNumber int) {
	if pieceNumber == TetrisPiece {
		ts.Total++
	}
}

func (ts *TetrisStat) Reset() {
	ts.Total = 0
	ts.Tetris = 0
}

func (ts *TetrisStat) GetPercent() string {
	if ts.Total == 0 {
		return "0%"
	}
	percent := (float64(ts.Tetris) / float64(ts.Total)) * 100
	return fmt.Sprintf("%.2f%%", percent)
}

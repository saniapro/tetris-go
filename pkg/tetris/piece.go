package tetris

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type Piece struct {
	Matrix [][]int
	X, Y   int
	Color  tcell.Color
}

const Fill = 1

func SpawnPiece(ts *TetrisStat) Piece {
	n := rand.Intn(len(Pieces))
	ts.AddTotal(n)
	return Pieces[n]
}

func RotatePiece(p Piece) Piece {
	n := len(p.Matrix)
	m := len(p.Matrix[0])
	newMatrix := make([][]int, m)
	for i := range newMatrix {
		newMatrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			newMatrix[j][n-1-i] = p.Matrix[i][j]
		}
	}
	return Piece{
		Matrix: newMatrix,
		X:      p.X,
		Y:      p.Y,
		Color:  p.Color,
	}
}

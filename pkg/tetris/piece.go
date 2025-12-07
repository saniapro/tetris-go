package tetris

import (
	"github.com/gdamore/tcell/v2"
)

// Piece represents a tetromino with its matrix, position, and color.
type Piece struct {
	Matrix [][]int     // 2D grid defining the piece shape (Fill=1, empty=0)
	X, Y   int         // Position on the board
	Color  tcell.Color // Rendering color
}

const Fill = 1 // Marker value for filled cells in piece matrices

// SpawnPiece creates a new tetromino from the pre-defined pieces array.
// Updates TetrisStat tracking and uses the bag generator index.
func SpawnPiece(ts *TetrisStat, n int) Piece {
	ts.AddTotal(n)
	return Pieces[n]
}

// RotatePiece rotates the given piece 90 degrees clockwise.
// Preserves the piece's position and color across rotations.
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

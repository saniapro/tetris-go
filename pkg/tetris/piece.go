package tetris

import (
	"github.com/gdamore/tcell/v2"
)

// Piece represents a tetromino with its matrix, position, and color.
type Piece struct {
	Matrix   [][]int     // 2D grid defining the piece shape (Fill=1, empty=0)
	X, Y     int         // Position on the board
	Color    tcell.Color // Rendering color
	ID       int         // Piece index (0-6) from the bag
	Rotation int         // Rotation state 0-3
}

const Fill = 1 // Marker value for filled cells in piece matrices

// SpawnPiece creates a new tetromino from the pre-defined pieces array.
// Updates TetrisRate tracking and uses the bag generator index.
func SpawnPiece(n int) Piece {
	// Return a copy so mutations (rotations/moves) don't change the global template
	template := Pieces[n]
	// deep copy matrix
	m := make([][]int, len(template.Matrix))
	for i := range template.Matrix {
		m[i] = make([]int, len(template.Matrix[i]))
		copy(m[i], template.Matrix[i])
	}
	return Piece{
		Matrix:   m,
		X:        template.X,
		Y:        template.Y,
		Color:    template.Color,
		ID:       n,
		Rotation: 0,
	}
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

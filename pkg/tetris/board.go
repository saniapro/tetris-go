package tetris

import (
	"slices"

	"github.com/gdamore/tcell/v2"
)

const (
	BoardWidth   = 10 // Standard Tetris board width
	BoardHeight  = 20 // Standard Tetris board height
	BoardXOffset = 2  // Horizontal offset for board display
	BoardYOffset = 2  // Vertical offset for board display
)

// Row represents a single row of the Tetris board, storing the color of each cell.
type Row []tcell.Color

// Board represents the Tetris playing field.
type Board struct {
	grid []Row
}

// NewBoard creates a new board with the standard Tetris dimensions (10x20).
// All cells are initially empty (color 0).
func NewBoard() *Board {
	grid := make([]Row, BoardHeight)
	for i := range grid {
		grid[i] = make(Row, BoardWidth)
	}
	return &Board{
		grid: grid,
	}
}

// CellFilled checks if a cell at (row, col) is occupied (non-zero color).
// Returns false for out-of-bounds queries.
func (b *Board) CellFilled(row, col int) bool {
	if row < 0 || row >= BoardHeight || col < 0 || col >= BoardWidth {
		return false // out of bounds
	}
	return b.grid[row][col] != 0
}

// SetCell places a colored block at the given (row, col) position.
// Silently ignores out-of-bounds assignments.
func (b *Board) SetCell(row, col int, value tcell.Color) {
	if row >= 0 && row < BoardHeight && col >= 0 && col < BoardWidth {
		b.grid[row][col] = tcell.Color(value)
	}
}

// Row returns the Row at the given index.
// Returns nil for out-of-bounds indices.
func (b *Board) Row(index int) Row {
	if index < 0 || index >= BoardHeight {
		return nil
	}
	return b.grid[index]
}

// CellFilled checks if a cell in the row at the given column is occupied.
// Returns false for out-of-bounds queries.
func (r Row) CellFilled(col int) bool {
	if col < 0 || col >= len(r) {
		return false // out of bounds
	}
	return r[col] != 0
}

// ClearLines removes all completed (fully filled) rows from the board.
// Completed rows are removed from the bottom up, and new empty rows are added at the top.
// Returns the count of rows cleared.
func (b *Board) ClearLines() int {
	linesCleared := 0
	for i := len(b.grid) - 1; i >= 0; i-- {
		if !slices.Contains(b.grid[i], 0) {
			linesCleared++
			// remove line
			b.grid = append(b.grid[:i], b.grid[i+1:]...)
			// add empty line at the top
			newLine := make(Row, BoardWidth)
			b.grid = append([]Row{newLine}, b.grid...)
			// check same line again
			i++
		}
	}
	return linesCleared
}

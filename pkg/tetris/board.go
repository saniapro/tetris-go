package tetris

import "github.com/gdamore/tcell/v2"

const (
	BoardWidth   = 10
	BoardHeight  = 20
	BoardXOffset = 2
	BoardYOffset = 2
)

type Row []tcell.Color

type Board struct {
	grid []Row
}

// NewBoard creates a new board with the standard Tetris dimensions
func NewBoard() *Board {
	grid := make([]Row, BoardHeight)
	for i := range grid {
		grid[i] = make(Row, BoardWidth)
	}
	return &Board{
		grid: grid,
	}
}

// Cell returns the value at the given row and column
func (b *Board) CellFilled(row, col int) bool {
	if row < 0 || row >= BoardHeight || col < 0 || col >= BoardWidth {
		return false // out of bounds
	}
	return b.grid[row][col] != 0
}

// SetCell sets the value at the given row and column
func (b *Board) SetCell(row, col int, value tcell.Color) {
	if row >= 0 && row < BoardHeight && col >= 0 && col < BoardWidth {
		b.grid[row][col] = tcell.Color(value)
	}
}

// Row returns a copy of the row at the given index
func (b *Board) Row(index int) Row {
	if index < 0 || index >= BoardHeight {
		return nil
	}
	return b.grid[index]
}

// Cell returns the value at the given row
func (r Row) CellFilled(col int) bool {
	if col < 0 || col >= len(r) {
		return false // out of bounds
	}
	return r[col] != 0
}

// ClearLines removes completed lines and returns the number of lines cleared
func (b *Board) ClearLines() int {
	linesCleared := 0
	for i := len(b.grid) - 1; i >= 0; i-- {
		fullLine := true
		for _, cell := range b.grid[i] {
			if cell == 0 {
				fullLine = false
				break
			}
		}
		if fullLine {
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

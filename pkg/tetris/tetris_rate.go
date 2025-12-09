package tetris

import "fmt"

const TetrisPiece = 0 // Index of the I-tetromino

// TetrisRate tracks tetris-specific statistics for performance analysis.
// Specifically monitors the ratio of 4-line clears to all lines cleared.
type TetrisRate struct {
	Tetris    int // Sum of 4-line clears (perfect tetrises)
	NonTetris int // Sum of non-tetris line clears
}

// AddTetraLines increments the tetris counter if a 4-line clear occurred.
func (tr *TetrisRate) AddTetraLines(lines int) {
	if lines == 4 {
		tr.Tetris += lines
	} else if lines > 0 {
		tr.NonTetris += lines
	}
}

// Reset clears all tetris statistics.
func (tr *TetrisRate) Reset() {
	tr.Tetris = 0
	tr.NonTetris = 0
}

// GetPercent returns the percentage of spawned I-pieces that resulted in tetrises.
// Returns "0%" if no I-pieces have been spawned.
func (tr *TetrisRate) GetPercent() string {
	if tr.NonTetris == 0 && tr.Tetris == 0 {
		return "0%"
	}
	percent := (float64(tr.Tetris) / float64(tr.NonTetris+tr.Tetris)) * 100
	return fmt.Sprintf("%.0f%%", percent)
}

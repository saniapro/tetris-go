package tetris

import "fmt"

const TetrisPiece = 0 // Index of the I-tetromino

// TetrisStat tracks tetris-specific statistics for performance analysis.
// Specifically monitors the ratio of 4-line clears to I-piece spawns.
type TetrisStat struct {
	Total  int // Count of I-pieces spawned
	Tetris int // Count of 4-line clears (perfect tetrises)
}

// AddTetraLines increments the tetris counter if a 4-line clear occurred.
func (ts *TetrisStat) AddTetraLines(lines int) {
	if lines == 4 {
		ts.Tetris++
	}
}

// AddTotal increments the total counter when an I-piece (index 0) is spawned.
func (ts *TetrisStat) AddTotal(pieceNumber int) {
	if pieceNumber == TetrisPiece {
		ts.Total++
	}
}

// Reset clears all tetris statistics.
func (ts *TetrisStat) Reset() {
	ts.Total = 0
	ts.Tetris = 0
}

// GetPercent returns the percentage of spawned I-pieces that resulted in tetrises.
// Returns "0%" if no I-pieces have been spawned.
func (ts *TetrisStat) GetPercent() string {
	if ts.Total == 0 {
		return "0%"
	}
	percent := (float64(ts.Tetris) / float64(ts.Total)) * 100
	return fmt.Sprintf("%.2f%%", percent)
}

package game

import (
	"time"

	"github.com/saniapro/tetris/pkg/tetris"
)

// GameState holds the complete state of a Tetris game including board, pieces, score, and rendering context.
type GameState struct {
	Board       *tetris.Board
	Current     tetris.Piece
	Next        tetris.Piece
	EventName   string
	SelectCount int
	Lines       int
	Score       int
	Level       tetris.Level
	GameOver    bool
	R           Renderer
	Ticker      *time.Ticker
	TetrisRate  *tetris.TetrisRate
	Generator   *tetris.BagGenerator
}

// TetrisRate tracks tetromino spawn statistics for gameplay analysis.
type TetrisRate struct {
	Total  int
	Tetris int
}

const msPerLevel = 10

// SetTickerInterval adjusts the game tick rate based on the level number.
// Higher levels result in faster piece drops.
// Interval scales from 500ms at level 1, decreasing by 10ms per level, with a minimum of 10ms.
func (gs *GameState) SetTickerInterval(LevelNumber int) {
	newInterval := max(time.Millisecond*time.Duration(500-(LevelNumber-1)*msPerLevel), time.Millisecond*msPerLevel)
	gs.Ticker.Reset(newInterval)
}

// MovePiece attempts to move the current piece by dx (horizontal) and dy (vertical) pixels.
// If the move results in a collision (fitFloor), the piece is locked.
// If the move is impossible (fitImpossible), it is reverted.
func (gs *GameState) MovePiece(dx, dy int) {
	gs.Current.X += dx
	gs.Current.Y += dy

	// collision detection and handling would go here
	if fit := gs.Fit(); fit != fitPossible {
		// revert move
		gs.Current.X -= dx
		gs.Current.Y -= dy
		if dy > 0 && fit == fitFloor {
			gs.LockPiece()
		}
	}
}

const (
	fitImpossible = iota // Piece cannot fit (collision with board edge or existing blocks)
	fitPossible          // Piece fits in the current position
	fitFloor             // Piece has reached the floor and should be locked
)

// Fit checks if the current piece can fit at its current position on the board.
// Returns one of: fitPossible, fitFloor (hit bottom), or fitImpossible (collision).
func (gs *GameState) Fit() int {
	for i, row := range gs.Current.Matrix {
		for j, cell := range row {
			if cell != 0 {
				if gs.Current.X+j < 0 || gs.Current.X+j >= tetris.BoardWidth {
					return fitImpossible
				}
				if gs.Current.Y+i >= tetris.BoardHeight ||
					gs.Board.CellFilled(gs.Current.Y+i, gs.Current.X+j) {
					return fitFloor
				}
			}
		}
	}
	return fitPossible
}

// LockPiece finalizes the current piece by placing it on the board.
// Spawns the next piece and checks for game over after locking.
func (gs *GameState) LockPiece() {
	for i, row := range gs.Current.Matrix {
		for j, cell := range row {
			if cell != 0 {
				gs.Board.SetCell(gs.Current.Y+i, gs.Current.X+j, gs.Current.Color)
			}
		}
	}
	gs.Current = gs.Next
	// check for game over
	if gs.IsGameOver() {
		gs.GameOver = true
	}
	gs.Next = tetris.SpawnPiece(gs.Generator.Next())
}

// HardDrop instantly drops the current piece to the bottom of the board and locks it.
// Continuously moves the piece down and redraws until it can no longer move.
func (gs *GameState) HardDrop() {
	for {
		gs.Current.Y++
		if gs.Fit() != fitPossible {
			gs.Current.Y--
			break
		}
		gs.DrawBoard()
	}
}

// RotatePiece rotates the current piece 90 degrees clockwise.
// If rotation would cause a collision, reverts to the original orientation.
func (gs *GameState) RotatePiece() {
	// Implement SRS-style rotation (clockwise). Uses simple JLSTZ and I kick tables.
	p := gs.Current
	// O piece does not change orientation in SRS
	if p.ID == 1 {
		return
	}

	origX, origY := p.X, p.Y
	origMatrix := p.Matrix
	origRotation := p.Rotation

	// rotated matrix (clockwise)
	rotated := rotateCW(p.Matrix)

	// SRS kick tests (dx, dy). dy values here follow standard SRS convention
	// where positive dy is upwards; board Y increases downward, so we'll
	// subtract dy when applying to piece Y.
	var tests [][2]int
	if p.ID == 0 { // I piece
		tests = [][2]int{{0, 0}, {-2, 0}, {1, 0}, {-2, -1}, {1, 2}}
	} else { // J, L, S, T, Z
		tests = [][2]int{{0, 0}, {-1, 0}, {-1, 1}, {0, -2}, {-1, -2}}
	}

	for _, t := range tests {
		dx, dy := t[0], t[1]
		tryX := origX + dx
		tryY := origY - dy

		gs.Current.Matrix = rotated
		gs.Current.X = tryX
		gs.Current.Y = tryY

		if gs.Fit() == fitPossible {
			gs.Current.Rotation = (origRotation + 1) % 4
			return
		}
	}

	// no valid kick found — revert
	gs.Current.Matrix = origMatrix
	gs.Current.X = origX
	gs.Current.Y = origY
	gs.Current.Rotation = origRotation
}

// rotateCW returns a new matrix representing the given matrix rotated 90 degrees clockwise.
func rotateCW(m [][]int) [][]int {
	n := len(m)
	if n == 0 {
		return [][]int{}
	}
	r := len(m[0])
	newM := make([][]int, r)
	for i := range newM {
		newM[i] = make([]int, n)
	}
	for i := range n {
		for j := range r {
			newM[j][n-1-i] = m[i][j]
		}
	}
	return newM
}

// ClearScreen clears the entire terminal display.
func (gs *GameState) ClearScreen() {
	if gs.R != nil {
		gs.R.Clear()
	}
}

// IncreaseLevel increments the level by 1 and marks it as manually set.
// Returns true if the level successfully increased.
func (gs *GameState) IncreaseLevel() bool {
	return gs.Level.Set(gs.Level.Number+1, true)
}

// DecreaseLevel decrements the level by 1 (minimum 1) and marks it as manually set.
// Returns true if the level successfully decreased, false if already at minimum.
func (gs *GameState) DecreaseLevel() bool {
	if gs.Level.Number > 1 {
		return gs.Level.Set(gs.Level.Number-1, true)
	}
	return false
}

// UpdateScore increments the score based on the number of lines cleared.
// Scoring follows standard Tetris rules, scaled by the current level.
func (gs *GameState) UpdateScore(lines int) {
	// tetris scrore calculation
	multiplier := gs.Level.Number
	if gs.Level.Number > 10 {
		multiplier++
	}
	switch lines {
	case 1:
		gs.Score += 40 * multiplier
	case 2:
		gs.Score += 100 * multiplier
	case 3:
		gs.Score += 300 * multiplier
	case 4:
		gs.Score += 1200 * multiplier
	}
}

// UpdateLevel automatically increases the level based on score, unless manually overridden.
// Level increases by 1 for every 10 lines cleared.
// Returns true if the level changed, false otherwise.
func (gs *GameState) UpdateLevel() bool {
	newLevel := gs.Lines/10 + 1
	return gs.Level.Set(newLevel, false)
}

// IsGameOver checks if the game has ended by testing if the top row is blocked.
// Game ends when the newly spawned piece collides with existing blocks.
func (gs *GameState) IsGameOver() bool {
	// Implementation to check if the game is over
	for j, cell := range gs.Current.Matrix[0] {
		if cell == tetris.Fill && gs.Board.CellFilled(gs.Current.Y, gs.Current.X+j) {
			// panic(fmt.Sprintf("Game Over detected at position (%d, %d) of piece %d", gs.Current.X, gs.Current.Y, j))
			return true
		}
	}
	return false
}

// ClearLines removes completed rows from the board and returns the number of rows cleared.
func (gs *GameState) ClearLines() int {
	return gs.Board.ClearLines()
}

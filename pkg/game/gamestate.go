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
	Level       Level
	GameOver    bool
	R           Renderer
	Ticker      *time.Ticker
	TetrisStat  *tetris.TetrisStat
	Generator   *tetris.BagGenerator
}

// Level represents the current game level with manual override capability.
type Level struct {
	Number       int
	ManualNumber bool
}

// TetrisStat tracks tetromino spawn statistics for gameplay analysis.
type TetrisStat struct {
	Total  int
	Tetris int
}

// Set updates the level number and marks whether it was manually set.
// Returns true if the level changed, false otherwise.
// If the level is already manually set and a non-manual update is attempted, it returns false.
// Minimum level is 1.
func (l *Level) Set(n int, manual bool) bool {
	if l.ManualNumber && !manual {
		return false
	}
	if n < 1 {
		n = 1
	}
	if l.Number != n {
		l.Number = n
		l.ManualNumber = manual
		return true
	}
	return false
}

// Get returns the current level number.
func (l *Level) Get() int {
	return l.Number
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
		if fit == fitFloor {
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
	gs.Next = tetris.SpawnPiece(gs.TetrisStat, gs.Generator.Next())
}

// HardDrop instantly drops the current piece to the bottom of the board and locks it.
// Continuously moves the piece down and redraws until it can no longer move.
func (gs *GameState) HardDrop() {
	for {
		prevY := gs.Current.Y
		gs.MovePiece(0, 1)
		if gs.Current.Y == prevY || (prevY > 0 && gs.Current.Y == 0) {
			break
		}
		gs.DrawBoard()
	}
}

// RotatePiece rotates the current piece 90 degrees clockwise.
// If rotation would cause a collision, reverts to the original orientation.
func (gs *GameState) RotatePiece() {
	tPiece := gs.Current
	gs.Current = tetris.RotatePiece(gs.Current)
	if gs.Fit() != fitPossible {
		gs.Current = tPiece
	}
}

// ClearScreen clears the entire terminal display.
func (gs *GameState) ClearScreen() {
	if gs.R != nil {
		gs.R.Clear()
	}
}

// UpdateScore increments the score based on the number of lines cleared.
// Scoring: 1 line = 100 pts, 2 lines = 300 pts, 3 lines = 500 pts, 4 lines = 800 pts.
func (gs *GameState) UpdateScore(lines int) {
	// tetris scrore calculation
	switch lines {
	case 1:
		gs.Score += 100
	case 2:
		gs.Score += 300
	case 3:
		gs.Score += 500
	case 4:
		gs.Score += 800
	}
}

// UpdateLevel automatically increases the level based on score, unless manually overridden.
// Level increases by 1 for every 500 points scored.
// Returns true if the level changed, false otherwise.
func (gs *GameState) UpdateLevel() bool {
	if gs.Level.ManualNumber {
		return false
	}
	newLevel := gs.Score/500 + 1
	if newLevel != gs.Level.Number {
		gs.Level.Number = newLevel
		return true
	}
	return false
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

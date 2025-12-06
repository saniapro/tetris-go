package game

import (
	"time"

	"github.com/saniapro/tetris/pkg/tetris"
)

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
}

type Level struct {
	Number       int
	ManualNumber bool
}

type TetrisStat struct {
	Total  int
	Tetris int
}

// Set sets the level number and marks whether it's manually set.
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
func (l *Level) Get() int {
	return l.Number
}

func (gs *GameState) IncreaseLevel() bool {
	return gs.Level.Set(gs.Level.Number+1, true)
}

func (gs *GameState) DecreaseLevel() bool {
	if gs.Level.Number > 1 {
		return gs.Level.Set(gs.Level.Number-1, true)
	}
	return false
}

const msPerLevel = 25

func (gs *GameState) SetTickerInterval(LevelNumber int) {
	newInterval := max(time.Millisecond*time.Duration(500-(LevelNumber-1)*msPerLevel), time.Millisecond*msPerLevel)
	gs.Ticker.Reset(newInterval)
}

func (gs *GameState) MovePiece(dx, dy int) {
	gs.Current.X += dx
	gs.Current.Y += dy

	// collision detection and handling would go here
	isBoardFloor := false
	for i, row := range gs.Current.Matrix {
		for j, cell := range row {
			if cell != 0 {
				xBoundsExceeded := gs.Current.X+j < 0 || gs.Current.X+j >= tetris.BoardWidth
				if !xBoundsExceeded && gs.Current.Y+i >= tetris.BoardHeight ||
					gs.Board.CellFilled(gs.Current.Y+i, gs.Current.X+j) {
					isBoardFloor = true
				}
				if xBoundsExceeded || isBoardFloor {
					// revert move
					gs.Current.X -= dx
					gs.Current.Y -= dy

					if isBoardFloor {
						// lock piece in place
						gs.LockPiece()
					}
					return
				}
			}
		}
	}
}

// LockPiece locks the current piece into the board
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
	gs.Next = tetris.SpawnPiece(gs.TetrisStat)
}

// Hard drop
func (gs *GameState) HardDrop() {
	for {
		prevY := gs.Current.Y
		gs.MovePiece(0, 1)
		if gs.Current.Y == prevY || (prevY > 0 && gs.Current.Y == 0) {
			break
		}
	}
}

func (gs *GameState) RotatePiece() {
	gs.Current = tetris.RotatePiece(gs.Current)
}

func (gs *GameState) ClearScreen() {
	if gs.R != nil {
		gs.R.Clear()
	}
}

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

func (gs *GameState) ClearLines() int {
	return gs.Board.ClearLines()
}

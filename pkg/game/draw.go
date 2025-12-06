package game

import (
	"strconv"

	"github.com/saniapro/tetris/pkg/tetris"
)

const strFill = "██"

func (gs *GameState) DrawBoard() {
	gs.ClearScreen()
	// Implementation to draw the game state on the screen

	for i := 0; i < tetris.BoardHeight; i++ {
		row := gs.Board.Row(i)
		for j := 0; j < len(row); j++ {
			if row.CellFilled(j) {

				gs.R.PutStrColor(tetris.BoardXOffset+j*2+1,
					i+tetris.BoardYOffset+1,
					strFill,
					row[j])
			}
		}
		// Draw borders
		if gs.R != nil {
			gs.R.PutStr(tetris.BoardXOffset, i+tetris.BoardYOffset+1, "║")
			gs.R.PutStr(tetris.BoardWidth*2+tetris.BoardXOffset+1, i+tetris.BoardYOffset+1, "║")
		}
	}
	const borderLine = "════════════════════"
	if gs.R != nil {
		gs.R.PutStr(tetris.BoardXOffset+1, tetris.BoardYOffset, borderLine)
		gs.R.PutStr(tetris.BoardXOffset+1, tetris.BoardHeight+tetris.BoardYOffset+1, borderLine)

		gs.R.PutStr(tetris.BoardXOffset, tetris.BoardYOffset, "╔")
		gs.R.PutStr(tetris.BoardWidth*2+tetris.BoardXOffset+1, tetris.BoardYOffset, "╗")
		gs.R.PutStr(tetris.BoardXOffset, tetris.BoardHeight+tetris.BoardYOffset+1, "╚")
		gs.R.PutStr(tetris.BoardWidth*2+tetris.BoardXOffset+1, tetris.BoardHeight+tetris.BoardYOffset+1, "╝")
	}

	//draw current piece
	gs.DrawPiece(gs.Current, tetris.BoardXOffset+1, tetris.BoardYOffset+1)

	//draw score and level
	xOffset := tetris.BoardWidth*2 + tetris.BoardXOffset + 4
	if gs.R != nil {
		gs.R.PutStr(xOffset, tetris.BoardYOffset, "Score:")
		tStr := strconv.Itoa(gs.Score)
		gs.R.PutStr(xOffset+9-len(tStr), tetris.BoardYOffset+1, tStr)
		gs.R.PutStr(xOffset, tetris.BoardYOffset+3, "Level:")
		gs.R.PutStr(xOffset+5, tetris.BoardYOffset+4, strconv.Itoa(gs.Level.Get()))
		// Lines counter
		gs.R.PutStr(xOffset, tetris.BoardYOffset+5, "Lines:")
		gs.R.PutStr(xOffset+6, tetris.BoardYOffset+6, strconv.Itoa(gs.Lines))
		gs.R.PutStr(xOffset, tetris.BoardYOffset+7, "Current:")
		gs.R.PutStr(xOffset, tetris.BoardYOffset+8, strconv.Itoa(gs.Current.X)+","+strconv.Itoa(gs.Current.Y))
		gs.R.PutStr(xOffset, tetris.BoardYOffset+10, "Next:")
		gs.DrawPiece(gs.Next, xOffset, tetris.BoardYOffset+11)
		gs.R.PutStr(xOffset, tetris.BoardYOffset+17, "Tetris: "+gs.TetrisStat.GetPercent())
	}
	gs.SelectCount = 0

	if gs.GameOver {
		if gs.R != nil {
			gs.R.PutStr(3, 10, "GAME OVER")
		}
	}
	if gs.R != nil {
		gs.R.Show()
	}
}

// DrawPiece draws a piece at its current position
func (gs *GameState) DrawPiece(p tetris.Piece, xOffset, yOffset int) {
	for i, row := range p.Matrix {
		for j, cell := range row {
			if cell == tetris.Fill {
				if gs.R != nil {
					gs.R.PutStrColor((p.X+j)*2+xOffset, p.Y+i+yOffset, strFill, p.Color)
				}
			}
		}
	}
}

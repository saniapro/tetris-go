package tetris

import "github.com/gdamore/tcell/v2"

// Pieces defines the seven standard Tetris tetrominoes (I, O, T, S, Z, J, L).
// Each piece is defined with its initial matrix orientation, spawn position, and color.
var Pieces = []Piece{
	// I Piece
	{
		Matrix: [][]int{
			{0, Fill, 0, 0},
			{0, Fill, 0, 0},
			{0, Fill, 0, 0},
			{0, Fill, 0, 0},
		},
		X:     3,
		Y:     0,
		Color: tcell.Color(tcell.ColorWhite),
	},
	// O Piece
	{
		Matrix: [][]int{
			{Fill, Fill},
			{Fill, Fill},
		},
		X:     4,
		Y:     0,
		Color: tcell.Color(tcell.ColorBlue),
	},
	// T Piece
	{
		Matrix: [][]int{
			{0, Fill, 0},
			{Fill, Fill, Fill},
		},
		X:     3,
		Y:     0,
		Color: tcell.Color(tcell.ColorRed),
	},
	// S Piece
	{
		Matrix: [][]int{
			{0, Fill, Fill},
			{Fill, Fill, 0},
		},
		X:     3,
		Y:     0,
		Color: tcell.Color(tcell.ColorTeal),
	},
	// Z Piece
	{
		Matrix: [][]int{
			{Fill, Fill, 0},
			{0, Fill, Fill},
		},
		X:     3,
		Y:     0,
		Color: tcell.Color(tcell.ColorTeal),
	},
	// J Piece
	{
		Matrix: [][]int{
			{Fill, 0, 0},
			{Fill, Fill, Fill},
		},
		X:     3,
		Y:     0,
		Color: tcell.Color(tcell.ColorGreen),
	},
	// L Piece
	{
		Matrix: [][]int{
			{0, 0, Fill},
			{Fill, Fill, Fill},
		},
		X:     3,
		Y:     0,
		Color: tcell.Color(tcell.ColorGreen),
	},
}

package tetris

import "github.com/gdamore/tcell/v2"

// Pieces defines the seven standard Tetris tetrominoes (I, O, T, S, Z, J, L).
// Each piece is defined with its initial matrix orientation, spawn position, and color.
var Pieces = []Piece{
	// I Piece
	{
		Matrix: [][]int{
			{Fill, Fill, Fill, Fill},
		},
		X:        3,
		Y:        0,
		Color:    tcell.Color(tcell.ColorAqua),
		ID:       0,
		Rotation: 0,
	},
	// O Piece
	{
		Matrix: [][]int{
			{Fill, Fill},
			{Fill, Fill},
		},
		X:        4,
		Y:        0,
		Color:    tcell.Color(tcell.ColorYellow),
		ID:       1,
		Rotation: 0,
	},
	// T Piece
	{
		Matrix: [][]int{
			{0, Fill, 0},
			{Fill, Fill, Fill},
		},
		X:        3,
		Y:        0,
		Color:    tcell.Color(tcell.ColorPurple),
		ID:       2,
		Rotation: 0,
	},
	// S Piece
	{
		Matrix: [][]int{
			{0, Fill, Fill},
			{Fill, Fill, 0},
		},
		X:        3,
		Y:        0,
		Color:    tcell.Color(tcell.ColorLime),
		ID:       3,
		Rotation: 0,
	},
	// Z Piece
	{
		Matrix: [][]int{
			{Fill, Fill, 0},
			{0, Fill, Fill},
		},
		X:        3,
		Y:        0,
		Color:    tcell.Color(tcell.ColorRed),
		ID:       4,
		Rotation: 0,
	},
	// J Piece
	{
		Matrix: [][]int{
			{Fill, 0, 0},
			{Fill, Fill, Fill},
		},
		X:        3,
		Y:        0,
		Color:    tcell.Color(tcell.ColorBlue),
		ID:       5,
		Rotation: 0,
	},
	// L Piece
	{
		Matrix: [][]int{
			{0, 0, Fill},
			{Fill, Fill, Fill},
		},
		X:        3,
		Y:        0,
		Color:    tcell.Color(tcell.ColorOrange),
		ID:       6,
		Rotation: 0,
	},
}

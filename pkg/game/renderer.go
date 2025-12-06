package game

import "github.com/gdamore/tcell/v2"

// Renderer encapsulates all screen operations used by the game.
// This decouples game logic from the concrete terminal implementation.
type Renderer interface {
	PutStr(x, y int, s string)
	Show()
	Clear()
	PollEvent() tcell.Event
	// PutStrStyled writes a string at position with given tcell style.
	PutStrColor(x, y int, s string, color tcell.Color)
}

// ScreenRenderer implements Renderer using the package-level `screen`.
type ScreenRenderer struct{}

func (r *ScreenRenderer) PutStr(x, y int, s string) { screen.PutStr(x, y, s) }
func (r *ScreenRenderer) Show()                     { screen.Show() }
func (r *ScreenRenderer) Clear()                    { screen.Clear() }
func (r *ScreenRenderer) PollEvent() tcell.Event    { return screen.PollEvent() }
func (r *ScreenRenderer) PutStrColor(x, y int, s string, color tcell.Color) {
	style := tcell.StyleDefault.Foreground(color)
	screen.PutStrStyled(x, y, s, style)
}

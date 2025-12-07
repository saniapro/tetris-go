package game

import "github.com/gdamore/tcell/v2"

// Renderer encapsulates all screen operations used by the game.
// This decouples game logic from the concrete terminal implementation.
type Renderer interface {
	// PutStr writes a string at position (x, y) without styling.
	PutStr(x, y int, s string)
	// Show commits all pending draw operations to the screen.
	Show()
	// Clear clears the entire screen.
	Clear()
	// PollEvent retrieves the next pending input event or blocks until one arrives.
	PollEvent() tcell.Event
	// PutStrColor writes a string at position (x, y) in the specified color.
	PutStrColor(x, y int, s string, color tcell.Color)
}

// ScreenRenderer implements Renderer using the package-level `screen` (tcell terminal).
type ScreenRenderer struct{}

// PutStr writes a plain text string to the screen.
func (r *ScreenRenderer) PutStr(x, y int, s string) { screen.PutStr(x, y, s) }

// Show flushes all pending drawing operations to the terminal.
func (r *ScreenRenderer) Show() { screen.Show() }

// Clear clears the entire terminal.
func (r *ScreenRenderer) Clear() { screen.Clear() }

// PollEvent blocks until an input event occurs and returns it.
func (r *ScreenRenderer) PollEvent() tcell.Event { return screen.PollEvent() }

// PutStrColor writes a colored string to the screen.
func (r *ScreenRenderer) PutStrColor(x, y int, s string, color tcell.Color) {
	style := tcell.StyleDefault.Foreground(color)
	screen.PutStrStyled(x, y, s, style)
}

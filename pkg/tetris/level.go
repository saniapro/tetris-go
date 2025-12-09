package tetris

// Level represents the current game level with manual override capability.
type Level struct {
	Number       int
	ManualNumber bool
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

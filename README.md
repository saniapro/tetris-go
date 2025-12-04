# Tetris Terminal Game

A classic Tetris game implemented in Go with terminal UI support.

## Features

- Classic Tetris gameplay
- Terminal-based graphics
- Smooth controls
- Score tracking
- Game over detection

## Prerequisites

- Go 1.21 or higher

## Building

```bash
go build -o tetris ./cmd/tetris
```

## Running

```bash
./tetris
```

## Controls

- **Arrow Left/Right**: Move tetromino left/right
- **Arrow Up**: Rotate tetromino
- **Arrow Down**: Speed up falling
- **Space**: Hard drop
- **Q**: Quit game

## Project Structure

```
.
├── cmd/tetris/        # Main application entry point
├── pkg/tetris/        # Core tetris game logic
├── pkg/game/          # Game engine and terminal UI
├── go.mod             # Go module definition
├── README.md          # This file
└── .github/
    └── copilot-instructions.md
```

## Development

This project is organized following Go best practices:

- `cmd/` - Command-line applications
- `pkg/` - Reusable packages
- `go.mod` - Module definition with dependencies

## License

MIT

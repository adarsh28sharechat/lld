# Tic-Tac-Toe (Go)

Simple CLI-based Tic-Tac-Toe in Go. Two players enter row/col indices (0-2).

## Run
```bash
cd tic-tac-toe
go run .
```

## Files
- `main.go`: starts the game
- `game.go`: game loop and turn switching
- `board.go`: board, move validation, win/draw checks
- `player.go`: player entity
- `cell.go`, `cell-type.go`: cell model

## How to play
- On your turn, enter row then column (0-2).
- Invalid/out-of-range or taken cells are rejected; try again.
- Game ends on win or draw.

## Description
This project implements a classic 3x3 Tic-Tac-Toe game for two human players in the terminal. Players alternate turns, entering coordinates to place their symbols (X or O). The game announces a winner when three symbols align in a row, column, or diagonal, or declares a draw when the board is full.

## Design patterns used
- Encapsulation: `GameBoard` hides the board structure and enforces rules via methods like `MakeMove`, preventing invalid state changes.
- Separation of Concerns (Single Responsibility): `Game` orchestrates turn flow and I/O, `GameBoard` manages rules and state, `Player` represents a participant; each has a focused responsibility.
- Constructor functions (Go idiom): `NewGameBoard`, `NewPlayer`, and `NewGame` centralize initialization to ensure valid objects and readable setup.
- Guard validation: `MakeMove` validates bounds and occupancy early, returning errors immediately to keep control flow clear and robust.

package main

import "fmt"

type GameBoard struct {
	Board      [][]*Cell
	MovesCount int
}

func NewGameBoard() *GameBoard {
	board := &GameBoard{
		Board:      make([][]*Cell, 3),
		MovesCount: 0,
	}

	for i := 0; i < 3; i++ {
		board.Board[i] = make([]*Cell, 3)
		for j := 0; j < 3; j++ {
			board.Board[i][j] = &Cell{Type: &CellType{Value: ""}}
		}
	}

	return board
}

func (gb *GameBoard) IsFull() bool {
	return gb.MovesCount == 9
}

func (gb *GameBoard) MakeMove(player *Player, row int, col int) error {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return fmt.Errorf("invalid move")
	}
	if gb.Board[row][col].Type.Value != "" {
		return fmt.Errorf("cell already taken")
	}

	gb.Board[row][col].Type.Value = player.Symbol
	gb.MovesCount++
	return nil
}

func (gb *GameBoard) HasWinner() bool {
	//check rows
	for i := 0; i < 3; i++ {
		if gb.Board[i][0].Type.Value != "" && gb.Board[i][0].Type.Value == gb.Board[i][1].Type.Value && gb.Board[i][1].Type.Value == gb.Board[i][2].Type.Value {
			return true
		}
	}

	//check cols
	for i := 0; i < 3; i++ {
		if gb.Board[0][i].Type.Value != "" && gb.Board[0][i].Type.Value == gb.Board[1][i].Type.Value && gb.Board[1][i].Type.Value == gb.Board[2][i].Type.Value {
			return true
		}
	}

	// check diagonal
	for i := 0; i < 3; i++ {
		if gb.Board[0][0].Type.Value != "" && gb.Board[0][0].Type.Value == gb.Board[1][1].Type.Value && gb.Board[1][1].Type.Value == gb.Board[2][2].Type.Value {
			return true
		}
	}

	return gb.Board[0][2].Type.Value != "" && gb.Board[0][2].Type.Value == gb.Board[1][1].Type.Value && gb.Board[1][1].Type.Value == gb.Board[2][0].Type.Value
}

func (gb *GameBoard) PrintGameBoard() {
	fmt.Println("Game Board")
	for _, cell := range gb.Board {
		for j, cell := range cell {
			val := "_"
			if cell != nil && cell.Type != nil && cell.Type.Value != "" {
				val = cell.Type.Value
			}
			fmt.Print(val, " ")

			if (j+1)%3 == 0 {
				fmt.Println()
			}
		}

	}
}

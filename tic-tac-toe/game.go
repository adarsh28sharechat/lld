package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Player1       *Player
	Player2       *Player
	Board         *GameBoard
	CurrentPlayer *Player
}

func NewGame(player1, player2 *Player) *Game {
	return &Game{Board: NewGameBoard(), CurrentPlayer: player1, Player1: player1, Player2: player2}
}

func (g *Game) Play() {
	g.Board.PrintGameBoard()

	for !g.Board.IsFull() && !g.Board.HasWinner() {
		fmt.Printf("%s's turn\n", g.CurrentPlayer.Name)
		row := getValidInput("Enter row btw  0-2: ")
		col := getValidInput("Enter column btw  0-2: ")

		err := g.Board.MakeMove(g.CurrentPlayer, row, col)
		if err != nil {
			fmt.Println(err)
			continue
		}
		g.SwitchPlayer()
		g.Board.PrintGameBoard()
	}

	if g.Board.HasWinner() {
		fmt.Printf("%s's winner\n", g.CurrentPlayer.Name)
	} else {
		fmt.Println("It's draw!")
	}
}

func (g *Game) SwitchPlayer() {
	if g.CurrentPlayer == g.Player1 {
		g.CurrentPlayer = g.Player2
	} else {
		g.CurrentPlayer = g.Player1
	}
}

func getValidInput(placeholder string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(placeholder)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		getValidInput(placeholder)
	}
	return num
}

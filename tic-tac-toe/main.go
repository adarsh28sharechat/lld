package main

import "fmt"

func main() {
	fmt.Println("tic-tac-toe")

	player1 := NewPlayer("Player1", "X")
	player2 := NewPlayer("Player2", "O")

	game := NewGame(player1, player2)
	game.Play()
}

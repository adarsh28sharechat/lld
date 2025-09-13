package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("snake-ladder game")

	gameManager := GetGameManager()
	wg := new(sync.WaitGroup)

	players1 := []string{"Player 1", "Player 2", "Player 3"}
	gameManager.StartNewGame(wg, players1)

	// Start game 2
	players2 := []string{"Player 4", "Player 5"}
	gameManager.StartNewGame(wg, players2)

	fmt.Println("Games started. Check game output above.")
	//Wait for all games to finish
	wg.Wait()
}

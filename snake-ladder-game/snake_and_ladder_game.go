package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type SnakeAndLadderGame struct {
	Id               int64
	Players          []*Player
	Dice             *Dice
	Board            *Board
	CurrentPlayerIdx int
}

func NewSnakeAndLadderGame(playerNames []string) *SnakeAndLadderGame {
	game := &SnakeAndLadderGame{
		Id:               generateID(),
		Players:          []*Player{},
		Dice:             NewDice(),
		Board:            NewBoard(),
		CurrentPlayerIdx: 0,
	}

	for _, name := range playerNames {
		game.Players = append(game.Players, NewPlayer(name))
	}
	return game
}

func (salg *SnakeAndLadderGame) Play(wg *sync.WaitGroup) {
	for !salg.isGameOver() {
		player := salg.Players[salg.CurrentPlayerIdx]
		roll := salg.Dice.Roll()

		newPosition := player.Position + roll

		if newPosition <= salg.Board.Size {
			player.Position = salg.Board.GetNewPosition(newPosition)
			fmt.Printf("Game: %d :- %s rolled a %d and moved to position %d\n", salg.Id, player.Name, roll, player.Position)
		}

		if newPosition == salg.Board.Size {
			fmt.Printf("For Game %d :- %s wins!\n", salg.Id, player.Name)
			break
		}

		salg.CurrentPlayerIdx = (salg.CurrentPlayerIdx + 1) % len(salg.Players)
		time.Sleep(10 * time.Millisecond)
	}
	wg.Done()
}

func (salg *SnakeAndLadderGame) isGameOver() bool {
	for _, player := range salg.Players {
		if player.Position == salg.Board.Size {
			return true
		}
	}
	return false
}

var idCounter int64

func generateID() int64 {
	return atomic.AddInt64(&idCounter, 1)
}

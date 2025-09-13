package main

import "sync"

type GameManger struct {
	Game []*SnakeAndLadderGame
}

var instance *GameManger

func GetGameManager() *GameManger {
	if instance == nil {
		instance = &GameManger{
			Game: []*SnakeAndLadderGame{},
		}
	}
	return instance
}

func (gm *GameManger) StartNewGame(wg *sync.WaitGroup, playerNames []string) {
	game := NewSnakeAndLadderGame(playerNames)
	gm.Game = append(gm.Game, game)
	wg.Add(1)
	go game.Play(wg)
}

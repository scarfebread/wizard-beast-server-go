package entity

import (
	"fmt"
	"log/slog"
)

type PlayerRepository struct {
	players []Player
	ch      chan Player
}

func CreatePlayerRepository() PlayerRepository {
	repository := PlayerRepository{}
	go repository.consumeNewPlayers()
	return repository
}

func (repository PlayerRepository) consumeNewPlayers() {
	repository.ch = make(chan Player)

	for i := range repository.ch {
		slog.Info(fmt.Sprintf("Registered %s", i.Name))
		repository.players = append(repository.players, i)
	}
}

func (repository PlayerRepository) AddPlayer(player Player) {
	repository.ch <- player
}

package player

import (
	"fmt"
	"log/slog"
)

type Repository struct {
	players []Player
	ch      chan Player
}

func CreatePlayerRepository() Repository {
	repository := Repository{}
	go repository.consumeNewPlayers()
	return repository
}

func (repository Repository) consumeNewPlayers() {
	repository.ch = make(chan Player)

	for i := range repository.ch {
		slog.Info(fmt.Sprintf("registered %s", i.Name))
		repository.players = append(repository.players, i)
	}
}

func (repository Repository) AddPlayer(player Player) {
	repository.ch <- player
}

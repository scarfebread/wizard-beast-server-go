package player

import (
	"fmt"
	"log/slog"
)

type Repository struct {
	players []Player
	ch      chan Player
}

func NewRepository() Repository {
	repository := Repository{}
	go repository.consumeNewPlayers()
	return repository
}

func (r Repository) Players() []Player {
	return r.players
}

func (r Repository) AddPlayer(player Player) {
	r.ch <- player
}

func (r Repository) consumeNewPlayers() {
	r.ch = make(chan Player)

	for i := range r.ch {
		slog.Info(fmt.Sprintf("registered %s", i.Name))
		r.players = append(r.players, i)
	}
}

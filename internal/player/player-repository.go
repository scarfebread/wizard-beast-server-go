package player

type Repository struct {
	players []Player
	ch      chan Player
}

func NewRepository() Repository {
	repository := Repository{
		players: make([]Player, 0),
		ch:      make(chan Player),
	}
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
	for i := range r.ch {
		r.players = append(r.players, i)
	}
}

package player

import "slices"

type Repository struct {
	players []Player
	ch      chan Player
}

func NewRepository() Repository {
	repository := Repository{
		players: make([]Player, 0),
		ch:      make(chan Player),
	}
	go repository.consume()
	return repository
}

func (r Repository) Players() []Player {
	return r.players
}

func (r Repository) Player(id string) Player {
	i := slices.IndexFunc(r.Players(), func(p Player) bool { return p.ID == id })

	return r.Players()[i]
}

func (r Repository) Remove(p Player) {
	i := r.index(p)
	r.players = slices.Delete(r.players, i, i+1) // TODO race condition
}

func (r Repository) AddPlayer(player Player) {
	r.ch <- player
}

func (r Repository) index(p Player) int {
	return slices.IndexFunc(r.players, func(p2 Player) bool { return p.ID == p2.ID })
}

func (r Repository) consume() {
	for p := range r.ch {
		if p.delete {
			i := r.index(p)
			r.players = slices.Delete(r.players, i, i+1) // TODO does this work?
		} else {
			r.players = append(r.players, p)
		}
	}
}

package state

import (
	"github.com/scarfebread/wizard-beast-server-go/internal/game/state/publishable"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"slices"
)

var snapshots = make(map[int64]State)

type State struct {
	players []player.Player
}

type delta struct {
	previous player.Player
	current  player.Player
}

type rule interface {
	apply()
}

func Snapshot(id int64, players []player.Player) {
	snapshots[id] = State{
		players: players,
	}

	delete(snapshots, id-64)
}

func BuildState(id int64, p player.Player) publishable.State {
	snapshot := snapshots[id]
	previous := previousState(p)

	delta := delta{
		previous: currentPlayer(previous.players, p),
		current:  currentPlayer(snapshot.players, p),
	}

	return publishable.State{
		ID:     id,
		Player: p,
	}
}

func previousState(p player.Player) State {
	if state, ok := snapshots[p.ConfirmedState]; ok {
		return state
	} else {
		return State{
			players: make([]player.Player, 0),
		}
	}
}

func currentPlayer(players []player.Player, p player.Player) player.Player {
	return players[slices.IndexFunc(players, func(p2 player.Player) bool { return p.ID == p2.ID })] // TODO duplication
}

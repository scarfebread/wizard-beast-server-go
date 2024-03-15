package state

import "github.com/scarfebread/wizard-beast-server-go/internal/player"

var snapshots = make(map[int64]State)

type State struct {
	players []player.Player
}

func Snapshot(id int64, players []player.Player) {
	snapshots[id] = State{
		players: players,
	}

	delete(snapshots, id-64)
}

func BuildState(snapshot int64, player player.Player) string {
	return ""
}

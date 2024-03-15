package game

import (
	"github.com/google/uuid"
	"github.com/scarfebread/wizard-beast-server-go/internal/game/state"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"time"
)

var lastTick time.Time

type Engine struct {
	PlayerRepository player.Repository
}

func (engine Engine) Start() error {
	tick := time.Second / 64
	snapshot := int64(0)

	for {
		var startTime = time.Now()

		engine.publish(snapshot)

		if elapsed := time.Now().Sub(startTime); elapsed < tick {
			time.Sleep(tick - elapsed)
		}

		snapshot++
		lastTick = startTime
	}
}

func (engine Engine) publish(id int64) {
	state.Snapshot(id, engine.PlayerRepository.Players())

	for _, p := range engine.PlayerRepository.Players() {
		p.Client.Send(
			"state",
			state.BuildState(id, p),
			uuid.New().String(),
		)
	}
}

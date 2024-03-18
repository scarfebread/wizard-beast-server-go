package game

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/scarfebread/wizard-beast-server-go/internal/game/state"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"log/slog"
	"time"
)

var lastTick time.Time

type Engine struct {
	PlayerRepository player.Repository
	Simulator        Simulator
}

func (engine Engine) Start() error {
	tick := time.Second / 64
	snapshot := int64(0)

	for {
		var startTime = time.Now()

		state.Snapshot(snapshot, engine.PlayerRepository.Players())
		engine.Simulator.ProcessMovement(lastTick, tick)
		engine.publish(snapshot)

		if elapsed := time.Now().Sub(startTime); elapsed < tick {
			time.Sleep(tick - elapsed)
		} else {
			slog.Warn("engine tick took longer than the allocated time")
		}

		snapshot++
		lastTick = startTime
	}
}

func (engine Engine) publish(id int64) {
	for _, p := range engine.PlayerRepository.Players() {
		jsonState, err := json.Marshal(state.BuildState(id, p))

		if err != nil {
			slog.Error("failed to serialise jsonState", err)
		}

		go p.Client.Send(
			"state",
			string(jsonState),
			uuid.New().String(),
		)
	}
}

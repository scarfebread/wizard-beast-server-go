package game

import (
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"time"
)

type Simulator struct {
	PlayerRepository player.Repository
}

func (sim Simulator) HandleInput(id string, actions []player.Action) {

}

func (sim Simulator) ProcessMovement(lastTick time.Time, tickLength time.Duration) {

}

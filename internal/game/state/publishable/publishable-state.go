package publishable

import (
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
)

type State struct {
	ID          int64          `json:"stateId"`
	Player      player.Player  `json:"player"`
	Players     []PlayerAction `json:"playerActions"`
	Projectiles []string       `json:"projectiles"`
	Enemies     []string       `json:"enemies"`
	Timestamp   int64          `json:"timestamp"`
}

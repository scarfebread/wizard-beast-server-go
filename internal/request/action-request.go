package request

import (
	"encoding/json"
	"github.com/scarfebread/wizard-beast-server-go/internal/game"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"log/slog"
)

type actionRequest struct {
	ID      string          `json:"id"`
	Actions []player.Action `json:"actions"`
}

func ProcessAction(
	data string,
	simulator game.Simulator,
) {
	var req actionRequest

	if err := json.Unmarshal([]byte(data), &req); err != nil {
		slog.Warn("cannot deserialise player action", "action", data, "error", err)
		return
	}

	simulator.HandleInput(req.ID, req.Actions)
}

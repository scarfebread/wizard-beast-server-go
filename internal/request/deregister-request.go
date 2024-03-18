package request

import (
	"encoding/json"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"github.com/scarfebread/wizard-beast-server-go/internal/udp"
	"log/slog"
)

type deregisterRequest struct {
	ID string `json:"id"`
}

func ProcessDeregistration(
	data string,
	repository player.Repository,
	client udp.Client,
) {
	var req deregisterRequest

	if err := json.Unmarshal([]byte(data), &req); err != nil {
		slog.Warn("failed to deserialise deregister request", "request", data)
		return
	}

	p := repository.Player(req.ID)

	if !p.Client.Validate(client) {
		slog.Warn("attempted deregister from another client", "client", client)
	}

	repository.Remove(p)
}

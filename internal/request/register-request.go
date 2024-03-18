package request

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"github.com/scarfebread/wizard-beast-server-go/internal/udp"
	"log/slog"
	"math/rand/v2"
)

type registerRequest struct {
	Name string `json:"name"`
}

func ProcessRegistration(
	id string,
	data string,
	repository player.Repository,
	client udp.Client,
) {
	var req registerRequest

	if err := json.Unmarshal([]byte(data), &req); err != nil {
		slog.Warn(fmt.Sprintf("cannot deserialise %s", data))
		client.Send("invalid", "failed to deserialise", id)
		return
	}

	p := player.Player{
		ID:     uuid.New().String(),
		Name:   req.Name,
		Client: client,
		X:      float32(rand.IntN(800 - 25)),
		Y:      float32(rand.IntN(480 - 25)),
	}
	repository.AddPlayer(p)

	serialisedPlayer, err := p.MarshalJSON()

	if err != nil {
		slog.Warn("failed to serialise player", "name", p.Name)
		client.Send("invalid", fmt.Sprintf("failed to serialise %s", p.Name), id)
		return
	}

	slog.Info("player registered", "name", p.Name, "tmp", id)
	client.Send("registered", string(serialisedPlayer), id)

	if err != nil {
		slog.Error("unable to send message back to the client", err)
	}
}

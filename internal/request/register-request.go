package request

import (
	"encoding/json"
	"fmt"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"github.com/scarfebread/wizard-beast-server-go/internal/udp"
	"log/slog"
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

	if err := req.deserialise(data); err != nil {
		slog.Warn(fmt.Sprintf("cannot deserialise %s", data))
		client.Send("invalid", "failed to deserialise", id)
		return
	}

	p := player.Player{
		Name:   req.Name,
		Client: client,
	}
	repository.AddPlayer(p)

	serialisedPlayer, err := p.Serialise()

	if err != nil {
		slog.Warn(fmt.Sprintf("cannot serialise p %s", p.Name))
		client.Send("invalid", fmt.Sprintf("failed to serialise %s", p.Name), id)
		return
	}

	client.Send("registered", string(serialisedPlayer), id)

	if err != nil {
		slog.Error("unable to send message back to the client", err)
	}
}

func (req *registerRequest) deserialise(data string) error {
	var personData map[string]interface{}

	err := json.Unmarshal([]byte(data), &personData)

	if err != nil {
		return err
	}

	req.Name = personData["name"].(string)

	return nil
}

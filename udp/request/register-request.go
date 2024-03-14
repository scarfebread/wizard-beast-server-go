package request

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"wizard-beast-server-go/player"
	"wizard-beast-server-go/udp"
)

type registerRequest struct {
	Name string `json:"name"`
}

func ProcessRegistration(
	id string,
	data string,
	addr *net.UDPAddr,
	repository player.Repository,
	client udp.Client,
) {
	var req registerRequest

	if err := req.deserialise(data); err != nil {
		slog.Warn(fmt.Sprintf("cannot deserialise %s", data))
		client.Send("invalid", "failed to deserialise", id)
		return
	}

	player := player.Player{
		Name: req.Name,
		Addr: addr,
	}
	repository.AddPlayer(player)

	serialisedPlayer, err := player.Serialise()

	if err != nil {
		slog.Warn(fmt.Sprintf("cannot serialise player %s", player.Name))
		client.Send("invalid", fmt.Sprintf("failed to serialise %s", player.Name), id)
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
